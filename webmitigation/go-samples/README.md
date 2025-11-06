# Go Samples for Day 2 Labs (OWASP Top 10)

Minimal Go service to exercise headers, rate limiting, and SSRF-safe fetching.

## Requirements
- Go 1.21+

## Run
- `cd webmitigation/go-samples`
- `go mod tidy`
- `go run .`
- Service listens on `:8085` (override with `PORT` env)

## Endpoints
- `GET /` → health check
- `GET /login?user=ID` → sets secure HttpOnly cookie `session` and returns header `X-Demo-User-ID`
- `GET /users?id=ID` → requires ownership: header `X-User-ID: ID` (or cookie set by /login)
- `GET /fetch?url=...` → SSRF-safe fetch: allows only http/https and blocks private/local IPs

## Middlewares
- Security headers: CSP default-src 'self', X-Content-Type-Options, X-Frame-Options, Referrer-Policy, HSTS (when HTTPS)
- Rate limiting: basic fixed window (100 req/min per IP) — demo only

## Test Quickly
- Headers: `curl -I http://localhost:8085/`
- Rate limit: `for /l %i in (1,1,120) do curl -s http://localhost:8085/ >nul` (PowerShell/CMD) — expect HTTP 429 past 100 req/min
- Ownership: `curl -H "X-User-ID: 2" "http://localhost:8085/users?id=2"` → 200; change header to `1` → 403
- SSRF: `curl "http://localhost:8085/fetch?url=http://example.com"` → OK; try `http://127.0.0.1` → blocked

## Notes
- For HSTS and Secure cookies, run behind TLS (e.g., reverse proxy/ngrok) or set `X-Forwarded-Proto: https` via proxy.
- This sample avoids external deps to keep labs simple.

