package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/transfer", transfer)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("session")
	if session != nil {
		fmt.Fprintf(w, "Welcome! Session: %s", session.Value)
	} else {
		fmt.Fprintf(w, `<a href="/login">Login</a>`)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// VULNERABLE: Predictable or Fixed
	fixed := r.URL.Query().Get("session_id")
	if fixed != "" {
		// Session Fixation
		http.SetCookie(w, &http.Cookie{Name: "session", Value: fixed})
		fmt.Fprintf(w, "Logged in with fixed session: %s", fixed)
	} else {
		// Predictable
		id := fmt.Sprintf("user%d", time.Now().Unix())
		http.SetCookie(w, &http.Cookie{Name: "session", Value: id})
		fmt.Fprintf(w, "Logged in: %s", id)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("session")
	if session != nil && session.Value == "hacked123" {
		fmt.Fprintf(w, "FLAG: ADMIN_ACCESS_GRANTED")
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func transfer(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("session")
	if session != nil {
		fmt.Fprintf(w, "Transfer $100 completed.")
	} else {
		http.Error(w, "Login required", 401)
	}
}

func secureRandomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
