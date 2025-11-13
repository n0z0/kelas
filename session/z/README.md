

Tentu, saya akan membantu Anda membuat slide presentasi lengkap dengan praktikum berbasis CLI dan Go, berdasarkan outline yang diberikan.

Struktur ini dirancang untuk presentasi yang interaktif, di mana setiap konsep teori diikuti oleh demonstrasi praktis. Kode Go akan dibuat sesederhana mungkin untuk fokus pada konsep keamanan, bukan kompleksitas aplikasi.

---

### **Struktur File untuk Praktikum**

Sebelum memulai, kita akan menyiapkan struktur file untuk praktikum. Ini akan membantu peserta mengikuti dengan teratur.

```
/session-hijacking-lab
├── main.go                 // Aplikasi web vulnerable (sebelum diperbaiki)
├── main_secure.go         // Aplikasi web yang sudah diamankan (setelah diperbaiki)
├── generate_cert.go       // Skrip untuk membuat sertifikat SSL/TLS
├── cookies.txt             // File untuk menyimpan cookie saat praktikum
└── README.md              // Petunjuk singkat (opsional)
```

---

### **Slide Presentasi dan Materi Praktikum**

Berikut adalah rincian slide beserta instruksi praktikumnya.

---

#### **Slide 1: Judul**

*   **Judul:** Memahami Serangan Session Hijacking dan Countermeasures
*   **Sub-judul:** Pelatihan Keamanan Aplikasi Web
*   **Nama Presenter:** [Nama Anda]
*   **Tanggal:** [Tanggal Presentasi]

---

#### **Slide 2: Tujuan Pembelajaran**

*   **Tujuan:**
    *   Memahami definisi dan dampak serangan Session Hijacking.
    *   Mengidentifikasi berbagai jenis serangan Session Hijacking (Fixation, Replay, Network-level).
    *   Mampu menggunakan alat (tools) untuk mendemonstrasikan serangan.
    *   Mengimplementasikan kontrol keamanan untuk mencegah serangan menggunakan Go.
    *   Memahami metode deteksi untuk mengidentifikasi upaya pembajakan sesi.

---

#### **Slide 3: Agenda**

1.  Pengenalan Session Hijacking
2.  Jenis-jenis Serangan Session Hijacking
    *   Session Fixation
    *   Session Replay
    *   Network-Level Hijacking
3.  Tools untuk Session Hijacking
4.  Praktikum: Menyerang Aplikasi Web Vulnerable
5.  Metode Deteksi
6.  Praktikum: Implementasi Pencegahan di Go
7.  Kesimpulan & Q&A

---

#### **Slide 4: Apa itu Session Hijacking?**

*   **Definisi:** Session Hijacking adalah eksploitasi sesi web yang valid—atau "kunci sesi"—untuk mendapatkan akses tidak sah ke informasi atau layanan dalam sistem komputer.
*   **Bagaimana Cara Kerjanya:** Serangan ini memanfaatkan cara aplikasi web mengelola status sesi (state), yang sering kali dilakukan dengan token sesi (Session ID/Token).
*   **Konsekuensi Dunia Nyata:**
    *   Pencurian data pribadi atau finansial.
    *   Pengambilalihan akun (misalnya, media sosial, perbankan).
    *   Akses tidak sah ke sistem internal perusahaan.

---

#### **Slide 5: Setup Lab: Aplikasi Web Vulnerable**

*   **Tujuan:** Membuat aplikasi web sederhana dengan Go yang memiliki kerentanan.
*   **Langkah-langkah:**
    1.  Pastikan Go terinstall.
    2.  Buat direktori `session-hijacking-lab`.
    3.  Buat file `main.go`.

*   **Instruksi Praktikum (Kode Go):**
    *   **Berikan kode berikut kepada peserta dan jelaskan bagian-bagian yang vulnerable:**
        *   Session ID tidak di-regenerasi setelah login (rentan terhadap Session Fixation).
        *   Komunikasi menggunakan HTTP (rentan terhadap sniffing/replay).
        *   Session ID tidak memiliki atribut keamanan (`Secure`, `HttpOnly`).

    ```go
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
    ```
    *   **Instruksi CLI:**
        *   `go run main.go`
        *   Buka `http://localhost:8080` di browser.

---

#### **Slide 6: Jenis Serangan: Session Fixation**

*   **Penjelasan:** Penyerang menetapkan (fix) session ID pengguna *sebelum* pengguna tersebut login. Setelah pengguna login dengan session ID tersebut, penyerang bisa menggunakannya untuk mengakses sesi yang sah.
*   **Alur Serangan:**
    1.  Penyerang mendapatkan session ID dari aplikasi.
    2.  Penyerang membujuk korban untuk menggunakan session ID tersebut (misal, melalui link: `http://app.com/login?session_id=XYZ`).
    3.  Korban login. Aplikasi *tidak* mengubah session ID.
    4.  Penyerang sekarang bisa mengakses akun korban dengan session ID `XYZ`.

*   **Instruksi Praktikum (CLI):**
    1.  **Penyerang mendapatkan session ID:**
        ```bash
        curl -c cookies.txt -v http://localhost:8080
        # Lihat output untuk cookie 'session_id' yang dibuat
        cat cookies.txt
        ```
    2.  **Penyerang mengirim link kepada korban (misal: `http://localhost:8080/login?session_id=session-1678886400`)**
    3.  **Korban login melalui link tersebut.** (Lakukan di browser atau dengan `curl`)
        ```bash
        # Simulasi login korban
        curl -c cookies.txt -b cookies.txt -X POST -d "username=admin&password=password" "http://localhost:8080/login?session_id=session-1678886400"
        ```
    4.  **Penyerang mengakses dashboard korban:**
        ```bash
        # Penyerang menggunakan cookie yang sama
        curl -b cookies.txt http://localhost:8080/dashboard
        # Output: "Selamat datang, admin! Session ID Anda: session-1678886400"
        # SERANGAN BERHASIL!
        ```

---

#### **Slide 7: Jenis Serangan: Session Replay**

*   **Penjelasan:** Penyerang menangkap transaksi yang valid (misalnya, request HTTP dengan session cookie) dan mengulanginya (replay) untuk mendapatkan akses.
*   **Alur Serangan:**
    1.  Penyerang berada di jaringan yang sama dengan korban (misal, WiFi publik).
    2.  Penyerang menggunakan packet sniffer (seperti Wireshark atau `tcpdump`) untuk menangkap traffic.
    3.  Korban login dan mengakses halaman dashboard.
    4.  Penyerang menemukan request `GET /dashboard` dengan cookie `session_id=...`.
    5.  Penyerang membuat request yang sama persis untuk mengakses dashboard.

*   **Instruksi Praktikum (CLI):**
    1.  **Mulai menangkap traffic di loopback interface:**
        ```bash
        # Buka terminal baru
        sudo tcpdump -i lo -A 'port 8080'
        ```
    2.  **Korban login dan akses dashboard (gunakan `curl` agar mudah dilihat):**
        ```bash
        # Di terminal lain
        curl -c cookies.txt -b cookies.txt -X POST -d "username=admin&password=password" http://localhost:8080/login
        curl -b cookies.txt http://localhost:8080/dashboard
        ```
    3.  **Penyerang menganalisis output `tcpdump`:**
        *   Cari baris `GET /dashboard HTTP/1.1` dan header `Cookie: session_id=...`.
    4.  **Penyerang melakukan replay:**
        ```bash
        # Salin cookie yang ditangkap dan lakukan request baru
        curl -H "Cookie: session_id=session-1678886400" http://localhost:8080/dashboard
        # Output: "Selamat datang, admin! Session ID Anda: session-1678886400"
        # SERANGAN BERHASIL!
        ```

---

#### **Slide 8: Pencegahan: Prinsip Utama**

*   **Gunakan HTTPS (TLS):** Mengenkripsi seluruh komunikasi antara klien dan server. Ini mencegah serangan Replay dan Network-level Sniffing.
*   **Regenerasi Session ID:** Selalu buat session ID baru setelah peristiwa penting seperti login atau perubahan hak akses. Ini mencegah Session Fixation.
*   **Atribut Cookie yang Aman:**
    *   `Secure`: Memastikan cookie hanya dikirim melalui HTTPS.
    *   `HttpOnly`: Mencegah akses cookie melalui JavaScript client-side (membantu mitigasi XSS).
    *   `SameSite=Strict` atau `Lax`: Melindungi dari Cross-Site Request Forgery (CSRF).
*   **Session ID yang Kuat dan Acak:** Gunakan generator angka acak yang kriptografis aman untuk membuat session ID.
*   **Timeout Sesi:** Berikan masa hidup yang singkat pada sesi yang tidak aktif.

---

#### **Slide 9: Praktikum: Mengamankan Aplikasi dengan Go**

*   **Tujuan:** Memperbaiki `main.go` menjadi `main_secure.go` dengan menerapkan semua prinsip pencegahan.
*   **Langkah-langkah:**
    1.  Buat file `main_secure.go`.
    2.  Generate sertifikat TLS untuk HTTPS.

*   **Instruksi Praktikum (Kode Go):**
    *   **Berikan kode berikut dan jelaskan perbaikan yang dilakukan.**

    ```go
    // main_secure.go
    package main

    import (
        "crypto/rand"
        "encoding/base64"
        "fmt"
        "net/http"
        "time"
    )

    var sessions = map[string]string{}

    func main() {
        http.HandleFunc("/", homePage)
        http.HandleFunc("/login", loginPage)
        http.HandleFunc("/dashboard", dashboardPage)
        
        // --- PENCEGAHAN: Gunakan HTTPS ---
        fmt.Println("Server aman dimulai di https://localhost:8443")
        // Untuk development, gunakan self-signed certificate
        err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
        if err != nil {
            fmt.Println("Gagal memulai server:", err)
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
                    HttpOnly: true,  // Tidak bisa diakses via JS
                    Secure:   true,  // Hanya dikirim via HTTPS
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
    ```

*   **Instruksi CLI (Generate Sertifikat & Jalankan Server):**
    1.  **Generate sertifikat self-signed:**
        ```bash
        go run generate_cert.go
        # (atau gunakan openssl: openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes)
        ```
        *Isi `generate_cert.go` dengan skrip sederhana untuk generate cert/key.*
    2.  **Jalankan server aman:**
        ```bash
        go run main_secure.go
        ```
    3.  **Akses di browser:** `https://localhost:8443` (akan ada peringatan keamanan, lanjutkan saja).

---

#### **Slide 10: Deteksi Session Hijacking**

*   **Metode Deteksi:**
    *   **Analisis Log & Perilaku Anomali:**
        *   Satu session ID digunakan dari banyak alamat IP secara bersamaan.
        *   Perubahan `User-Agent` header untuk satu session ID yang sama.
        *   Lonjakan aktivitas yang tidak wajar dari satu sesi.
    *   **Intrusion Detection System (IDS) / Web Application Firewall (WAF):**
        *   Banyak WAF modern memiliki aturan bawaan untuk mendeteksi anomali sesi.
        *   IDS dapat memantau jaringan untuk pola traffic yang mencurigakan.
*   **Praktikum (Konseptual di Go):**
    *   Tambahkan middleware ke `main_secure.go` untuk logging IP dan User-Agent.
    *   Jika IP atau User-Agent berubah untuk session ID yang sama, log sebagai peringatan.

    ```go
    // Middleware untuk logging dan deteksi (contoh)
    func loggingMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            sessionID, err := r.Cookie("session_id")
            if err == nil {
                // Di sini Anda bisa menyimpan IP dan User-Agent pertama kali
                // dan membandingkannya di request selanjutnya.
                fmt.Printf("[LOG] SessionID: %s, IP: %s, UserAgent: %s\n", sessionID.Value, r.RemoteAddr, r.UserAgent())
                // Logika deteksi anomali bisa ditambahkan di sini
            }
            next.ServeHTTP(w, r)
        })
    }
    // Gunakan middleware ini di main.go:
    // http.Handle("/", loggingMiddleware(http.HandlerFunc(homePage)))
    ```

---

#### **Slide 11: Kesimpulan**

*   **Ringkasan:**
    *   Session Hijacking adalah ancaman serius yang dapat mengakibatkan kerugian besar.
    *   Memahami berbagai jenis serangan (Fixation, Replay) adalah kunci untuk pertahanan.
    *   Pencegahan jauh lebih baik daripada mengatasi serangan.
*   **Pilar Keamanan Sesi:**
    1.  **HTTPS SELALU.**
    2.  **Regenerasi Session ID** setelah login.
    3.  Gunakan atribut cookie **`Secure`**, **`HttpOnly`**, dan **`SameSite`**.
    4.  Buat **Session ID yang acak dan tidak dapat ditebak**.
    5.  Implementasikan **timeout sesi** yang wajar.
    6.  Aktifkan **logging dan monitoring** untuk deteksi dini.

---

#### **Slide 12: Q&A**

*   **Judul:** Tanya Jawab
*   **Konten:** Terima kasih atas perhatiannya. Silakan ajukan pertanyaan.

---

### **Catatan Tambahan untuk Presenter:**

*   **File HTML:** Siapkan file `index.html` dan `login.html` yang sangat sederhana agar praktikum berjalan lancar.
*   **Interaktivitas:** Dorong peserta untuk menjalankan perintah CLI sendiri dan memodifikasi kode Go untuk melihat efeknya.
*   **Waktu:** Alokasikan waktu yang cukup untuk setiap sesi praktikum. Sesi praktikum biasanya memakan waktu paling lama.
*   **Koneksi Internet:** Pastikan ada koneksi internet untuk menginstall dependensi jika diperlukan (meskipun lab ini dirancang tanpa dependensi luar).