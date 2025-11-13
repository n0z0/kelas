

Tentu, berikut adalah versi lengkap dari presentasi dan praktikum menggunakan PHP. Struktur ini tetap mempertahankan interaktivitas dan fokus pada konsep keamanan, dengan implementasi yang spesifik untuk ekosistem PHP.

---

### **Struktur File untuk Praktikum PHP**

Kita akan memisahkan aplikasi yang rentan dan yang sudah diamankan ke dalam folder yang berbeda untuk memudahkan perbandingan.

```
/session-hijacking-lab-php
├── vulnerable/              // Aplikasi web rentan
│   ├── index.php
│   ├── login.php
│   ├── dashboard.php
│   └── logout.php
├── secure/                  // Aplikasi web yang sudah diamankan
│   ├── index.php
│   ├── login.php
│   ├── dashboard.php
│   ├── logout.php
│   └── session_config.php   // Konfigurasi sesi yang aman
└── README.md                // Petunjuk singkat (opsional)
```

---

### **Slide Presentasi dan Materi Praktikum (Versi PHP)**

---

#### **Slide 1: Judul**

*   **Judul:** Memahami Serangan Session Hijacking dan Countermeasures
*   **Sub-judul:** Pelatihan Keamanan Aplikasi Web dengan PHP
*   **Nama Presenter:** [Nama Anda]
*   **Tanggal:** [Tanggal Presentasi]

---

#### **Slide 2: Tujuan Pembelajaran**

*   **Tujuan:**
    *   Memahami definisi dan dampak serangan Session Hijacking.
    *   Mengidentifikasi berbagai jenis serangan Session Hijacking (Fixation, Replay, Network-level).
    *   Mampu menggunakan alat (tools) untuk mendemonstrasikan serangan.
    *   Mengimplementasikan kontrol keamanan untuk mencegah serangan menggunakan PHP.
    *   Memahami metode deteksi untuk mengidentifikasi upaya pembajakan sesi.

---

#### **Slide 3: Agenda**

1.  Pengenalan Session Hijacking
2.  Jenis-jenis Serangan Session Hijacking
    *   Session Fixation
    *   Session Replay
    *   Network-Level Hijacking
3.  Tools untuk Session Hijacking
4.  Praktikum: Menyerang Aplikasi Web Vulnerable (PHP)
5.  Metode Deteksi
6.  Praktikum: Implementasi Pencegahan di PHP
7.  Kesimpulan & Q&A

---

#### **Slide 4: Apa itu Session Hijacking?**

*   **Definisi:** Session Hijacking adalah eksploitasi sesi web yang valid—atau "kunci sesi"—untuk mendapatkan akses tidak sah ke informasi atau layanan dalam sistem komputer.
*   **Bagaimana Cara Kerjanya:** Serangan ini memanfaatkan cara aplikasi web mengelola status sesi (state), yang di PHP sering kali dilakukan dengan *session cookie* (`PHPSESSID`).
*   **Konsekuensi Dunia Nyata:**
    *   Pencurian data pribadi atau finansial.
    *   Pengambilalihan akun (misalnya, media sosial, perbankan).
    *   Akses tidak sah ke sistem internal perusahaan.

---

#### **Slide 5: Setup Lab: Aplikasi Web Vulnerable dengan PHP**

*   **Tujuan:** Membuat aplikasi web sederhana dengan PHP yang memiliki kerentanan Session Fixation.
*   **Langkah-langkah:**
    1.  Pastikan PHP dan CLI terinstall.
    2.  Buat struktur file seperti di atas.
    3.  Buat file-file di dalam folder `vulnerable/`.

*   **Instruksi Praktikum (Kode PHP):**
    *   **Berikan kode berikut kepada peserta dan jelaskan bagian-bagian yang vulnerable:**
        *   Session ID bisa di-set dari luar (`$_GET`).
        *   Tidak ada `session_regenerate_id()` setelah login.

    **File: `vulnerable/login.php`**
    ```php
    <?php
    // --- VULNERABILITY: Session Fixation ---
    // Mengizinkan session ID untuk diatur dari URL
    if (isset($_GET['PHPSESSID'])) {
        session_id($_GET['PHPSESSID']);
    }

    session_start();

    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
        $username = $_POST['username'] ?? '';
        $password = $_POST['password'] ?? '';

        if ($username === 'admin' && $password === 'password') {
            $_SESSION['logged_in'] = true;
            $_SESSION['username'] = $username;
            
            // --- VULNERABILITY: TIDAK ADA REGENERASI SESSION ID ---
            // Session ID yang sama digunakan sebelum dan sesudah login
            
            header('Location: dashboard.php');
            exit();
        } else {
            $error = "Username atau password salah!";
        }
    }
    ?>
    <!DOCTYPE html>
    <html>
    <head><title>Login (Vulnerable)</title></head>
    <body>
        <h1>Login</h1>
        <?php if (isset($error)) echo "<p style='color:red;'>$error</p>"; ?>
        <form method="post">
            Username: <input type="text" name="username"><br><br>
            Password: <input type="password" name="password"><br><br>
            <input type="submit" value="Login">
        </form>
    </body>
    </html>
    ```
    **File: `vulnerable/dashboard.php`**
    ```php
    <?php
    session_start();
    if (!isset($_SESSION['logged_in'])) {
        header('Location: login.php');
        exit();
    }
    ?>
    <!DOCTYPE html>
    <html>
    <head><title>Dashboard</title></head>
    <body>
        <h1>Selamat datang, <?php echo htmlspecialchars($_SESSION['username']); ?>!</h1>
        <p>Session ID Anda: <?php echo session_id(); ?></p>
        <a href="logout.php">Logout</a>
    </body>
    </html>
    ```
    *(Buat file `index.php` dan `logout.php` sederhana jika diperlukan)*

*   **Instruksi CLI:**
    *   Buka terminal, arahkan ke folder `session-hijacking-lab-php`.
    *   Jalankan server bawaan PHP:
        ```bash
        php -S localhost:8080 -t vulnerable
        ```
    *   Buka `http://localhost:8080` di browser.

---

#### **Slide 6: Jenis Serangan: Session Fixation (Demo PHP)**

*   **Penjelasan:** Penyerang menetapkan (fix) session ID pengguna *sebelum* pengguna tersebut login.
*   **Alur Serangan:**
    1.  Penyerang membuat session ID (misal `abc123`).
    2.  Penyerang membujuk korban untuk login melalui link: `http://localhost:8080/login.php?PHPSESSID=abc123`.
    3.  Korban login. Aplikasi *tidak* mengubah session ID.
    4.  Penyerang sekarang bisa mengakses akun korban dengan session ID `abc123`.

*   **Instruksi Praktikum (CLI):**
    1.  **Penyerang membuat link berbahaya:**
        `http://localhost:8080/login.php?PHPSESSID=sayabajak`
    2.  **Korban mengklik link dan login** (lakukan di browser atau dengan `curl`).
        ```bash
        # Simulasi login korban dengan session ID yang sudah ditetapkan
        curl -c cookies.txt -b "PHPSESSID=sayabajak" -X POST -d "username=admin&password=password" http://localhost:8080/login.php
        ```
    3.  **Penyerang mengakses dashboard korban menggunakan session ID yang sama:**
        ```bash
        # Penyerang tidak perlu cookie, karena session ID-nya sudah diketahui
        curl -b "PHPSESSID=sayabajak" http://localhost:8080/dashboard.php
        # Output: "<h1>Selamat datang, admin!</h1>...<p>Session ID Anda: sayabajak</p>"
        # SERANGAN BERHASIL!
        ```

---

#### **Slide 7: Jenis Serangan: Session Replay (Demo PHP)**

*   **Penjelasan:** Penyerang menangkap transaksi yang valid dan mengulanginya (replay).
*   **Instruksi Praktikum (CLI):**
    1.  **Mulai menangkap traffic di loopback interface:**
        ```bash
        # Buka terminal baru
        sudo tcpdump -i lo -A 'port 8080'
        ```
    2.  **Korban login dan akses dashboard (gunakan `curl`):**
        ```bash
        # Di terminal lain
        curl -c cookies.txt -X POST -d "username=admin&password=password" http://localhost:8080/login.php
        curl -b cookies.txt http://localhost:8080/dashboard.php
        ```
    3.  **Penyerang menganalisis output `tcpdump`:**
        *   Cari request `GET /dashboard.php HTTP/1.1` dan header `Cookie: PHPSESSID=...`.
    4.  **Penyerang melakukan replay:**
        ```bash
        # Salin cookie yang ditangkap dan lakukan request baru
        curl -H "Cookie: PHPSESSID=[session_id_yang_ditangkap]" http://localhost:8080/dashboard.php
        # Output: "<h1>Selamat datang, admin!</h1>..."
        # SERANGAN BERHASIL!
        ```

---

#### **Slide 8: Pencegahan: Prinsip Utama di PHP**

*   **Gunakan HTTPS:** Mengenkripsi seluruh komunikasi. Mencegah serangan Replay dan Sniffing.
*   **Regenerasi Session ID:** Gunakan `session_regenerate_id(true)` setelah login. `true` akan menghapus sesi lama, mencegah *session fixation*.
*   **Atribut Cookie Aman:** Gunakan `session_set_cookie_params()` untuk mengatur:
    *   `lifetime`: Batas waktu hidup cookie.
    *   `path` & `domain`: Membatasi pengiriman cookie.
    *   `secure`: `true` agar cookie hanya dikirim via HTTPS.
    *   `httponly`: `true` agar tidak bisa diakses via JavaScript (mitigasi XSS).
    *   `samesite`: `'Lax'` atau `'Strict'` untuk melindungi dari CSRF.
*   **Session ID yang Kuat:** PHP modern sudah menggunakan generator yang aman secara default.
*   **Timeout Sesi:** Implementasikan logika timeout di `$_SESSION`.

---

#### **Slide 9: Praktikum: Mengamankan Aplikasi dengan PHP**

*   **Tujuan:** Memperbaiki aplikasi di folder `secure/` dengan menerapkan semua prinsip pencegahan.
*   **Langkah-langkah:**
    1.  Buat file-file di dalam folder `secure/`.

*   **Instruksi Praktikum (Kode PHP):**
    *   **Jelaskan perbaikan yang dilakukan pada setiap file.**

    **File: `secure/session_config.php` (Konfigurasi Pusat)**
    ```php
    <?php
    // Konfigurasi cookie sesi yang aman
    $secureCookieParams = [
        'lifetime' => 600, // Sesi berakhir dalam 10 menit
        'path' => '/',
        'domain' => $_SERVER['HTTP_HOST'],
        'secure' => false, // >>> LIHAT CATATAN DI BAWAH <<<
        'httponly' => true,
        'samesite' => 'Lax'
    ];
    session_set_cookie_params($secureCookieParams);
    ?>
    ```

    **File: `secure/login.php`**
    ```php
    <?php
    require_once 'session_config.php'; // Muat konfigurasi aman
    session_start();

    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
        $username = $_POST['username'] ?? '';
        $password = $_POST['password'] ?? '';

        if ($username === 'admin' && $password === 'password') {
            $_SESSION['logged_in'] = true;
            $_SESSION['username'] = $username;
            $_SESSION['last_activity'] = time(); // Untuk timeout
            
            // --- PENCEGAHAN: Regenerasi Session ID ---
            // Buat session ID baru dan hapus yang lama
            session_regenerate_id(true);

            header('Location: dashboard.php');
            exit();
        }
    }
    ?>
    <!DOCTYPE html>
    <html>
    <head><title>Login (Secure)</title></head>
    <body>
        <h1>Login</h1>
        <form method="post">
            Username: <input type="text" name="username"><br><br>
            Password: <input type="password" name="password"><br><br>
            <input type="submit" value="Login">
        </form>
    </body>
    </html>
    ```
    *(Buat `dashboard.php` dan `logout.php` yang serupa, pastikan `require_once 'session_config.php'; session_start();` ada di setiap halaman yang membutuhkan sesi).*

*   **Instruksi CLI:**
    1.  **Jalankan server aman di port berbeda:**
        ```bash
        php -S localhost:8443 -t secure
        ```
    2.  **Akses di browser:** `http://localhost:8443`
    3.  **Ulangi serangan Session Fixation dari Slide 6.** Perhatikan bahwa session ID akan berubah setelah login, sehingga serangan gagal.

> **⚠️ Catatan Penting tentang HTTPS:**
> Server bawaan PHP (`php -S`) tidak memiliki dukungan TLS/HTTPS bawaan. Oleh karena itu, kita set `'secure' => false` pada konfigurasi cookie agar sesi tetap berjalan di lab ini.
> **Di lingkungan produksi, Anda WAJIB menggunakan web server (Nginx, Apache) dengan HTTPS dan mengatur `'secure' => true`.**

---

#### **Slide 10: Deteksi Session Hijacking di PHP**

*   **Metode Deteksi:**
    *   **Analisis Log & Perilaku Anomali:**
        *   Satu session ID digunakan dari banyak alamat IP.
        *   Perubahan `User-Agent` header untuk satu session ID yang sama.
*   **Praktikum (Kode PHP):**
    *   Tambahkan kode berikut di awal setiap halaman yang dilindungi (setelah `session_start()`).

    ```php
    <?php
    session_start();
    // ... kode login lainnya ...

    // --- DETEKSI ANOMALI SESI ---
    if (isset($_SESSION['user_ip']) === false) {
        $_SESSION['user_ip'] = $_SERVER['REMOTE_ADDR'];
    }

    if (isset($_SESSION['user_agent']) === false) {
        $_SESSION['user_agent'] = $_SERVER['HTTP_USER_AGENT'];
    }

    if ($_SESSION['user_ip'] !== $_SERVER['REMOTE_ADDR'] || $_SESSION['user_agent'] !== $_SERVER['HTTP_USER_AGENT']) {
        // ANOMALI TERDETEKSI!
        // Log ke file, kirim notifikasi, atau hancurkan sesi
        error_log("Session Hijacking Detected for Session ID: " . session_id());
        session_destroy();
        die("Akses diblokir karena aktivitas mencurigakan.");
    }
    // ... lanjutkan dengan logik normal ...
    ?>
    ```

---

#### **Slide 11: Kesimpulan**

*   **Ringkasan:**
    *   Session Hijacking adalah ancaman serius.
    *   Memahami serangan (Fixation, Replay) adalah kunci pertahanan.
    *   Pencegahan jauh lebih baik daripada mengatasi serangan.
*   **Pilar Keamanan Sesi di PHP:**
    1.  **HTTPS SELALU** (di produksi).
    2.  **`session_regenerate_id(true)`** setelah login.
    3.  Gunakan **`session_set_cookie_params()`** dengan flag `Secure`, `HttpOnly`, `SameSite`.
    4.  Simpan **IP Address dan User-Agent** di sesi untuk deteksi anomali.
    5.  Implementasikan **timeout sesi** yang wajar.

---

#### **Slide 12: Q&A**

*   **Judul:** Tanya Jawab
*   **Konten:** Terima kasih atas perhatiannya. Silakan ajukan pertanyaan.