Berikut rancangan slide + ide praktik (lab) berbasis CLI dan Golang sesuai outline di PDF. 

Kamu bisa langsung copy struktur ini ke PowerPoint/Google Slides/LibreOffice Impress.

---

## Slide 1 – Judul & Objective

**Judul:**
**Understanding Session Hijacking Attacks and Countermeasures**

**Poin:**

* Memahami konsep *session hijacking* pada aplikasi web
* Mengenali berbagai jenis serangan session hijacking
* Menggunakan tools untuk mendeteksi dan mencegah session hijacking
* Menerapkan kontrol keamanan di level aplikasi (Golang) dan jaringan

**Praktik (persiapan):**

* Siapkan 1 VM / mesin lab (mis: Ubuntu + Go + curl + tcpdump/Wireshark)
* Install Go: `sudo apt install golang-go`
* Buat folder kerja: `mkdir session-lab && cd session-lab`

---

## Slide 2 – Konsep Dasar Session & Session Hijacking

**Poin:**

* Session = mekanisme server untuk mengingat state pengguna (login)
* Biasanya diwakili oleh **session ID** dalam cookie (`SID`, `PHPSESSID`, dll.)
* *Session hijacking* = penyerang mengambil/menyalahgunakan session ID korban
* Akibat: penyerang bisa “login” sebagai korban tanpa tahu password

**Praktik Ringan (CLI):**

* Akses website demo (atau app Go yang kita buat nanti) dengan `curl`:

  ```bash
  curl -i http://localhost:8080/
  ```
* Perlihatkan header `Set-Cookie: SID=...`

---

## Slide 3 – Gambaran Umum Serangan Session Hijacking

**Poin:**

* Sumber session ID bocor:

  * Sniffing di jaringan (HTTP, WiFi publik)
  * XSS / injeksi di browser
  * Malware (Man-in-the-Browser)
  * Server misconfig (log, referer, dll)
* Konsekuensi:

  * *Account takeover*
  * Transaksi finansial tanpa izin
  * Pencurian data sensitif

**Praktik (demonstrasi konsep):**

* Tunjukkan perbedaan:

  ```bash
  curl -i http://example.com      # HTTP
  curl -i https://example.com     # HTTPS
  ```
* Tekankan bahwa HTTP ⇒ data (termasuk cookie) tidak terenkripsi.

---

## Slide 4 – Types of Session Hijacking (Overview)

**Poin:**

* Spoofing
* Application-level session hijacking
* Man-in-the-Browser (MITB)
* Session replay
* Session fixation
* Network-level session hijacking (TCP/IP hijacking)
* Tools yang sering dipakai: Wireshark, Burp Suite, dkk.

**Praktik (diskusi singkat):**

* Minta mahasiswa menyebutkan skenario nyata (WiFi publik, warnet, dll.)
* Kelompokkan skenario ke jenis serangan di atas.

---

## Slide 5 – Lab Setup: Aplikasi Go Sederhana (Vulnerable Session)

**Poin:**

* Kita pakai **web server Go sederhana** sebagai target latihan
* Fitur:

  * `/login` → set cookie `SID`
  * `/profile` → hanya bisa diakses jika SID valid
* Awalnya sengaja **tidak aman** (tanpa `HttpOnly`, `Secure`, dll.) untuk demo hijacking

**Kode Go (ringkas, 1 file: `main.go`):**

```go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

var sessions = map[string]string{} // SID -> username

func newSID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Demo: username statis
	username := "alice"
	sid := newSID()
	sessions[sid] = username

	http.SetCookie(w, &http.Cookie{
		Name:  "SID",
		Value: sid,
		// sengaja tidak di-set Secure/HttpOnly untuk demo
	})
	fmt.Fprintf(w, "Logged in as %s, SID=%s\n", username, sid)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("SID")
	if err != nil || sessions[c.Value] == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome %s! This is your profile.\n", sessions[c.Value])
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile", profileHandler)
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Praktik:**

```bash
go run main.go
curl -i http://localhost:8080/login
curl -i --cookie "SID=<copy_dari_output_login>" http://localhost:8080/profile
```

---

## Slide 6 – Spoofing dalam Session Hijacking

**Poin:**

* Spoofing = memalsukan identitas (IP, cookie, header, user-agent)
* Dalam konteks session hijacking: **memalsukan session ID** milik korban
* Bisa dilakukan setelah penyerang:

  * Mengendus session ID
  * Menebak (jika pola SID lemah)
  * Mendapatkan dari XSS

**Praktik (CLI – spoofing cookie SID):**

1. Login sebagai “korban” di browser → dapatkan cookie `SID` dari DevTools.
2. Di terminal penyerang:

   ```bash
   curl -i --cookie "SID=<SID_korban>" http://localhost:8080/profile
   ```
3. Tunjukkan bahwa penyerang bisa akses profil seolah-olah korban.

---

## Slide 7 – Application-Level Session Hijacking

**Poin:**

* Terjadi saat aplikasi web:

  * Menggunakan session ID yang bisa ditebak
  * Menyimpan session di tempat tidak aman (URL, log, localStorage tanpa proteksi)
  * Tidak memvalidasi session dengan baik (mis: tidak cek IP/UA/timeout)
* Contoh: SID dikirim via URL `http://site.com/profile?sid=12345`

**Praktik (modifikasi kecil aplikasi Go):**

* Ubah sementara menjadi **SID di URL** (sengaja tidak aman):

```go
func profileHandler(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Query().Get("sid")
	if sid == "" || sessions[sid] == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome %s! This is your profile.\n", sessions[sid])
}
```

**Demo CLI:**

```bash
# 1. login
curl -i http://localhost:8080/login
# 2. gunakan SID di URL
curl "http://localhost:8080/profile?sid=<SID_korban>"
```

Jelaskan bahwa SID di URL mudah masuk ke log, history, referer ⇒ rawan hijack.

---

## Slide 8 – Man-in-the-Browser (MITB)

**Poin:**

* MITB = malware/add-on jahat di browser yang:

  * Mengubah konten halaman
  * Mengintersep & memodifikasi request (termasuk cookie, form)
* Berbeda dengan MITM di jaringan — ini ada di **browser client**
* Sulit dideteksi di jaringan karena kelihatan “normal”

**Praktik (simulasi pakai proxy seperti Burp/OWASP ZAP):**

* Set browser menggunakan proxy lokal (mis: Burp/ZAP)
* Buka `http://localhost:8080/profile`
* Di proxy, ubah manual header `Cookie: SID=...` jadi `SID=<SID_penyerang>`
* Tunjukkan bahwa browser seakan-akan mengirim session korban → hijack.

*(Jika tidak ingin pakai GUI: cukup jelaskan, lalu kembali fokus ke CLI/Go.)*

---

## Slide 9 – Session Replay Attack

**Poin:**

* Penyerang **merekam** request yang valid (termasuk cookie/session token)
* Lalu **memutar ulang** (replay) request tersebut untuk melakukan aksi berulang
* Berbahaya jika:

  * Session tidak kadaluarsa cepat
  * Tidak ada proteksi anti-replay (nonce, timestamp, dll.)

**Praktik (CLI – replay sederhana):**

1. Simpan request login:

   ```bash
   curl -i http://localhost:8080/login -c cookies.txt
   curl -i http://localhost:8080/profile -b cookies.txt
   ```
2. Tunjukkan bahwa file `cookies.txt` bisa disalin ke mesin lain.
3. Di mesin lain (anggap penyerang):

   ```bash
   curl -i http://ip-server:8080/profile -b cookies.txt
   ```
4. Jelaskan bahwa selama SID valid, replay akan berhasil.

---

## Slide 10 – Session Fixation Attack

**Poin:**

* Penyerang **memaksa korban** memakai session ID yang sudah diketahui penyerang
* Langkah umum:

  1. Penyerang membuat SID (atau menebak)
  2. Mengarahkan korban ke URL dengan SID tersebut / memasang cookie di browser korban
  3. Korban login ⇒ SID tetap sama (tidak di-rotate)
  4. Penyerang pakai SID itu untuk akses akun korban
* Inti masalah: aplikasi **tidak mengganti SID setelah login**

**Praktik (Go – demo session fixation):**

1. Modifikasi loginHandler supaya **tidak mengganti SID** jika sudah ada cookie `SID`:

   ```go
   func loginHandler(w http.ResponseWriter, r *http.Request) {
       c, err := r.Cookie("SID")
       var sid string
       if err == nil && c.Value != "" {
           sid = c.Value // FIXATION: pakai SID lama
       } else {
           sid = newSID()
       }
       sessions[sid] = "alice"
       http.SetCookie(w, &http.Cookie{Name: "SID", Value: sid})
       fmt.Fprintf(w, "Logged in with SID=%s\n", sid)
   }
   ```

2. Simulasi:

   * Penyerang: set cookie SID di browser korban (misal via XSS / link).
   * Korban login → server pakai SID yang ditanam penyerang.
   * Penyerang pakai SID yang sama di curl untuk akses `/profile`.

3. Diskusikan cara mitigasi (lihat slide berikutnya).

---

## Slide 11 – Network-Level Session Hijacking

**Poin:**

* Terjadi di level jaringan (TCP/IP):

  * Sniffing paket (HTTP, telnet, dll.)
  * TCP hijacking (menebak sequence number)
  * ARP spoofing, WiFi rogue AP
* Session ID yang lewat di HTTP bisa dilihat dengan Wireshark/tcpdump.

**Praktik (CLI + tcpdump):**

1. Jalankan server Go di `:8080` (HTTP, bukan HTTPS).
2. Di terminal lain:

   ```bash
   sudo tcpdump -i <iface> -A 'tcp port 8080' | grep "Cookie:"
   ```
3. Akses `http://server:8080/profile` dari browser korban.
4. Tunjukkan bahwa header `Cookie: SID=...` muncul di output tcpdump.

---

## Slide 12 – Session Hijacking Tools

**Poin:**

* **Wireshark**: sniffing paket & melihat cookie
* **Burp Suite / OWASP ZAP**: intercept, modify, replay HTTP request
* **Browser devtools**: melihat/mengubah cookie
* Tools khusus (dulu: Firesheep, dll.) mempermudah hijack di WiFi publik

**Praktik (demo singkat):**

* Tampilkan Wireshark filter `http && http.cookie`
* Tunjukkan salah satu request ke server Go:

  * Lihat `Cookie: SID=...`
* (Opsional) Pakai Burp/ZAP untuk replay request ke `/profile`.

---

## Slide 13 – Detection Methods (IDS, WAF, Monitoring)

**Poin:**

* IDS (Intrusion Detection System):

  * Memonitor pola trafik mencurigakan
  * Contoh pola: banyak request dengan SID yang sama dari IP berbeda
* WAF (Web Application Firewall):

  * Memfilter XSS, injection, dll. yang bisa menyebabkan session hijacking
* Log aplikasi:

  * Deteksi login ganda, IP berubah drastis, jam akses tidak wajar

**Praktik (log sederhana di Go):**

* Tambah logging IP dan SID di handler:

  ```go
  func profileHandler(w http.ResponseWriter, r *http.Request) {
      ip := r.RemoteAddr
      c, _ := r.Cookie("SID")
      log.Printf("Access /profile from %s SID=%v", ip, c)
      // ...lanjutan seperti sebelumnya
  }
  ```
* Tampilkan bahwa ketika SID yang sama dipakai dari IP berbeda, log bisa jadi indikator hijack.

---

## Slide 14 – Prevention Tools & Best Practices

**Poin:**

* Gunakan **HTTPS** di semua halaman (bukan hanya login)
* Set cookie:

  * `Secure` (hanya dikirim lewat HTTPS)
  * `HttpOnly` (tidak bisa diakses JavaScript)
  * `SameSite` (mencegah CSRF)
* Rotate session ID:

  * Setiap login, privilege change, dll.
* Session timeout & idle timeout
* Hindari SID di URL

**Praktik (perbaikan kode Go):**

* Ubah `SetCookie` menjadi:

  ```go
  http.SetCookie(w, &http.Cookie{
      Name:     "SID",
      Value:    sid,
      HttpOnly: true,
      Secure:   false, // true jika pakai HTTPS reverse proxy
      SameSite: http.SameSiteStrictMode,
  })
  ```
* Diskusikan:

  * Apa efek `HttpOnly` terhadap XSS?
  * Apa efek `SameSite=Strict` terhadap CSRF?

---

## Slide 15 – Ringkasan & Tugas Mandiri

**Poin Ringkasan:**

* Session hijacking = pengambilalihan session ID
* Jenis: spoofing, application-level, MITB, replay, fixation, network-level
* Tools: Wireshark, Burp/ZAP, logging aplikasi
* Proteksi: HTTPS, cookie aman, session rotation, timeout, IDS/WAF

**Tugas Mandiri (Praktik Lanjutan):**

1. Modifikasi aplikasi Go:

   * Tambah fitur logout (hapus SID dari map dan cookie).
   * Tambah *session timeout* (simulasi dengan menyimpan timestamp).
2. Buat skenario serangan sederhana:

   * Rekam request dengan `curl -v` lalu replay.
   * Catat apa yang berubah ketika session timeout diberlakukan.
3. Tuliskan laporan singkat:

   * 1 halaman: “Skenario Session Hijacking & Cara Mitigasinya di Aplikasi Go”.

---

Kalau mau, saya bisa lanjutkan dengan:

* File **PPTX** yang isinya sudah pakai struktur di atas,
* Atau **versi LaTeX Beamer** kalau kamu biasa presentasi pakai PDF.
