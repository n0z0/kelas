

Tentu, saya akan bantu membuat konten presentasi lengkap dengan sesi praktik berdasarkan outline yang Anda berikan. Presentasi ini dirancang untuk sesi pelatihan satu hari, dengan kombinasi teori, demonstrasi, dan hands-on lab.

---

### **Presentasi: Memahami Serangan Denial-of-Service (DoS) dan Distributed Denial-of-Service (DDoS)**

---

#### **Slide 1: Judul**

**Memahami Serangan Denial-of-Service (DoS) dan Distributed Denial-of-Service (DDoS)**

*   Teknik, Deteksi, dan Strategi Perlindungan
*   [Nama Instruktur/Fasilitator]
*   [Tanggal]

---

#### **Slide 2: Tujuan Pembelajaran**

**Tujuan Pembelajaran:**

Setelah mengikuti sesi ini, peserta diharapkan mampu untuk:

1.  **Memahami** konsep dasar, karakteristik, dan perbedaan antara serangan DoS dan DDoS.
2.  **Mengidentifikasi** berbagai teknik serangan DoS/DDoS yang umum digunakan.
3.  **Menganalisis** peran botnet dalam melancarkan serangan DDoS.
4.  **Mendeteksi** indikasi serangan DoS/DDoS menggunakan teknik analisis lalu lintas.
5.  **Menerapkan** kontrol keamanan dan menggunakan alat untuk mitigasi serangan.

---

#### **Slide 3: Agenda Hari Ini**

**Agenda:**

1.  **(09:00 - 10:00) Sesi 1: Gambaran Umum Serangan DoS dan DDoS**
    *   Definisi, Karakteristik, dan Dampak
    *   Studi Kasus di Dunia Nyata
2.  **(10:00 - 10:45) Sesi 2: Botnet**
    *   Penjelasan dan Peran dalam DDoS
    *   Cara Kerja dan Deteksi
3.  **(10:45 - 11:00) Coffee Break**
4.  **(11:00 - 12:30) Sesi 3: Teknik Serangan DoS/DDoS**
    *   SYN Flood, UDP Flood, HTTP Flood, dll.
    *   **Praktik 1: Simulasi Serangan**
5.  **(12:30 - 13:30) Istirahat Makan Siang**
6.  **(13:30 - 14:30) Sesi 4: Alat untuk Serangan DoS/DDoS**
    *   LOIC, HOIC, Slowloris, Hping3
    *   **Praktik 2: Hands-on dengan Alat Serangan**
7.  **(14:30 - 15:30) Sesi 5: Teknik Deteksi Serangan**
    *   Anomaly-based, Signature-based, Traffic Analysis
    *   **Praktik 3: Analisis Lalu Lintas Serangan**
8.  **(15:30 - 15:45) Coffee Break**
9.  **(15:45 - 16:45) Sesi 6: Alat dan Strategi Perlindungan**
    *   Cloudflare, Akamai, Konfigurasi Firewall
    *   **Praktik 4: Implementasi Perlindungan Dasar**
10. **(16:45 - 17:00) Sesi 7: Kesimpulan dan Tanya Jawab**

---

### **Persiapan Lingkungan Praktik (Lab)**

*   **Mesin Attacker:** Kali Linux (atau distro Linux lain dengan tools seperti `hping3`, `slowloris`).
*   **Mesin Target:** Server sederhana (bisa Ubuntu/Debian) dengan layanan web (Apache/Nginx) dan SSH.
*   **Monitoring Tool:** Wireshark terinstal di mesin attacker atau target.
*   **Jaringan:** Kedua mesin berada dalam satu jaringan lokal yang sama (misalnya, dalam VirtualBox Host-Only Network).
*   **Penting:** Semua praktik dilakukan di lingkungan terisolasi (lab virtual). DILARANG KERAS melakukan serangan ke jaringan atau sistem di luar lab yang telah disediakan.

---

### **Konten dan Praktik Per Sesi**

---

#### **Sesi 1: Gambaran Umum Serangan DoS dan DDoS**

*   **Definisi DoS:** Serangan yang bertujuan untuk membuat layanan tidak dapat diakses oleh pengguna yang sah dengan menghabiskan sumber daya (bandwidth, memori, CPU).
*   **Definisi DDoS:** Serangan DoS yang dilakukan dari banyak sumber secara bersamaan dan terdistribusi, biasanya melalui botnet.
*   **Karakteristik Utama:**
    *   Volume lalu lintas yang tidak wajar.
    *   Permintaan dari berbagai alamat IP.
    *   Layanan menjadi sangat lambat atau tidak merespons.
*   **Dampak:**
    *   Kerugian finansial (gangguan bisnis).
    *   Kerusakan reputasi.
    *   Kegagalan infrastruktur kritikal.

**Praktik/Demonstrasi 1: Diskusi Kelompok - Analisis Kasus**

*   **Aktivitas:** Bagi peserta menjadi beberapa kelompok. Berikan mereka 2 studi kasus terkenal:
    1.  **Serangan DDoS pada Dyn (2016):** Menyebabkan situs besar seperti Twitter, Netflix, dan Spotify down di AS.
    2.  **Serangan DDoS pada GitHub (2018):** Serangan terbesar pada saat itu dengan volume 1.35 Tbps.
*   **Tugas Diskusi:**
    *   Apa kemungkinan target serangan ini?
    *   Menurut Anda, apa dampak bisnis yang ditimbulkan?
    *   Teknik serangan apa yang mungkin digunakan (berdasarkan pengetahuan awal)?
*   **Hasil:** Setiap kelompok mempresentasikan hasil diskusinya. Fasilitator memberikan penjelasan lebih lanjut.

---

#### **Sesi 2: Botnet**

*   **Apa itu Botnet?** Jaringan komputer yang telah disusupi malware (bot) dan dapat dikendalikan secara remote oleh penyerang (botmaster).
*   **Peran dalam DDoS:** Botnet adalah sumber daya utama untuk melancarkan serangan DDoS karena volume serangan yang besar dan berasal dari ribuan IP berbeda.
*   **Cara Kerja:**
    *   **Infeksi:** Melalui email phishing, eksploitasi kerentanan, drive-by download.
    *   **Komunikasi (C&C - Command & Control):** Botmaster mengirimkan perintah ke bot melalui server pusat (tradisional) atau jaringan peer-to-peer (modern).
    *   **Eksekusi:** Bot menerima perintah (misalnya, "serang target X dengan SYN flood") dan menjalankannya secara bersamaan.

**Praktik/Demonstrasi 2: Simulasi Arus Lalu Lintas Botnet**

*   **Tujuan:** Menunjukkan bagaimana lalu lintas dari banyak sumber (terlihat seperti botnet) membanjiri sebuah target.
*   **Langkah-langkah:**
    1.  Di mesin **target**, jalankan web server dan buka terminal untuk memantau koneksi masuk: `watch -n 1 'netstat -an | grep :80 | grep SYN_RECV'`.
    2.  Di mesin **attacker**, gunakan `hping3` untuk mensimulasikan beberapa IP yang menyerang. Gunakan skrip shell sederhana untuk menjalankan `hping3` di latar belakang dengan IP spoofing (atau dari beberapa VM jika memungkinkan).
        ```bash
        # Contoh skrip di mesin attacker
        for i in {1..10}
        do
          hping3 --rand-source --flood -p 80 -S [IP_TARGET] &
        done
        ```
    3.  **Amati:** Lihat bagaimana koneksi `SYN_RECV` di mesin target melonjak drastis. Ini mensimulasikan bagaimana botnet membanjiri sumber daya koneksi server.

---

#### **Sesi 3 & 4: Teknik & Alat Serangan (Digabung untuk Praktik)**

*   **Teknik Umum:**
    *   **SYN Flood:** Mengeksploitasi proses TCP three-way handshake dengan mengirimkan banyak paket SYN tanpa pernah menyelesaikannya, menghabiskan tabel koneksi server.
    *   **UDP Flood:** Mengirimkan sejumlah besar paket UDP ke port acak. Server akan mencoba memprosesnya dan gagal menemukan aplikasi yang mendengarkan, sehingga menghabiskan sumber daya.
    *   **HTTP Flood:** Serangan Layer 7. Menyerang layer aplikasi dengan mengirimkan banyak permintaan HTTP GET/POST yang terlihat valid (seperti pengguna asli), menyebabkan server kehabisan sumber daya untuk memproses halaman web.
    *   **DNS Amplification:** Serangan refleksi. Penyerang mengirimkan permintaan DNS ke server DNS terbuka dengan alamat IP target sebagai sumbernya. Server DNS akan merespons ke target dengan volume data yang jauh lebih besar.
*   **Alat-alat Serangan:**
    *   **LOIC (Low Orbit Ion Cannon):** Sederhana, GUI-based, sering digunakan untuk serangan yang di-koordinasikan oleh banyak orang (seperti Anonymous).
    *   **HOIC (High Orbit Ion Cannon):** Penerus LOIC, lebih kuat, dapat menyerang beberapa target sekaligus dengan booster files.
    *   **Slowloris:** Menargetkan web server dengan membuka banyak koneksi HTTP dan mengirimkannya sangat lambat, memegang koneksi tersebut terbuka dan menghabiskan thread pool server.
    *   **Hping3:** Alat analisis paket jaringan yang sangat fleksibel, bisa digunakan untuk membuat berbagai jenis serangan (SYN, UDP, ICMP flood).

**Praktik 3: Hands-on dengan Alat Serangan (Slowloris & Hping3)**

*   **Tujuan:** Peserta merasakan langsung cara kerja dan dampak dari alat serangan.
*   **Langkah-langkah:**
    1.  **Serangan Slowloris:**
        *   Di mesin **attacker**, buka terminal. Jalankan Slowloris ke arah IP mesin target di port 80.
            ```bash
            slowloris.pl -dns [IP_TARGET] -port 80
            ```
        *   Di mesin **target**, coba akses web server dari browser (dari mesin host atau VM lain). Amati bahwa loading halaman akan menjadi sangat lambat atau bahkan tidak bisa dibuka.
        *   Di mesin **target**, jalankan `netstat -an | grep :80 | grep ESTABLISHED`. Akan terlihat banyak koneksi dalam status `ESTABLISHED` yang "menggantung".
    2.  **Serangan SYN Flood dengan Hping3:**
        *   Di mesin **target**, pantau koneksi seperti pada praktik botnet.
        *   Di mesin **attacker**, jalankan perintah:
            ```bash
            hping3 --flood -p 80 -S [IP_TARGET]
            ```
        *   **Amati:** Lonjakan koneksi `SYN_RECV` di mesin target. Akses web akan terganggu.

---

#### **Sesi 5: Teknik Deteksi Serangan**

*   **Signature-based Detection:** Mencocokkan pola lalu lintas dengan basis data "tanda tangan" serangan yang sudah diketahui. Cepat tetapi tidak efektif untuk serangan baru (zero-day).
*   **Anomaly-based Detection:** Membuat profil "normal" dari lalu lintas jaringan. Apapun yang menyimpang secara signifikan dari profil ini akan ditandai sebagai anomali. Lebih baik untuk mendeteksi serangan baru.
*   **Traffic Analysis:** Memantau metrik seperti:
    *   Kenaikan tiba-tiba volume lalu lintas (bandwidth).
    *   Kenaikan jumlah paket per detik (PPS).
    *   Banyaknya koneksi dari sumber yang tidak biasa.
    *   Rasio jenis paket yang aneh (misalnya, SYN jauh lebih banyak daripada ACK).

**Praktik 4: Analisis Lalu Lintas Serangan dengan Wireshark**

*   **Tujuan:** Mengidentifikasi karakteristik serangan dari tangkapan paket (packet capture).
*   **Langkah-langkah:**
    1.  Di mesin **target** (atau mesin monitor di jaringan yang sama), mulai menangkap paket dengan Wireshark pada interface yang terhubung ke jaringan lab.
    2.  Jalankan salah satu serangan dari praktik sebelumnya (misalnya, SYN Flood).
    3.  Hentikan tangkapan dan analisis hasilnya di Wireshark:
        *   **Filter:** Gunakan filter `tcp.flags.syn==1 and tcp.flags.ack==0` untuk melihat hanya paket SYN.
        *   **Statistik:** Lihat `Statistics -> Protocol Hierarchy`. Akan terlihat dominasi protokol TCP.
        *   **Konversasi:** Lihat `Statistics -> Conversations`. Akan terlihat banyak percakapan dari berbagai IP ke port 80 target, tetapi hanya satu arah (atau tidak lengkap). Ini adalah ciri khas SYN Flood.

---

#### **Sesi 6: Alat dan Strategi Perlindungan**

*   **Tingkat Jaringan (Network-level):**
    *   **Blackholing/Null-routing:** Membuang semua lalu lintas yang menuju target. Efektif menghentikan serangan, tetapi juga memutus akses pengguna legit.
    *   **Rate Limiting:** Membatasi jumlah paket per detik dari sebuah sumber. Sederhana, tetapi bisa dibypass oleh DDoS besar.
*   **Tingkat Aplikasi (Application-level):**
    *   **Web Application Firewall (WAF):** Memfilter request HTTP yang berbahaya.
*   **Solusi Berbasis Cloud (CDN & DDoS Mitigation):**
    *   **Cloudflare, Akamai, AWS Shield:** Lalu lintas melewati jaringan mereka yang besar. Mereka memiliki sumber daya untuk "menyaring" (scrubbing) lalu lintas berbahaya dan hanya meneruskan lalu lintas yang bersih ke server asli Anda.

**Praktik 5: Implementasi Perlindungan Dasar dengan `iptables`**

*   **Tujuan:** Menunjukkan cara kerja salah satu metode mitigasi paling dasar: rate limiting.
*   **Langkah-langkah:**
    1.  Di mesin **target**, pastikan `iptables` terinstal.
    2.  Jalankan serangan SYN Flood dari mesin attacker untuk membuktikan target rentan.
    3.  Di mesin **target**, terapkan aturan `iptables` untuk membatasi laju paket SYN baru:
        ```bash
        # Menerima maksimal 1 koneksi SYN per detik dari satu sumber, dengan burst 3
        iptables -A INPUT -p tcp --syn -m limit --limit 1/s --limit-burst 3 -j ACCEPT
        
        # Jika melebihi limit, maka drop (tolak)
        iptables -A INPUT -p tcp --syn -j DROP
        ```
    4.  Jalankan kembali serangan SYN Flood dari mesin attacker.
    5.  **Amarti:** Di mesin target, jalankan `watch 'iptables -L -v -n'` untuk melihat counter paket yang di-drop. Akses web dari IP yang tidak menyerang seharusnya masih tetap bisa dilakukan. Dampak serangan akan jauh berkurang.

---

#### **Slide Penutup**

**Kesimpulan:**

*   Serangan DoS/DDoS adalah ancaman nyata yang dapat mengganggu operasional bisnis.
*   Pemahaman teknik serangan adalah kunci untuk membangun pertahanan yang efektif.
*   Deteksi dini sangat penting untuk meminimalkan dampak.
*   Perlindungan berlapis (defense-in-depth), mulai dari konfigurasi jaringan hingga menggunakan layanan cloud, adalah strategi terbaik.

**Sesi Tanya Jawab**

**Terima Kasih!**
