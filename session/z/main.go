// main.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var sessions = map[string]string{"user123": "attacker"} // Simulasi sesi yang sudah dibajak

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/dashboard", dashboardPage)
	fmt.Println("Server vulnerable dimulai di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html") // Asumsikan ada file HTML sederhana
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			// --- VULNERABILITY: Session Fixation ---
			// Menggunakan session ID yang sudah ada atau yang mudah ditebak
			sessionID, err := r.Cookie("session_id")
			if err != nil {
				// Jika tidak ada cookie, buat baru yang mudah ditebak
				newSessionID := "session-" + strconv.FormatInt(time.Now().Unix(), 10)
				http.SetCookie(w, &http.Cookie{
					Name:  "session_id",
					Value: newSessionID,
				})
				sessionID = &http.Cookie{Value: newSessionID}
			}
			// --- TIDAK ADA REGENERASI SESSION ID ---

			// Simpan sesi (dalam kasus nyata, di database)
			sessions[sessionID.Value] = username
			http.Redirect(w, r, "/dashboard?session_id="+sessionID.Value, http.StatusSeeOther)
			return
		}
	}
	http.ServeFile(w, r, "login.html")
}

func dashboardPage(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		// Coba ambil dari query string untuk keperluan demo fixation
		sessionID = &http.Cookie{Value: r.URL.Query().Get("session_id")}
	}

	username, ok := sessions[sessionID.Value]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Selamat datang, %s! Session ID Anda: %s", username, sessionID.Value)
}
