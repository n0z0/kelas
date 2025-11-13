// main_secure.go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	certgen "session/z/z"
)

var sessions = map[string]string{}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/dashboard", dashboardPage)

	// --- PENCEGAHAN: Gunakan HTTPS ---
	fmt.Println("Server aman dimulai di https://localhost:8443")

	// Check if certificates exist, generate them if they don't
	certPath := "cert.pem"
	keyPath := "key.pem"

	if _, err := os.Stat(certPath); os.IsNotExist(err) || os.IsNotExist(os.IsNotExist(err)) {
		fmt.Println("Certificates not found, generating new ones...")
		generateCertificates()
	}

	// Untuk development, gunakan self-signed certificate
	err := http.ListenAndServeTLS(":8443", certPath, keyPath, nil)
	if err != nil {
		fmt.Println("Gagal memulai server:", err)
	}
}

func generateCertificates() {
	config := certgen.DefaultConfig()
	config.OutputDir = "." // Save in current directory
	config.Domains = []string{"localhost", "127.0.0.1"}

	if err := certgen.GenerateCertificateWithConfig(config); err != nil {
		fmt.Printf("Failed to generate certificates: %v\n", err)
		os.Exit(1)
	}

	// Rename files to match expected names
	if err := os.Rename("localhost.crt", "cert.pem"); err != nil {
		fmt.Printf("Failed to rename certificate: %v\n", err)
	}

	if err := os.Rename("localhost.key", "key.pem"); err != nil {
		fmt.Printf("Failed to rename private key: %v\n", err)
	}
}

// Fungsi untuk membuat session ID yang aman dan acak
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			// --- PENCEGAHAN: Regenerasi Session ID ---
			// Buat session ID baru yang aman setelah login berhasil
			newSessionID, err := generateSessionID()
			if err != nil {
				http.Error(w, "Could not create session", http.StatusInternalServerError)
				return
			}

			// --- PENCEGAHAN: Atribut Cookie Aman ---
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    newSessionID,
				HttpOnly: true, // Tidak bisa diakses via JS
				Secure:   true, // Hanya dikirim via HTTPS
				Path:     "/",
				SameSite: http.SameSiteLaxMode,
				MaxAge:   300, // 5 menit
			})

			sessions[newSessionID] = username
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}
	http.ServeFile(w, r, "login.html")
}

func dashboardPage(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Unauthorized: No session", http.StatusUnauthorized)
		return
	}

	username, ok := sessions[sessionID.Value]
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Selamat datang di dashboard aman, %s!", username)
}
