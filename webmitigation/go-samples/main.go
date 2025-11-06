package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "strings"
    "sync"
    "time"
)

func main() {
    mux := http.NewServeMux()

    // Public endpoints
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ok")
    })

    // Login demo: sets secure cookie (for HTTPS demo; see README)
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        userID := r.URL.Query().Get("user")
        if userID == "" {
            userID = "1"
        }
        http.SetCookie(w, &http.Cookie{
            Name:     "session",
            Value:    "u:" + userID,
            Path:     "/",
            HttpOnly: true,
            Secure:   isHTTPS(r),
            SameSite: http.SameSiteStrictMode,
        })
        // also pass user via header for lab simplicity
        w.Header().Set("X-Demo-User-ID", userID)
        fmt.Fprintf(w, "logged in as %s\n", userID)
    })

    // Ownership-protected resource: /users?id=123 requires X-User-ID==id
    mux.Handle("/users", AuthorizeOwnership(func(id string) (string, error) {
        // demo: owner == id
        if id == "" {
            return "", errors.New("missing id")
        }
        return id, nil
    }, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        _ = json.NewEncoder(w).Encode(map[string]any{"id": id, "name": "user" + id})
    })))

    // SSRF-safe fetch: /fetch?url=...
    mux.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {
        raw := r.URL.Query().Get("url")
        if !validateURL(raw) {
            http.Error(w, "invalid or blocked URL", http.StatusBadRequest)
            return
        }
        ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
        defer cancel()
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
        if err != nil {
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            http.Error(w, "fetch error", http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        io.CopyN(w, resp.Body, 2048) // limit response size for demo
    })

    // Compose middlewares
    h := SecurityHeaders(RateLimit(mux))

    addr := ":8085"
    if v := os.Getenv("PORT"); v != "" {
        if _, err := strconv.Atoi(v); err == nil {
            addr = ":" + v
        }
    }
    log.Printf("listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, h))
}

// --- Middlewares ---

// SecurityHeaders adds common security headers. In production, enable HSTS only over HTTPS.
func SecurityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("Referrer-Policy", "no-referrer")
        if isHTTPS(r) {
            w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        }
        next.ServeHTTP(w, r)
    })
}

// RateLimit implements a basic fixed-window limit per IP (demo only).
func RateLimit(next http.Handler) http.Handler {
    var (
        mu    sync.Mutex
        hits  = map[string]int{}
        start = time.Now()
        max   = 100 // requests per minute
    )
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        if time.Since(start) > time.Minute {
            hits = map[string]int{}
            start = time.Now()
        }
        ip := clientIP(r)
        hits[ip]++
        n := hits[ip]
        mu.Unlock()
        if n > max {
            http.Error(w, "too many requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// AuthorizeOwnership ensures X-User-ID header matches resource owner.
type OwnerFunc func(id string) (string, error)

func AuthorizeOwnership(getOwner OwnerFunc, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := userIDFromRequest(r)
        id := r.URL.Query().Get("id")
        owner, err := getOwner(id)
        if err != nil || userID == "" || owner != userID {
            http.Error(w, "forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// --- Helpers ---

func userIDFromRequest(r *http.Request) string {
    // Prefer header set by upstream auth; fallback parse demo cookie
    if v := r.Header.Get("X-User-ID"); v != "" {
        return v
    }
    if c, err := r.Cookie("session"); err == nil {
        if strings.HasPrefix(c.Value, "u:") {
            return strings.TrimPrefix(c.Value, "u:")
        }
    }
    return ""
}

func clientIP(r *http.Request) string {
    if xf := r.Header.Get("X-Forwarded-For"); xf != "" {
        parts := strings.Split(xf, ",")
        return strings.TrimSpace(parts[0])
    }
    host, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return r.RemoteAddr
    }
    return host
}

func isHTTPS(r *http.Request) bool {
    if r.TLS != nil || strings.EqualFold(r.Header.Get("X-Forwarded-Proto"), "https") {
        return true
    }
    return false
}

func validateURL(raw string) bool {
    u, err := url.Parse(raw)
    if err != nil {
        return false
    }
    if u.Scheme != "http" && u.Scheme != "https" {
        return false
    }
    host := u.Hostname()
    // Disallow localhost and metadata endpoints explicitly
    if host == "localhost" || host == "127.0.0.1" || host == "::1" || host == "169.254.169.254" {
        return false
    }
    ips, err := net.LookupIP(host)
    if err != nil || len(ips) == 0 {
        return false
    }
    for _, ip := range ips {
        if isPrivateIP(ip) {
            return false
        }
    }
    return true
}

func isPrivateIP(ip net.IP) bool {
    privateCIDRs := []string{
        "10.0.0.0/8",
        "172.16.0.0/12",
        "192.168.0.0/16",
        "127.0.0.0/8",
        "::1/128",
        "fc00::/7",
    }
    for _, c := range privateCIDRs {
        _, block, _ := net.ParseCIDR(c)
        if block.Contains(ip) {
            return true
        }
    }
    return false
}

