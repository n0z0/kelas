# MODUL PRAKTIKUM
# Network Scanning and Enumeration in Information Security
## Menggunakan Kali Linux - 8 Jam Pelajaran

---

## 📋 INFORMASI MODUL

**Durasi:** 8 Jam (480 Menit)  
**Level:** Pemula hingga Menengah  
**Platform:** Kali Linux 2024.x  
**Target Lab:** Metasploitable 2/3  
**Metode:** Hands-on Practice

---

## 🎯 TUJUAN PEMBELAJARAN

Setelah menyelesaikan praktikum ini, peserta diharapkan mampu:

1. ✅ Mengonfigurasi lab environment untuk penetration testing
2. ✅ Melakukan network reconnaissance dan host discovery
3. ✅ Menggunakan Nmap untuk scanning port dan service detection
4. ✅ Melakukan OS fingerprinting dan version detection
5. ✅ Menjalankan berbagai teknik enumeration (NetBIOS, SMB, SNMP, LDAP, DNS, SMTP)
6. ✅ Menganalisis hasil scanning dan enumeration
7. ✅ Mengidentifikasi vulnerability dari hasil enumeration
8. ✅ Membuat laporan profesional dari hasil penetration testing
9. ✅ Memahami aspek legal dan etika dalam security testing

---

## 📦 KEBUTUHAN LAB

### Hardware Minimum:
- **CPU:** Intel i5 atau AMD Ryzen 5 (dengan VT-x/AMD-V enabled)
- **RAM:** 8 GB (Recommended: 16 GB)
- **Storage:** 50 GB free space
- **Network:** Koneksi internet untuk download

### Software Requirements:
- **Hypervisor:** VirtualBox 7.x atau VMware Workstation
- **Kali Linux:** 2024.x (ISO dapat diunduh dari kali.org)
- **Target Machine:** Metasploitable 2 atau 3
- **Optional:** DVWA, bWAPP untuk web application testing

### Network Configuration:
- **Network Mode:** Host-Only atau NAT Network
- **Kali Linux IP:** 192.168.56.101 (contoh)
- **Metasploitable IP:** 192.168.56.102 (contoh)

---

## 📅 JADWAL PEMBELAJARAN 8 JAM

### **SESI 1: Setup Lab Environment (60 menit)**
- Instalasi dan konfigurasi Kali Linux
- Setup target machine (Metasploitable)
- Network configuration dan testing
- Pengenalan tools security testing

### **Break 1: 15 menit**

### **SESI 2: Network Reconnaissance & Host Discovery (90 menit)**
- Passive reconnaissance
- Active host discovery dengan Nmap
- Network mapping
- Dokumentasi hasil scanning

### **ISTIRAHAT MAKAN: 60 menit**

### **SESI 3: Port Scanning & Service Detection (75 menit)**
- TCP scanning techniques
- UDP scanning
- Service version detection
- OS fingerprinting

### **Break 2: 15 menit**

### **SESI 4: Advanced Scanning Techniques (60 menit)**
- Stealth scanning
- Firewall evasion
- NSE (Nmap Scripting Engine)
- Output formatting

### **Break 3: 15 menit**

### **SESI 5: Enumeration Techniques - Part 1 (60 menit)**
- NetBIOS enumeration
- SMB enumeration
- SNMP enumeration

### **Break 4: 15 menit**

### **SESI 6: Enumeration Techniques - Part 2 (60 menit)**
- LDAP enumeration
- DNS enumeration
- SMTP enumeration
- NTP enumeration

### **Break 5: 15 menit**

### **SESI 7: Vulnerability Assessment & Analysis (45 menit)**
- Analisis hasil scanning
- Cross-reference dengan vulnerability database
- Risk assessment
- Prioritas remediation

### **SESI 8: Report Writing & Documentation (45 menit)**
- Professional report structure
- Evidence documentation
- Recommendations
- Presentation findings

---

# SESI 1: SETUP LAB ENVIRONMENT (60 MENIT)

## 1.1 Instalasi Kali Linux di VirtualBox

### Langkah 1: Download Kali Linux
```bash
# Download dari website resmi
URL: https://www.kali.org/get-kali/

# Pilih: Kali Linux 64-bit (Installer atau Pre-built VM)
# Recommended: Pre-built VM untuk kemudahan
```

### Langkah 2: Import Kali Linux ke VirtualBox

**Jika menggunakan Pre-built VM:**
```
1. Buka VirtualBox
2. File → Import Appliance
3. Pilih file .ova yang sudah didownload
4. Klik Import
5. Tunggu proses import selesai
```

**Konfigurasi VM:**
```
- Name: Kali-Linux-Praktikum
- RAM: 4096 MB (4 GB)
- CPU: 2 cores
- Network: Host-Only Adapter (vboxnet0)
- Video Memory: 128 MB
```

### Langkah 3: First Boot Kali Linux

**Default Credentials:**
```
Username: kali
Password: kali
```

**Update System:**
```bash
# Login sebagai root atau gunakan sudo
sudo apt update
sudo apt upgrade -y
sudo apt dist-upgrade -y

# Reboot jika ada kernel update
sudo reboot
```

---

## 1.2 Setup Metasploitable 2 (Target Machine)

### Download Metasploitable 2
```
URL: https://sourceforge.net/projects/metasploitable/

File: metasploitable-linux-2.0.0.zip
```

### Import ke VirtualBox

**Langkah-langkah:**
```
1. Extract file .zip
2. Buka VirtualBox
3. New → Create Virtual Machine
4. Name: Metasploitable2
5. Type: Linux
6. Version: Ubuntu (64-bit)
7. RAM: 512 MB
8. Hard Disk: Use existing → pilih .vmdk dari extract
9. Network: Host-Only Adapter (vboxnet0) - sama dengan Kali
```

**Start Metasploitable:**
```
Username: msfadmin
Password: msfadmin
```

---

## 1.3 Network Configuration

### Konfigurasi Host-Only Network di VirtualBox

**Langkah 1: Create Host-Only Network**
```
VirtualBox → File → Host Network Manager
Klik "Create"

Network Details:
- IPv4 Address: 192.168.56.1
- IPv4 Network Mask: 255.255.255.0
- DHCP Server: Enabled (optional)
```

### Konfigurasi IP di Kali Linux

**Check current IP:**
```bash
ip addr show
# atau
ifconfig
```

**Set static IP (optional):**
```bash
# Edit network configuration
sudo nano /etc/network/interfaces

# Tambahkan:
auto eth0
iface eth0 inet static
    address 192.168.56.101
    netmask 255.255.255.0
    gateway 192.168.56.1

# Restart networking
sudo systemctl restart networking
```

### Test Connectivity

**Dari Kali Linux:**
```bash
# Check IP Metasploitable
# Login ke Metasploitable dan jalankan: ifconfig

# Ping Metasploitable dari Kali
ping -c 4 192.168.56.102

# Expected output:
# 64 bytes from 192.168.56.102: icmp_seq=1 ttl=64 time=0.xxx ms
```

---

## 1.4 Pengenalan Tools di Kali Linux

### Essential Tools untuk Lab

**1. Nmap - Network Scanner**
```bash
# Check version
nmap --version

# Basic help
nmap --help
man nmap
```

**2. Netdiscover - Network Discovery**
```bash
# Install jika belum ada
sudo apt install netdiscover -y
```

**3. Enumeration Tools**
```bash
# Install essential tools
sudo apt install enum4linux smbclient nbtscan snmp onesixtyone ldap-utils -y
```

**4. Metasploit Framework**
```bash
# Start Metasploit database
sudo msfdb init

# Start Metasploit console
msfconsole
```

**5. Wireshark - Network Analyzer**
```bash
# Start Wireshark
sudo wireshark
```

---

## 1.5 Pre-Lab Checklist

**✅ Checklist Setup:**
```
□ Kali Linux berjalan dengan baik
□ Metasploitable 2 berjalan dengan baik
□ Kedua VM dalam network yang sama (Host-Only)
□ Kali dapat ping Metasploitable
□ Metasploitable dapat ping Kali
□ Nmap terinstall dan berfungsi
□ Tools enumeration terinstall
□ Screenshot tools untuk dokumentasi ready
```

---

## 🎯 LATIHAN SESI 1

### **Latihan 1.1: Network Discovery**
```bash
# Temukan semua host di network
sudo netdiscover -r 192.168.56.0/24

# Atau gunakan Nmap
sudo nmap -sn 192.168.56.0/24
```

**Pertanyaan:**
1. Berapa banyak host yang ditemukan?
2. Apa MAC address dari Metasploitable?
3. Apa IP address dari Kali Linux Anda?

### **Latihan 1.2: Basic Connectivity Test**
```bash
# Test dengan ping
ping -c 10 192.168.56.102

# Test dengan traceroute
traceroute 192.168.56.102

# Check route
ip route show
```

**Dokumentasi:**
- Screenshot hasil netdiscover
- Screenshot hasil ping
- Catat IP address semua device

---

# SESI 2: NETWORK RECONNAISSANCE & HOST DISCOVERY (90 MENIT)

## 2.1 Passive Reconnaissance

### Konsep Passive Reconnaissance

**Definisi:**
Mengumpulkan informasi tanpa interaksi langsung dengan target.

**Tools dan Teknik:**
```bash
# 1. WHOIS Lookup
whois scanme.nmap.org

# 2. DNS Enumeration (Passive)
nslookup scanme.nmap.org
dig scanme.nmap.org

# 3. Search Engine (Google Dorking)
# Dilakukan via browser:
site:target.com
filetype:pdf site:target.com
intitle:"index of" site:target.com

# 4. Shodan (IoT Search Engine)
# Website: https://www.shodan.io
```

---

## 2.2 Active Host Discovery dengan Nmap

### List Scan (No Packets Sent)

**Konsep:**
Hanya melakukan DNS reverse lookup tanpa mengirim packet.

**Command:**
```bash
# List scan pada subnet
nmap -sL 192.168.56.0/24

# List scan pada range
nmap -sL 192.168.56.100-110
```

**Output yang Diharapkan:**
```
Starting Nmap 7.94 ( https://nmap.org )
Nmap scan report for 192.168.56.101
Nmap scan report for 192.168.56.102
...
```

---

### Ping Scan (Host Discovery)

**Berbagai Teknik Ping:**

**1. Default Ping Scan (-sn)**
```bash
# Ping sweep pada subnet
nmap -sn 192.168.56.0/24

# Simpan hasil ke file
nmap -sn 192.168.56.0/24 -oN ping_scan.txt
```

**2. ICMP Echo Ping (-PE)**
```bash
# Traditional ICMP echo request
sudo nmap -PE 192.168.56.102
```

**3. TCP SYN Ping (-PS)**
```bash
# TCP SYN pada port 80 (default)
sudo nmap -PS 192.168.56.102

# TCP SYN pada port spesifik
sudo nmap -PS22,80,443 192.168.56.102
```

**4. TCP ACK Ping (-PA)**
```bash
# TCP ACK ping
sudo nmap -PA 192.168.56.102

# Multiple ports
sudo nmap -PA22,80,443 192.168.56.102
```

**5. UDP Ping (-PU)**
```bash
# UDP ping pada port 53 (DNS)
sudo nmap -PU53 192.168.56.102
```

**6. ARP Ping (-PR)**
```bash
# ARP ping (paling akurat dalam local network)
sudo nmap -PR 192.168.56.102

# Verify dengan arp-scan
sudo arp-scan --localnet
```

---

### Disable Ping (No Ping)

**Kenapa Disable Ping?**
Host mungkin memblok ICMP tapi service tetap berjalan.

**Command:**
```bash
# Skip ping, langsung port scan
nmap -Pn 192.168.56.102

# Useful jika firewall block ICMP
```

---

## 2.3 Network Mapping

### Traceroute dengan Nmap

**Command:**
```bash
# Traceroute ke target
sudo nmap --traceroute 192.168.56.102

# Kombinasi dengan ping scan
sudo nmap -sn --traceroute 192.168.56.102
```

### Network Topology Discovery

**Identify Gateway:**
```bash
# Show routing table
ip route show

# Identify default gateway
ip route | grep default
```

**Map the Network:**
```bash
# Comprehensive network map
sudo nmap -sn -PE -PS22,80,443 -PA80 -PU53 --traceroute 192.168.56.0/24

# Output options
-oN network_map.txt      # Normal output
-oX network_map.xml      # XML output
-oG network_map.gnmap    # Grepable output
-oA network_map          # All formats
```

---

## 2.4 Dokumentasi Hasil Scanning

### Create Scan Report Directory

```bash
# Buat directory untuk hasil scan
mkdir -p ~/scans/$(date +%Y-%m-%d)
cd ~/scans/$(date +%Y-%m-%d)

# Buat subdirectory per target
mkdir -p metasploitable/{nmap,enum,screenshots}
```

### Scan dengan Multiple Output Formats

```bash
# Comprehensive scan dengan semua output format
sudo nmap -sn -PE -PS -PA -PU --traceroute 192.168.56.0/24 \
  -oA metasploitable/nmap/host_discovery
```

---

## 🎯 LATIHAN SESI 2

### **Latihan 2.1: Host Discovery Comparison**

**Task:** Bandingkan berbagai teknik host discovery

```bash
# 1. ICMP Echo Ping
sudo nmap -PE -sn 192.168.56.0/24 -oN icmp_ping.txt

# 2. TCP SYN Ping
sudo nmap -PS -sn 192.168.56.0/24 -oN tcp_syn_ping.txt

# 3. ARP Ping
sudo nmap -PR -sn 192.168.56.0/24 -oN arp_ping.txt

# 4. Kombinasi semua
sudo nmap -PE -PS -PA -PU -sn 192.168.56.0/24 -oN combined_ping.txt
```

**Pertanyaan:**
1. Teknik mana yang menemukan host paling banyak?
2. Teknik mana yang paling cepat?
3. Mengapa hasilnya bisa berbeda?

---

### **Latihan 2.2: Network Mapping**

**Task:** Buat peta lengkap network lab Anda

```bash
# Comprehensive network mapping
sudo nmap -sn -PE -PS22,80,443 -PA80 -PU53 --traceroute \
  --reason --packet-trace 192.168.56.0/24 \
  -oA network_map_complete
```

**Deliverables:**
1. List semua host yang ditemukan
2. Identifikasi MAC address masing-masing host
3. Gambar topologi network (manual sketch)
4. Screenshot hasil scanning

---

### **Latihan 2.3: Timing and Performance**

**Task:** Eksperimen dengan timing options

```bash
# T0 - Paranoid (slowest, IDS evasion)
sudo nmap -T0 -sn 192.168.56.102

# T1 - Sneaky
sudo nmap -T1 -sn 192.168.56.102

# T2 - Polite
sudo nmap -T2 -sn 192.168.56.102

# T3 - Normal (default)
sudo nmap -T3 -sn 192.168.56.102

# T4 - Aggressive (faster)
sudo nmap -T4 -sn 192.168.56.102

# T5 - Insane (fastest, may be inaccurate)
sudo nmap -T5 -sn 192.168.56.102
```

**Analisis:**
- Catat waktu eksekusi masing-masing
- Bandingkan akurasi hasil
- Kapan menggunakan timing yang berbeda?

---

# SESI 3: PORT SCANNING & SERVICE DETECTION (75 MENIT)

## 3.1 TCP Port Scanning Techniques

### TCP Connect Scan (-sT)

**Konsep:**
Full TCP three-way handshake (SYN → SYN-ACK → ACK)

**Karakteristik:**
- ✅ Tidak perlu root privileges
- ✅ Paling akurat
- ❌ Tercatat di log server
- ❌ Mudah terdeteksi

**Command:**
```bash
# Basic TCP connect scan
nmap -sT 192.168.56.102

# Scan specific ports
nmap -sT -p 22,80,443 192.168.56.102

# Scan port range
nmap -sT -p 1-1000 192.168.56.102

# Scan all ports
nmap -sT -p- 192.168.56.102
```

---

### TCP SYN Scan (-sS) - Stealth Scan

**Konsep:**
Half-open scan (SYN → SYN-ACK → RST)

**Karakteristik:**
- ✅ Stealth (tidak tercatat di application log)
- ✅ Lebih cepat dari -sT
- ✅ Default scan type (jika root)
- ⚠️ Perlu root privileges

**Command:**
```bash
# TCP SYN scan (stealth)
sudo nmap -sS 192.168.56.102

# Top 1000 ports (default)
sudo nmap -sS 192.168.56.102

# All ports
sudo nmap -sS -p- 192.168.56.102

# Fast scan (top 100 ports)
sudo nmap -sS -F 192.168.56.102
```

---

### Advanced TCP Scan Types

**1. TCP FIN Scan (-sF)**
```bash
# FIN scan (firewall evasion)
sudo nmap -sF 192.168.56.102
```

**2. TCP NULL Scan (-sN)**
```bash
# NULL scan (no flags set)
sudo nmap -sN 192.168.56.102
```

**3. TCP Xmas Scan (-sX)**
```bash
# Xmas scan (FIN, PSH, URG flags)
sudo nmap -sX 192.168.56.102
```

**4. TCP ACK Scan (-sA)**
```bash
# ACK scan (firewall rule detection)
sudo nmap -sA 192.168.56.102
```

**5. TCP Window Scan (-sW)**
```bash
# Window scan (check window size)
sudo nmap -sW 192.168.56.102
```

---

## 3.2 UDP Port Scanning

### UDP Scan (-sU)

**Konsep:**
UDP adalah connectionless protocol, lebih sulit untuk scan.

**Karakteristik:**
- ⚠️ Sangat lambat
- ⚠️ Kurang akurat
- ✅ Penting untuk service seperti DNS, SNMP, DHCP

**Command:**
```bash
# UDP scan pada common ports
sudo nmap -sU --top-ports 20 192.168.56.102

# UDP scan specific ports
sudo nmap -sU -p 53,161,162 192.168.56.102

# Combined TCP and UDP scan
sudo nmap -sS -sU -p T:80,443,U:53,161 192.168.56.102
```

**Port States pada UDP:**
- **open:** UDP response diterima
- **closed:** ICMP port unreachable diterima
- **open|filtered:** No response (timeout)

---

## 3.3 Service Version Detection

### Version Detection (-sV)

**Command:**
```bash
# Service version detection
nmap -sV 192.168.56.102

# Intensity level (0-9)
nmap -sV --version-intensity 5 192.168.56.102

# Light version detection (faster)
nmap -sV --version-light 192.168.56.102

# Aggressive version detection (slower, more accurate)
nmap -sV --version-all 192.168.56.102
```

**Example Output:**
```
PORT     STATE SERVICE     VERSION
21/tcp   open  ftp         vsftpd 2.3.4
22/tcp   open  ssh         OpenSSH 4.7p1 Debian 8ubuntu1
80/tcp   open  http        Apache httpd 2.2.8 ((Ubuntu) DAV/2)
3306/tcp open  mysql       MySQL 5.0.51a-3ubuntu5
```

---

## 3.4 OS Fingerprinting

### OS Detection (-O)

**Command:**
```bash
# OS detection
sudo nmap -O 192.168.56.102

# Aggressive OS detection
sudo nmap -O --osscan-guess 192.168.56.102

# Limit OS detection (faster)
sudo nmap -O --osscan-limit 192.168.56.102
```

**Example Output:**
```
Running: Linux 2.6.X
OS CPE: cpe:/o:linux:linux_kernel:2.6
OS details: Linux 2.6.9 - 2.6.33
```

---

## 3.5 Comprehensive Scanning

### Aggressive Scan (-A)

**Command:**
```bash
# Aggressive scan (OS + Version + Script + Traceroute)
sudo nmap -A 192.168.56.102

# Aggressive scan specific ports
sudo nmap -A -p 21,22,80 192.168.56.102
```

**What -A includes:**
- `-O` (OS detection)
- `-sV` (Version detection)
- `-sC` (Default scripts)
- `--traceroute` (Traceroute)

---

## 🎯 LATIHAN SESI 3

### **Latihan 3.1: Port Scanning Comparison**

**Task:** Bandingkan hasil berbagai teknik scanning

```bash
# Create results directory
mkdir -p ~/scans/port_scanning

# 1. TCP Connect Scan
nmap -sT -p 1-1000 192.168.56.102 -oN ~/scans/port_scanning/tcp_connect.txt

# 2. TCP SYN Scan
sudo nmap -sS -p 1-1000 192.168.56.102 -oN ~/scans/port_scanning/tcp_syn.txt

# 3. UDP Scan (top 100)
sudo nmap -sU --top-ports 100 192.168.56.102 -oN ~/scans/port_scanning/udp.txt
```

**Analisis:**
1. Berapa port terbuka yang ditemukan di TCP?
2. Berapa port terbuka yang ditemukan di UDP?
3. Bandingkan waktu eksekusi
4. Service apa saja yang berjalan?

---

### **Latihan 3.2: Service and OS Detection**

**Task:** Identifikasi service dan OS pada Metasploitable

```bash
# Comprehensive scan dengan output
sudo nmap -sV -O -sC -p- 192.168.56.102 \
  -oA ~/scans/port_scanning/comprehensive_scan

# Analyze output
cat ~/scans/port_scanning/comprehensive_scan.nmap
```

**Dokumentasi:**
1. List semua service beserta versinya
2. Identifikasi OS yang terdeteksi
3. Temukan service yang potentially vulnerable
4. Screenshot hasil scanning

---

### **Latihan 3.3: Top 10 Vulnerable Services**

**Task:** Scan dan identifikasi top 10 service yang paling umum

```bash
# Scan common ports
sudo nmap -sV --top-ports 20 192.168.56.102 -oN top_services.txt

# Detail scan pada port yang terbuka
# Misal jika port 21 (FTP) terbuka:
sudo nmap -sV -sC -p 21 192.168.56.102
```

**List Service yang Harus Dicheck:**
1. FTP (21)
2. SSH (22)
3. Telnet (23)
4. SMTP (25)
5. DNS (53)
6. HTTP (80/8080)
7. POP3 (110)
8. NetBIOS (139)
9. SMB (445)
10. MySQL (3306)

---

# SESI 4: ADVANCED SCANNING TECHNIQUES (60 MENIT)

## 4.1 Stealth Scanning Techniques

### Why Stealth Scanning?

**Tujuan:**
- Menghindari deteksi IDS/IPS
- Bypass firewall rules
- Reduce noise dalam logs

### NULL, FIN, and Xmas Scans

**1. NULL Scan (-sN)**
```bash
# No TCP flags set
sudo nmap -sN 192.168.56.102

# With verbose
sudo nmap -sN -v 192.168.56.102
```

**Bagaimana Kerjanya:**
- Packet tanpa flags
- Open port: No response
- Closed port: RST response

---

**2. FIN Scan (-sF)**
```bash
# Only FIN flag set
sudo nmap -sF 192.168.56.102
```

**Bagaimana Kerjanya:**
- Packet dengan FIN flag
- Open port: No response
- Closed port: RST response

---

**3. Xmas Scan (-sX)**
```bash
# FIN, PSH, URG flags set (lights up like Christmas tree)
sudo nmap -sX 192.168.56.102
```

---

## 4.2 Firewall Evasion Techniques

### Fragment Packets (-f)

**Command:**
```bash
# Fragment packets into 8 bytes
sudo nmap -f 192.168.56.102

# Fragment into 16 bytes
sudo nmap -ff 192.168.56.102

# Custom MTU
sudo nmap --mtu 24 192.168.56.102
```

---

### Decoy Scanning (-D)

**Command:**
```bash
# Use decoys
sudo nmap -D RND:10 192.168.56.102

# Specific decoys
sudo nmap -D 192.168.56.50,192.168.56.51,ME,192.168.56.52 192.168.56.102
```

**Konsep:**
Target melihat scan dari multiple IP addresses.

---

### Idle Scan (Zombie Scan)

**Command:**
```bash
# Find zombie host first
sudo nmap -O -v 192.168.56.0/24

# Use zombie for scanning
sudo nmap -sI <zombie_ip> 192.168.56.102
```

---

### Source Port Spoofing

**Command:**
```bash
# Spoof source port (common: 53, 80, 443)
sudo nmap --source-port 53 192.168.56.102

# Short form
sudo nmap -g 53 192.168.56.102
```

---

### Timing Templates

**Six Timing Templates:**
```bash
# T0 - Paranoid (IDS evasion, very slow)
sudo nmap -T0 -p 22 192.168.56.102

# T1 - Sneaky
sudo nmap -T1 -p 22 192.168.56.102

# T2 - Polite (won't crash target)
sudo nmap -T2 192.168.56.102

# T3 - Normal (default)
sudo nmap -T3 192.168.56.102

# T4 - Aggressive (fast, assumes good network)
sudo nmap -T4 192.168.56.102

# T5 - Insane (extremely fast, may miss ports)
sudo nmap -T5 192.168.56.102
```

---

## 4.3 Nmap Scripting Engine (NSE)

### Introduction to NSE

**Categories:**
- auth: Authentication scripts
- broadcast: Network broadcast scripts
- brute: Brute force attacks
- default: Default scripts (-sC)
- discovery: Network/service discovery
- dos: Denial of service
- exploit: Exploit vulnerabilities
- fuzzer: Fuzzing scripts
- intrusive: Intrusive scripts
- malware: Malware detection
- safe: Safe scripts
- version: Version detection enhancement
- vuln: Vulnerability detection

---

### Using NSE Scripts

**List Available Scripts:**
```bash
# List all scripts
ls /usr/share/nmap/scripts/

# Search scripts by keyword
ls /usr/share/nmap/scripts/ | grep ftp

# Script help
nmap --script-help ftp-anon
```

---

**Run Specific Scripts:**
```bash
# Single script
nmap --script=ftp-anon 192.168.56.102

# Multiple scripts
nmap --script=ftp-anon,ftp-vsftpd-backdoor -p 21 192.168.56.102

# Category
nmap --script=vuln 192.168.56.102

# Multiple categories
nmap --script="default,safe" 192.168.56.102
```

---

**Useful NSE Scripts Examples:**

**1. Vulnerability Scanning:**
```bash
# Scan for vulnerabilities
sudo nmap -sV --script=vuln 192.168.56.102

# Specific CVE
sudo nmap --script=smb-vuln-ms17-010 -p 445 192.168.56.102
```

**2. Brute Force:**
```bash
# FTP brute force
nmap --script=ftp-brute -p 21 192.168.56.102

# SSH brute force
nmap --script=ssh-brute -p 22 192.168.56.102

# Custom wordlist
nmap --script=ftp-brute --script-args userdb=/path/to/users.txt,passdb=/path/to/passwords.txt 192.168.56.102
```

**3. Service Specific:**
```bash
# HTTP enumeration
nmap --script=http-enum -p 80 192.168.56.102

# SMB enumeration
nmap --script=smb-enum-shares,smb-enum-users -p 445 192.168.56.102

# DNS zone transfer
nmap --script=dns-zone-transfer --script-args dns-zone-transfer.domain=example.com -p 53 <dns_server>
```

---

## 4.4 Output Formats and Reporting

### Output Options

**1. Normal Output (-oN)**
```bash
nmap -oN scan_results.txt 192.168.56.102
```

**2. XML Output (-oX)**
```bash
nmap -oX scan_results.xml 192.168.56.102
```

**3. Grepable Output (-oG)**
```bash
nmap -oG scan_results.gnmap 192.168.56.102
```

**4. All Formats (-oA)**
```bash
# Creates .nmap, .xml, and .gnmap
nmap -oA scan_results 192.168.56.102
```

---

### Advanced Output Options

**Append to File:**
```bash
nmap --append-output -oN scan_results.txt 192.168.56.102
```

**Verbose Output:**
```bash
# Verbose level 1
nmap -v 192.168.56.102

# Verbose level 2
nmap -vv 192.168.56.102

# Debug mode
nmap -d 192.168.56.102
```

---

## 🎯 LATIHAN SESI 4

### **Latihan 4.1: Stealth Scanning**

**Task:** Praktik berbagai stealth technique

```bash
# Directory untuk hasil
mkdir -p ~/scans/stealth_scanning

# 1. TCP SYN Scan (baseline)
sudo nmap -sS -p 21,22,80 192.168.56.102 -oN ~/scans/stealth_scanning/syn_scan.txt

# 2. NULL Scan
sudo nmap -sN -p 21,22,80 192.168.56.102 -oN ~/scans/stealth_scanning/null_scan.txt

# 3. FIN Scan
sudo nmap -sF -p 21,22,80 192.168.56.102 -oN ~/scans/stealth_scanning/fin_scan.txt

# 4. Xmas Scan
sudo nmap -sX -p 21,22,80 192.168.56.102 -oN ~/scans/stealth_scanning/xmas_scan.txt
```

**Analisis:**
1. Bandingkan hasil keempat scan
2. Apakah ada perbedaan dalam deteksi port?
3. Teknik mana yang paling stealth?

---

### **Latihan 4.2: Firewall Evasion**

**Task:** Bypass potential firewall rules

```bash
# 1. Normal scan (baseline)
sudo nmap -sS -p 80 192.168.56.102

# 2. Fragment packets
sudo nmap -sS -f -p 80 192.168.56.102

# 3. Use decoys
sudo nmap -sS -D RND:5 -p 80 192.168.56.102

# 4. Source port spoofing
sudo nmap -sS --source-port 53 -p 80 192.168.56.102

# 5. Slow timing
sudo nmap -sS -T1 -p 80 192.168.56.102
```

---

### **Latihan 4.3: NSE Scripts Practice**

**Task:** Gunakan NSE untuk vulnerability scanning

```bash
# 1. Default scripts
sudo nmap -sC -sV 192.168.56.102 -oN default_scripts.txt

# 2. Vulnerability scanning
sudo nmap -sV --script=vuln 192.168.56.102 -oN vuln_scan.txt

# 3. FTP specific scripts
sudo nmap --script="ftp-*" -p 21 192.168.56.102 -oN ftp_scripts.txt

# 4. SMB vulnerability check
sudo nmap --script=smb-vuln* -p 445 192.168.56.102 -oN smb_vuln.txt

# 5. HTTP enumeration
sudo nmap --script=http-enum -p 80 192.168.56.102 -oN http_enum.txt
```

**Dokumentasi:**
1. List vulnerability yang ditemukan
2. Severity rating masing-masing vulnerability
3. Recommendations untuk remediation

---

# SESI 5: ENUMERATION TECHNIQUES - PART 1 (60 MENIT)

## 5.1 NetBIOS Enumeration

### Introduction to NetBIOS

**NetBIOS (Network Basic Input Output System):**
- Port: 137 (Name Service), 138 (Datagram), 139 (Session)
- Windows network sharing protocol
- Provides information about:
  - Computer names
  - Domain/Workgroup
  - Shared resources
  - User accounts

---

### NetBIOS Enumeration Tools

**1. Nmap NSE Scripts:**
```bash
# NetBIOS enumeration
sudo nmap -sU -sS --script nbstat.nse -p 137,139 192.168.56.102

# SMB OS discovery
sudo nmap --script smb-os-discovery -p 445 192.168.56.102
```

---

**2. nbtscan:**
```bash
# Install if not present
sudo apt install nbtscan -y

# Scan single host
nbtscan 192.168.56.102

# Scan network range
nbtscan 192.168.56.0/24

# Verbose output
nbtscan -v 192.168.56.102
```

---

**3. nmblookup (Samba tool):**
```bash
# Lookup NetBIOS name
nmblookup -A 192.168.56.102

# Find master browser
nmblookup -M -- -

# Find workgroup
nmblookup -W
```

---

**4. enum4linux:**
```bash
# Comprehensive NetBIOS/SMB enumeration
enum4linux 192.168.56.102

# Verbose mode
enum4linux -v 192.168.56.102

# All information
enum4linux -a 192.168.56.102

# Specific options:
enum4linux -U 192.168.56.102  # Users
enum4linux -S 192.168.56.102  # Shares
enum4linux -G 192.168.56.102  # Groups
enum4linux -P 192.168.56.102  # Password policy
```

---

## 5.2 SMB Enumeration

### Introduction to SMB

**SMB (Server Message Block):**
- Port: 445 (Direct SMB), 139 (SMB over NetBIOS)
- File/printer sharing protocol
- Critical for Windows networks
- Target umum untuk exploitation

---

### SMB Enumeration dengan Nmap

**1. SMB Version Detection:**
```bash
# Detect SMB version
sudo nmap -p 445 --script smb-protocols 192.168.56.102
```

**2. SMB Security Mode:**
```bash
# Check SMB security mode
sudo nmap -p 445 --script smb-security-mode 192.168.56.102
```

**3. SMB Shares Enumeration:**
```bash
# Enumerate shares
sudo nmap -p 445 --script smb-enum-shares 192.168.56.102

# With credentials
sudo nmap -p 445 --script smb-enum-shares --script-args smbusername=guest,smbpassword= 192.168.56.102
```

**4. SMB Users Enumeration:**
```bash
# Enumerate users
sudo nmap -p 445 --script smb-enum-users 192.168.56.102
```

**5. SMB Vulnerabilities:**
```bash
# Check for known SMB vulnerabilities
sudo nmap -p 445 --script smb-vuln* 192.168.56.102

# Specific: EternalBlue (MS17-010)
sudo nmap -p 445 --script smb-vuln-ms17-010 192.168.56.102
```

---

### SMB Enumeration dengan smbclient

**Install smbclient:**
```bash
sudo apt install smbclient -y
```

**1. List Shares:**
```bash
# Anonymous access
smbclient -L //192.168.56.102 -N

# With credentials
smbclient -L //192.168.56.102 -U username
```

**2. Connect to Share:**
```bash
# Connect to share
smbclient //192.168.56.102/tmp -N

# SMB commands inside:
smb: \> ls          # List files
smb: \> get file    # Download file
smb: \> put file    # Upload file
smb: \> help        # Show commands
```

---

### SMB Enumeration dengan smbmap

**Install smbmap:**
```bash
sudo apt install smbmap -y
```

**Commands:**
```bash
# List shares with permissions
smbmap -H 192.168.56.102

# List contents of all shares
smbmap -H 192.168.56.102 -R

# With credentials
smbmap -H 192.168.56.102 -u username -p password

# Execute command
smbmap -H 192.168.56.102 -u username -p password -x 'ipconfig'
```

---

## 5.3 SNMP Enumeration

### Introduction to SNMP

**SNMP (Simple Network Management Protocol):**
- Port: 161 (agent), 162 (trap)
- Protocol: UDP
- Network device management
- Community strings: "public" (read), "private" (read-write)

**Information Gathered:**
- Device information
- Network interfaces
- Routing tables
- Running processes
- Installed software
- Open ports

---

### SNMP Enumeration Tools

**1. snmpwalk:**
```bash
# Basic snmpwalk (SNMPv1)
snmpwalk -v1 -c public 192.168.56.102

# SNMPv2c
snmpwalk -v2c -c public 192.168.56.102

# System information
snmpwalk -v2c -c public 192.168.56.102 system

# Running processes
snmpwalk -v2c -c public 192.168.56.102 hrSWRunName

# Installed software
snmpwalk -v2c -c public 192.168.56.102 hrSWInstalledName

# User accounts
snmpwalk -v2c -c public 192.168.56.102 1.3.6.1.4.1.77.1.2.25

# Network interfaces
snmpwalk -v2c -c public 192.168.56.102 interfaces
```

---

**2. snmp-check:**
```bash
# Install if not present
sudo apt install snmp-check -y

# Comprehensive SNMP enumeration
snmp-check 192.168.56.102

# With custom community string
snmp-check -c private 192.168.56.102
```

---

**3. onesixtyone (SNMP scanner):**
```bash
# Install
sudo apt install onesixtyone -y

# Scan with default community strings
onesixtyone 192.168.56.102

# Scan with custom wordlist
onesixtyone -c /usr/share/wordlists/SecLists/Discovery/SNMP/common-snmp-community-strings.txt 192.168.56.102

# Scan network range
onesixtyone -c public 192.168.56.0/24
```

---

**4. Nmap SNMP Scripts:**
```bash
# SNMP info
sudo nmap -sU -p 161 --script=snmp-info 192.168.56.102

# SNMP processes
sudo nmap -sU -p 161 --script=snmp-processes 192.168.56.102

# SNMP interfaces
sudo nmap -sU -p 161 --script=snmp-interfaces 192.168.56.102

# SNMP brute force community strings
sudo nmap -sU -p 161 --script=snmp-brute 192.168.56.102
```

---

## 🎯 LATIHAN SESI 5

### **Latihan 5.1: NetBIOS/SMB Full Enumeration**

**Task:** Lakukan enumeration lengkap pada Metasploitable

```bash
# Create directory
mkdir -p ~/scans/enumeration/netbios_smb

# 1. NetBIOS scan dengan nbtscan
nbtscan 192.168.56.102 > ~/scans/enumeration/netbios_smb/nbtscan.txt

# 2. SMB enumeration dengan Nmap
sudo nmap -p 139,445 --script "smb-enum-*,smb-vuln-*,smb-os-discovery" 192.168.56.102 \
  -oN ~/scans/enumeration/netbios_smb/nmap_smb.txt

# 3. Comprehensive enum4linux
enum4linux -a 192.168.56.102 > ~/scans/enumeration/netbios_smb/enum4linux.txt

# 4. SMB shares dengan smbclient
smbclient -L //192.168.56.102 -N > ~/scans/enumeration/netbios_smb/smbclient.txt

# 5. SMB shares dengan smbmap
smbmap -H 192.168.56.102 -R > ~/scans/enumeration/netbios_smb/smbmap.txt
```

**Analisis:**
1. List semua users yang ditemukan
2. List semua shares yang accessible
3. Identifikasi vulnerability SMB yang ada
4. Screenshot masing-masing tool

---

### **Latihan 5.2: SNMP Enumeration**

**Task:** Enumerate SNMP service (jika tersedia)

```bash
# Create directory
mkdir -p ~/scans/enumeration/snmp

# 1. Check if SNMP is running
sudo nmap -sU -p 161 192.168.56.102

# 2. SNMP info with Nmap
sudo nmap -sU -p 161 --script=snmp-info 192.168.56.102 \
  -oN ~/scans/enumeration/snmp/nmap_snmp_info.txt

# 3. snmpwalk system information
snmpwalk -v2c -c public 192.168.56.102 system \
  > ~/scans/enumeration/snmp/snmpwalk_system.txt

# 4. snmp-check comprehensive
snmp-check 192.168.56.102 > ~/scans/enumeration/snmp/snmp_check.txt

# 5. Brute force community strings
onesixtyone -c /usr/share/wordlists/SecLists/Discovery/SNMP/common-snmp-community-strings.txt 192.168.56.102 \
  > ~/scans/enumeration/snmp/community_brute.txt
```

**Note:** Jika SNMP tidak tersedia di Metasploitable 2, dokumentasikan command yang akan digunakan.

---

### **Latihan 5.3: Compare Enumeration Tools**

**Task:** Bandingkan hasil dari berbagai tools

**Create Comparison Table:**
```
┌─────────────┬──────────┬──────────┬──────────┬─────────┐
│ Information │ nbtscan  │ enum4linux│ smbclient│ smbmap │
├─────────────┼──────────┼──────────┼──────────┼─────────┤
│ Users       │          │     ✓    │          │         │
│ Shares      │          │     ✓    │    ✓     │    ✓    │
│ Groups      │          │     ✓    │          │         │
│ OS Info     │    ✓     │     ✓    │          │         │
│ Domain      │    ✓     │     ✓    │          │         │
└─────────────┴──────────┴──────────┴──────────┴─────────┘
```

**Deliverables:**
1. Tabel comparison lengkap
2. Rekomendasi tool terbaik untuk masing-masing task
3. Dokumentasi lengkap hasil enumeration

---

# SESI 6: ENUMERATION TECHNIQUES - PART 2 (60 MENIT)

## 6.1 LDAP Enumeration

### Introduction to LDAP

**LDAP (Lightweight Directory Access Protocol):**
- Port: 389 (LDAP), 636 (LDAPS)
- Directory service protocol
- Active Directory menggunakan LDAP
- Information: users, groups, computers, organizational structure

---

### LDAP Enumeration Tools

**1. Nmap LDAP Scripts:**
```bash
# LDAP root DSE
sudo nmap -p 389 --script ldap-rootdse 192.168.56.102

# LDAP search
sudo nmap -p 389 --script ldap-search --script-args 'ldap.username="cn=admin,dc=example,dc=com",ldap.password=password' 192.168.56.102

# LDAP brute force
sudo nmap -p 389 --script ldap-brute 192.168.56.102
```

---

**2. ldapsearch:**
```bash
# Anonymous bind (if allowed)
ldapsearch -x -h 192.168.56.102 -b "dc=example,dc=com"

# With credentials
ldapsearch -x -h 192.168.56.102 -D "cn=admin,dc=example,dc=com" -W -b "dc=example,dc=com"

# Search for users
ldapsearch -x -h 192.168.56.102 -b "dc=example,dc=com" "(objectClass=person)"

# Search for groups
ldapsearch -x -h 192.168.56.102 -b "dc=example,dc=com" "(objectClass=group)"

# Get all attributes
ldapsearch -x -h 192.168.56.102 -b "dc=example,dc=com" "*"
```

---

**3. ldapdomaindump:**
```bash
# Install
pip3 install ldapdomaindump

# Dump domain information
ldapdomaindump -u 'DOMAIN\user' -p password 192.168.56.102 -o ~/ldap_dump/
```

---

## 6.2 DNS Enumeration

### Introduction to DNS

**DNS (Domain Name System):**
- Port: 53 (TCP and UDP)
- Translates domain names to IP addresses
- Information: subdomains, mail servers, name servers

---

### DNS Enumeration Tools

**1. nslookup:**
```bash
# Basic lookup
nslookup example.com

# Specify DNS server
nslookup example.com 8.8.8.8

# Query specific record type
nslookup -type=A example.com
nslookup -type=MX example.com
nslookup -type=NS example.com
nslookup -type=TXT example.com
```

---

**2. dig (Domain Information Groper):**
```bash
# Basic query
dig example.com

# Query specific record
dig example.com A
dig example.com MX
dig example.com NS
dig example.com TXT

# Zone transfer attempt
dig axfr @ns1.example.com example.com

# Reverse DNS lookup
dig -x 192.168.56.102

# Trace DNS path
dig +trace example.com
```

---

**3. host:**
```bash
# Simple lookup
host example.com

# Verbose output
host -v example.com

# Zone transfer
host -l example.com ns1.example.com
```

---

**4. dnsenum:**
```bash
# Install
sudo apt install dnsenum -y

# Basic enumeration
dnsenum example.com

# With zone transfer
dnsenum --enum example.com

# Brute force subdomains
dnsenum --subfile /usr/share/wordlists/subdomains.txt example.com
```

---

**5. dnsrecon:**
```bash
# Install
sudo apt install dnsrecon -y

# Standard enumeration
dnsrecon -d example.com

# Zone transfer
dnsrecon -d example.com -t axfr

# Brute force
dnsrecon -d example.com -t brt -D /usr/share/wordlists/dns-subdomains.txt

# Reverse lookup
dnsrecon -r 192.168.56.0/24
```

---

**6. Nmap DNS Scripts:**
```bash
# DNS brute force
sudo nmap --script dns-brute example.com

# DNS zone transfer
sudo nmap --script dns-zone-transfer --script-args dns-zone-transfer.domain=example.com -p 53 <dns_server>

# DNS service discovery
sudo nmap --script dns-service-discovery -p 53 192.168.56.102
```

---

## 6.3 SMTP Enumeration

### Introduction to SMTP

**SMTP (Simple Mail Transfer Protocol):**
- Port: 25 (SMTP), 587 (Submission), 465 (SMTPS)
- Email transmission protocol
- Commands: VRFY (verify), EXPN (expand), RCPT TO

**Information Gathered:**
- Valid email addresses
- Username enumeration
- Mail server version

---

### SMTP Enumeration Tools

**1. Manual SMTP Commands:**
```bash
# Connect to SMTP
telnet 192.168.56.102 25

# Or with netcat
nc 192.168.56.102 25

# SMTP commands:
HELO attacker.com
VRFY root
VRFY admin
EXPN root
MAIL FROM: test@example.com
RCPT TO: root@target.com
QUIT
```

---

**2. smtp-user-enum:**
```bash
# Install
sudo apt install smtp-user-enum -y

# VRFY mode
smtp-user-enum -M VRFY -U /usr/share/wordlists/usernames.txt -t 192.168.56.102

# EXPN mode
smtp-user-enum -M EXPN -U /usr/share/wordlists/usernames.txt -t 192.168.56.102

# RCPT mode
smtp-user-enum -M RCPT -U /usr/share/wordlists/usernames.txt -t 192.168.56.102

# Single user
smtp-user-enum -M VRFY -u root -t 192.168.56.102
```

---

**3. Nmap SMTP Scripts:**
```bash
# SMTP commands enumeration
sudo nmap -p 25 --script smtp-commands 192.168.56.102

# SMTP user enumeration
sudo nmap -p 25 --script smtp-enum-users --script-args smtp-enum-users.methods={VRFY,EXPN,RCPT} 192.168.56.102

# SMTP open relay
sudo nmap -p 25 --script smtp-open-relay 192.168.56.102

# SMTP NTLM info
sudo nmap -p 25 --script smtp-ntlm-info 192.168.56.102
```

---

**4. Metasploit SMTP Modules:**
```bash
# Start Metasploit
msfconsole

# SMTP version scanner
use auxiliary/scanner/smtp/smtp_version
set RHOSTS 192.168.56.102
run

# SMTP user enumeration
use auxiliary/scanner/smtp/smtp_enum
set RHOSTS 192.168.56.102
set USER_FILE /usr/share/wordlists/metasploit/unix_users.txt
run
```

---

## 6.4 NTP Enumeration

### Introduction to NTP

**NTP (Network Time Protocol):**
- Port: 123 (UDP)
- Time synchronization protocol
- Information: connected hosts, network topology

---

### NTP Enumeration Tools

**1. ntpq:**
```bash
# Query NTP server
ntpq -c readlist 192.168.56.102
ntpq -c peers 192.168.56.102
ntpq -c associations 192.168.56.102

# Monlist (list of recent clients)
ntpdc -n -c monlist 192.168.56.102
```

---

**2. ntpdate:**
```bash
# Query NTP server time
ntpdate -q 192.168.56.102
```

---

**3. Nmap NTP Scripts:**
```bash
# NTP info
sudo nmap -sU -p 123 --script ntp-info 192.168.56.102

# NTP monlist (if enabled, shows recent clients)
sudo nmap -sU -p 123 --script ntp-monlist 192.168.56.102
```

---

## 🎯 LATIHAN SESI 6

### **Latihan 6.1: LDAP Enumeration**

**Task:** Enumerate LDAP service (simulasi)

```bash
# Create directory
mkdir -p ~/scans/enumeration/ldap

# 1. Check LDAP port
sudo nmap -p 389,636 192.168.56.102

# 2. LDAP root DSE
sudo nmap -p 389 --script ldap-rootdse 192.168.56.102 \
  -oN ~/scans/enumeration/ldap/ldap_rootdse.txt

# 3. Anonymous bind attempt
ldapsearch -x -h 192.168.56.102 -b "" -s base \
  > ~/scans/enumeration/ldap/ldap_anonymous.txt 2>&1
```

**Note:** Jika LDAP tidak tersedia, dokumentasikan command yang akan digunakan untuk real Active Directory environment.

---

### **Latihan 6.2: DNS Enumeration**

**Task:** Enumerate DNS untuk domain scanme.nmap.org

```bash
# Create directory
mkdir -p ~/scans/enumeration/dns

# 1. Basic lookup
dig scanme.nmap.org ANY > ~/scans/enumeration/dns/dig_any.txt

# 2. Specific records
dig scanme.nmap.org A > ~/scans/enumeration/dns/dig_a.txt
dig scanme.nmap.org MX > ~/scans/enumeration/dns/dig_mx.txt
dig scanme.nmap.org NS > ~/scans/enumeration/dns/dig_ns.txt

# 3. Zone transfer attempt (likely will fail)
dig axfr @ns1.scanme.nmap.org scanme.nmap.org \
  > ~/scans/enumeration/dns/zone_transfer.txt 2>&1

# 4. DNS enumeration dengan dnsrecon
dnsrecon -d scanme.nmap.org -t std \
  > ~/scans/enumeration/dns/dnsrecon.txt

# 5. Subdomain brute force (small wordlist)
dnsrecon -d scanme.nmap.org -t brt -D /usr/share/wordlists/dnsmap.txt \
  > ~/scans/enumeration/dns/subdomain_brute.txt
```

---

### **Latihan 6.3: SMTP Enumeration**

**Task:** Enumerate SMTP service di Metasploitable

```bash
# Create directory
mkdir -p ~/scans/enumeration/smtp

# 1. Check SMTP port
sudo nmap -p 25 192.168.56.102

# 2. SMTP commands
sudo nmap -p 25 --script smtp-commands 192.168.56.102 \
  -oN ~/scans/enumeration/smtp/smtp_commands.txt

# 3. SMTP user enumeration
sudo nmap -p 25 --script smtp-enum-users 192.168.56.102 \
  -oN ~/scans/enumeration/smtp/smtp_users_nmap.txt

# 4. smtp-user-enum dengan wordlist
smtp-user-enum -M VRFY -U /usr/share/wordlists/metasploit/unix_users.txt -t 192.168.56.102 \
  > ~/scans/enumeration/smtp/smtp_user_enum.txt

# 5. Manual enumeration dengan netcat
echo -e "HELO attacker.com\nVRFY root\nVRFY admin\nQUIT" | nc 192.168.56.102 25 \
  > ~/scans/enumeration/smtp/manual_smtp.txt
```

---

### **Latihan 6.4: Comprehensive Enumeration Report**

**Task:** Compile semua hasil enumeration

```bash
# Create comprehensive report directory
mkdir -p ~/scans/final_report

# Copy all enumeration results
cp -r ~/scans/enumeration/* ~/scans/final_report/

# Create summary file
cat > ~/scans/final_report/ENUMERATION_SUMMARY.md << 'EOF'
# Enumeration Report - Metasploitable 2

## Target Information
- IP Address: 192.168.56.102
- Hostname: metasploitable
- OS: Linux 2.6.x

## Services Enumerated

### 1. NetBIOS/SMB (Ports 139, 445)
- Users found: [list users]
- Shares accessible: [list shares]
- Vulnerabilities: [list vulnerabilities]

### 2. SNMP (Port 161)
- Community strings: [found strings]
- System information: [details]
- Running processes: [key processes]

### 3. DNS (Port 53)
- Name servers: [list]
- Subdomains found: [list]
- Zone transfer: [status]

### 4. SMTP (Port 25)
- Valid users: [list]
- Mail server version: [version]
- Open relay: [yes/no]

### 5. LDAP (Port 389)
- Base DN: [details]
- Users: [count]
- Groups: [count]

## Recommendations
1. [recommendation 1]
2. [recommendation 2]
...

EOF
```

---

# SESI 7: VULNERABILITY ASSESSMENT & ANALYSIS (45 MENIT)

## 7.1 Analisis Hasil Scanning

### Organize Scan Results

**Create Analysis Workspace:**
```bash
# Create directory structure
mkdir -p ~/vulnerability_assessment/{raw_scans,analysis,reports,evidence}

# Copy all scan results
cp -r ~/scans/* ~/vulnerability_assessment/raw_scans/
```

---

### Extract Key Information

**1. Extract Open Ports:**
```bash
# From Nmap output
grep "open" ~/vulnerability_assessment/raw_scans/port_scanning/*.nmap > \
  ~/vulnerability_assessment/analysis/open_ports.txt

# Using grep patterns
cat ~/vulnerability_assessment/raw_scans/port_scanning/*.nmap | \
  grep -E "^[0-9]+/tcp.*open" > \
  ~/vulnerability_assessment/analysis/tcp_services.txt
```

---

**2. Extract Service Versions:**
```bash
# Create service inventory
cat ~/vulnerability_assessment/raw_scans/port_scanning/comprehensive_scan.nmap | \
  grep "open" | \
  awk '{print $1, $3, $4, $5, $6}' > \
  ~/vulnerability_assessment/analysis/service_inventory.txt
```

---

**3. Extract Vulnerabilities from NSE:**
```bash
# Extract vulnerability findings
grep -i "vuln" ~/vulnerability_assessment/raw_scans/**/*.txt > \
  ~/vulnerability_assessment/analysis/vulnerabilities_found.txt
```

---

## 7.2 Cross-Reference dengan Vulnerability Database

### Using searchsploit (Exploit-DB)

**Install and Update:**
```bash
# Update exploit database
sudo apt update
sudo apt install exploitdb -y
searchsploit -u
```

---

**Search for Exploits:**
```bash
# Search by service name
searchsploit vsftpd
searchsploit openssh 4.7

# Search by CVE
searchsploit CVE-2021-41773

# Search by platform
searchsploit -p linux vsftpd

# Show exploit path
searchsploit -p vsftpd 2.3.4

# Copy exploit to current directory
searchsploit -m exploits/unix/remote/49757.py
```

---

**Example Workflow:**
```bash
# Create exploit research directory
mkdir -p ~/vulnerability_assessment/exploits

# From service inventory, search each service
# Example: vsftpd 2.3.4
searchsploit vsftpd 2.3.4 > ~/vulnerability_assessment/exploits/vsftpd_exploits.txt

# Example: OpenSSH 4.7p1
searchsploit openssh 4.7 > ~/vulnerability_assessment/exploits/openssh_exploits.txt

# Example: Apache 2.2.8
searchsploit apache 2.2.8 > ~/vulnerability_assessment/exploits/apache_exploits.txt
```

---

### Online Vulnerability Databases

**1. CVE Details (https://www.cvedetails.com/)**
```bash
# Search manually di website
# Format: [Product] [Version]
# Example: vsftpd 2.3.4
```

**2. Exploit-DB Online (https://www.exploit-db.com/)**
```bash
# Search di website atau gunakan searchsploit
```

**3. National Vulnerability Database (https://nvd.nist.gov/)**
```bash
# Advanced search by product, version, CVSS score
```

---

### Create Vulnerability Matrix

**Vulnerability Assessment Spreadsheet:**
```bash
# Create CSV file
cat > ~/vulnerability_assessment/analysis/vulnerability_matrix.csv << 'EOF'
Port,Service,Version,Vulnerability,CVE,CVSS,Exploitable,Impact,Recommendation
21,FTP,vsftpd 2.3.4,Backdoor,CVE-2011-2523,10.0,Yes,Critical,Update to latest version
22,SSH,OpenSSH 4.7p1,User Enumeration,CVE-2018-15473,5.3,Yes,Medium,Update to latest version
80,HTTP,Apache 2.2.8,Multiple,Various,7.5,Yes,High,Update to latest version
139,SMB,Samba 3.x,RCE,CVE-2017-7494,10.0,Yes,Critical,Upgrade Samba version
3306,MySQL,MySQL 5.0.51a,Multiple,Various,9.8,Yes,Critical,Update MySQL
EOF

# View with column
column -t -s ',' ~/vulnerability_assessment/analysis/vulnerability_matrix.csv
```

---

## 7.3 Risk Assessment

### CVSS Scoring

**CVSS (Common Vulnerability Scoring System):**
```
Score Range    Severity    Action Required
0.1 - 3.9      Low         Plan for patching
4.0 - 6.9      Medium      Patch within 30 days
7.0 - 8.9      High        Patch within 7 days
9.0 - 10.0     Critical    Patch immediately (24 hours)
```

---

### Vulnerability Prioritization

**Factors to Consider:**
1. **CVSS Score** - Severity rating
2. **Exploitability** - Is exploit publicly available?
3. **Asset Criticality** - How important is the system?
4. **Exposure** - Internet-facing or internal?
5. **Compensating Controls** - Firewall, IDS/IPS protection?

---

**Create Priority Matrix:**
```bash
cat > ~/vulnerability_assessment/analysis/priority_matrix.txt << 'EOF'
Priority 1 (Critical - Immediate Action):
- vsftpd 2.3.4 backdoor (Port 21)
  CVSS: 10.0, Public exploit available, Remote Code Execution
  
- Samba RCE CVE-2017-7494 (Port 139/445)
  CVSS: 10.0, Public exploit available, Remote Code Execution

Priority 2 (High - 7 Days):
- Apache 2.2.8 vulnerabilities (Port 80)
  CVSS: 7.5, Multiple exploits, Information disclosure & DoS
  
- MySQL 5.0.51a (Port 3306)
  CVSS: 9.8, Authentication bypass possible

Priority 3 (Medium - 30 Days):
- OpenSSH 4.7p1 User Enumeration (Port 22)
  CVSS: 5.3, Information disclosure
  
- Weak SMB configuration (Port 139/445)
  CVSS: 5.0, Anonymous access enabled

Priority 4 (Low - Maintenance Window):
- Banner disclosure on various services
  CVSS: 3.5, Information leakage
EOF
```

---

## 7.4 Evidence Collection

### Screenshot Best Practices

**What to Screenshot:**
1. Port scanning results
2. Service version detection
3. Vulnerability findings
4. Successful exploitation (if performed)
5. Proof of concept

---

**Using gnome-screenshot:**
```bash
# Full screen
gnome-screenshot

# Window selection
gnome-screenshot -w

# Area selection
gnome-screenshot -a

# Save to specific file
gnome-screenshot -f ~/vulnerability_assessment/evidence/nmap_scan_$(date +%Y%m%d_%H%M%S).png
```

---

### Log All Commands

**Create Command Log:**
```bash
# Start logging all commands
script ~/vulnerability_assessment/evidence/terminal_session_$(date +%Y%m%d_%H%M%S).log

# Do your scanning work...

# Stop logging
exit
```

---

## 🎯 LATIHAN SESI 7

### **Latihan 7.1: Vulnerability Research**

**Task:** Research vulnerabilities untuk semua service yang ditemukan

```bash
# Create research directory
mkdir -p ~/vulnerability_assessment/research

# For each service found, search for exploits
# Example for vsftpd 2.3.4:
searchsploit vsftpd 2.3.4 | tee ~/vulnerability_assessment/research/vsftpd_research.txt

# Get detailed info
searchsploit -x exploits/unix/remote/49757.py > ~/vulnerability_assessment/research/vsftpd_exploit_detail.txt

# Repeat for other services:
# - OpenSSH
# - Apache
# - Samba
# - MySQL
# - etc.
```

**Documentation:**
1. List of all vulnerabilities found
2. CVE numbers
3. CVSS scores
4. Exploit availability
5. Potential impact

---

### **Latihan 7.2: Create Vulnerability Report**

**Task:** Buat vulnerability assessment report

```bash
cat > ~/vulnerability_assessment/reports/VULNERABILITY_REPORT.md << 'EOF'
# Vulnerability Assessment Report

## Executive Summary
This report presents findings from a vulnerability assessment conducted on target system 192.168.56.102 (Metasploitable 2).

## Scope
- Target: 192.168.56.102
- Date: [Current Date]
- Methodology: Network scanning, service enumeration, vulnerability analysis

## Findings Summary
- Critical: X vulnerabilities
- High: X vulnerabilities  
- Medium: X vulnerabilities
- Low: X vulnerabilities

## Detailed Findings

### Finding 1: vsftpd 2.3.4 Backdoor
**Severity:** Critical (CVSS 10.0)
**Description:** vsftpd version 2.3.4 contains a backdoor that allows remote command execution.
**Evidence:** [Screenshot/Output]
**Impact:** Complete system compromise
**Recommendation:** Immediately update to vsftpd 3.x or later
**References:** CVE-2011-2523

[Continue for each finding...]

## Recommendations
1. Implement patch management process
2. Disable unnecessary services
3. Apply security hardening
4. Implement network segmentation
5. Deploy IDS/IPS

## Appendices
- Appendix A: Raw scan results
- Appendix B: Tool outputs
- Appendix C: Evidence screenshots

EOF
```

---

### **Latihan 7.3: Risk Assessment Matrix**

**Task:** Create comprehensive risk assessment

**Risk Calculation:**
```
Risk = Likelihood × Impact

Likelihood:
- Very High (5): Exploit publicly available, easy to execute
- High (4): Exploit available, moderate difficulty
- Medium (3): Exploit available, high difficulty
- Low (2): No public exploit, theoretical vulnerability
- Very Low (1): Requires specific conditions

Impact:
- Critical (5): Complete system compromise, data loss
- High (4): Significant data access, service disruption
- Medium (3): Limited data access, minor service impact
- Low (2): Information disclosure only
- Very Low (1): Minimal impact
```

**Create Risk Matrix:**
```bash
cat > ~/vulnerability_assessment/analysis/risk_matrix.csv << 'EOF'
Vulnerability,Likelihood,Impact,Risk Score,Risk Level
vsftpd Backdoor,5,5,25,Critical
Samba RCE,5,5,25,Critical
Apache Vulnerabilities,4,4,16,High
MySQL Weak Config,4,4,16,High
SSH User Enum,3,2,6,Medium
Banner Disclosure,5,1,5,Low
EOF

# View formatted
column -t -s ',' ~/vulnerability_assessment/analysis/risk_matrix.csv | tee ~/vulnerability_assessment/analysis/risk_matrix_formatted.txt
```

---

# SESI 8: REPORT WRITING & DOCUMENTATION (45 MENIT)

## 8.1 Professional Report Structure

### Standard Penetration Testing Report Format

**Report Sections:**
```
1. Executive Summary
2. Scope and Methodology
3. Findings Summary
4. Detailed Technical Findings
5. Risk Assessment
6. Recommendations
7. Conclusion
8. Appendices
```

---

### 1. Executive Summary

**Template:**
```markdown
# Executive Summary

## Purpose
This penetration test was conducted to assess the security posture of [Target System] and identify potential vulnerabilities that could be exploited by malicious actors.

## Scope
- Target: [IP/Domain]
- Date: [Start Date] - [End Date]
- Testing Type: Network Penetration Test
- Methodology: OWASP Testing Guide, PTES

## Key Findings
- [Number] Critical vulnerabilities identified
- [Number] High-risk vulnerabilities identified
- [Number] Medium-risk vulnerabilities identified
- [Number] Low-risk vulnerabilities identified

## Critical Issues
1. [Critical Issue 1]
2. [Critical Issue 2]

## Recommendations
The following actions should be taken immediately:
1. [Priority 1 recommendation]
2. [Priority 2 recommendation]
3. [Priority 3 recommendation]

## Risk Level
Overall Risk: [Critical/High/Medium/Low]
```

---

### 2. Scope and Methodology

**Template:**
```markdown
# Scope and Methodology

## Testing Scope
### In Scope
- IP Range: 192.168.56.0/24
- Services: All network services
- Testing Window: [Date/Time]

### Out of Scope
- Physical security testing
- Social engineering
- Denial of Service testing
- Production data modification

## Methodology
### Phase 1: Information Gathering
- Passive reconnaissance
- Active host discovery
- Network mapping

### Phase 2: Scanning and Enumeration
- Port scanning
- Service identification
- Version detection
- OS fingerprinting

### Phase 3: Vulnerability Analysis
- Automated scanning
- Manual verification
- Exploit research

### Phase 4: Reporting
- Evidence collection
- Risk assessment
- Report compilation

## Tools Used
- Nmap 7.94
- Metasploit Framework
- enum4linux
- searchsploit
- Various enumeration tools
```

---

### 3. Findings Summary

**Template:**
```markdown
# Findings Summary

## Vulnerability Distribution

| Severity | Count | Percentage |
|----------|-------|------------|
| Critical | X     | XX%        |
| High     | X     | XX%        |
| Medium   | X     | XX%        |
| Low      | X     | XX%        |
| **Total**| X     | 100%       |

## Services Assessed

| Port | Service     | Version         | Status       |
|------|-------------|-----------------|--------------|
| 21   | FTP         | vsftpd 2.3.4    | Vulnerable   |
| 22   | SSH         | OpenSSH 4.7p1   | Outdated     |
| 80   | HTTP        | Apache 2.2.8    | Vulnerable   |
| 139  | NetBIOS     | Samba 3.x       | Vulnerable   |
| 445  | SMB         | Samba 3.x       | Vulnerable   |
| 3306 | MySQL       | MySQL 5.0.51a   | Vulnerable   |

## Top 5 Critical Findings
1. vsftpd 2.3.4 Backdoor (CVE-2011-2523)
2. Samba Remote Code Execution (CVE-2017-7494)
3. MySQL Weak Configuration
4. Apache Multiple Vulnerabilities
5. Anonymous SMB Access
```

---

### 4. Detailed Technical Findings

**Finding Template:**
```markdown
## Finding X: [Vulnerability Name]

### Overview
**Severity:** [Critical/High/Medium/Low]  
**CVSS Score:** X.X  
**CVE:** CVE-XXXX-XXXXX  
**Affected Asset:** [IP:Port]  
**Service:** [Service Name]  

### Description
[Detailed description of the vulnerability]

### Evidence
```
[Command output or screenshot]
```

### Proof of Concept
```bash
[Commands to reproduce]
```

### Impact
[Explanation of potential impact to business/system]

### Recommendation
**Immediate Actions:**
1. [Action 1]
2. [Action 2]

**Long-term Solutions:**
1. [Solution 1]
2. [Solution 2]

### References
- [CVE Link]
- [Vendor Advisory]
- [Additional Resources]
```

---

### Example Complete Finding

```markdown
## Finding 1: vsftpd 2.3.4 Backdoor

### Overview
**Severity:** Critical  
**CVSS Score:** 10.0  
**CVE:** CVE-2011-2523  
**Affected Asset:** 192.168.56.102:21  
**Service:** vsftpd 2.3.4  

### Description
The vsftpd 2.3.4 version contains a backdoor that was inserted by an unknown intruder. This backdoor can be triggered by sending a username containing a ":)" smiley face, which opens a command shell on port 6200.

### Evidence
```
$ nmap -sV -p 21 192.168.56.102

PORT   STATE SERVICE VERSION
21/tcp open  ftp     vsftpd 2.3.4

$ searchsploit vsftpd 2.3.4
exploits/unix/remote/49757.py
```

### Proof of Concept
```bash
# Connect to FTP
telnet 192.168.56.102 21

# Send backdoor trigger
USER anonymous:)
PASS password

# Connect to backdoor
telnet 192.168.56.102 6200
# Command shell obtained
```

### Impact
An attacker can gain complete control of the system with root privileges. This allows:
- Complete data exfiltration
- System modification
- Pivot to other systems
- Installation of persistent backdoors
- Service disruption

Business Impact:
- Confidentiality: High - All data accessible
- Integrity: High - System can be modified
- Availability: High - System can be crashed

### Recommendation
**Immediate Actions:**
1. Immediately disable the FTP service
2. Disconnect the system from the network
3. Perform incident response procedures
4. Check logs for signs of compromise

**Long-term Solutions:**
1. Update vsftpd to version 3.x or later
2. If FTP is not required, disable the service permanently
3. If FTP is required, consider SFTP as alternative
4. Implement network segmentation to limit FTP access
5. Deploy IDS/IPS to detect exploitation attempts
6. Regular security assessments and patch management

### References
- https://nvd.nist.gov/vuln/detail/CVE-2011-2523
- https://www.exploit-db.com/exploits/49757
- https://security.appspot.com/vsftpd.html
```

---

## 8.2 Evidence Documentation

### Organize Evidence

**Directory Structure:**
```bash
mkdir -p ~/final_report/{screenshots,logs,scan_results,exploits}

# Copy evidence
cp ~/vulnerability_assessment/evidence/*.png ~/final_report/screenshots/
cp ~/vulnerability_assessment/evidence/*.log ~/final_report/logs/
cp ~/vulnerability_assessment/raw_scans/**/*.txt ~/final_report/scan_results/
```

---

### Screenshot Naming Convention

```bash
# Format: [date]_[target]_[finding]_[description].png
# Examples:
20250101_192.168.56.102_vsftpd_version_detection.png
20250101_192.168.56.102_smb_vulnerable_output.png
20250101_192.168.56.102_apache_exploit_poc.png
```

---

## 8.3 Recommendations Section

### Short-term Recommendations (0-30 days)

**Template:**
```markdown
# Recommendations

## Immediate Actions (0-7 days)

### 1. Patch Critical Vulnerabilities
**Priority:** Critical  
**Effort:** Medium  
**Cost:** Low  

**Actions:**
- Update vsftpd to version 3.0.5 or later
- Update Samba to version 4.6.4 or later
- Apply all available security patches

**Expected Outcome:**
- Elimination of remote code execution vulnerabilities
- Significant reduction in attack surface

---

### 2. Disable Unnecessary Services
**Priority:** High  
**Effort:** Low  
**Cost:** None  

**Actions:**
- Disable FTP service if not required
- Disable Telnet service
- Review and disable all non-essential services

**Expected Outcome:**
- Reduced attack surface
- Better resource utilization

---

## Short-term Actions (7-30 days)

### 3. Implement Network Segmentation
**Priority:** High  
**Effort:** High  
**Cost:** Medium  

**Actions:**
- Separate critical systems into secure network zones
- Implement VLAN segmentation
- Deploy firewall rules between segments

---

### 4. Deploy Intrusion Detection System
**Priority:** Medium  
**Effort:** Medium  
**Cost:** Medium  

**Actions:**
- Install Snort or Suricata IDS
- Configure signatures for known exploits
- Set up alerting mechanism

---

## Long-term Actions (30+ days)

### 5. Security Awareness Training
**Priority:** Medium  
**Effort:** Medium  
**Cost:** Low-Medium  

**Actions:**
- Conduct regular security training
- Implement phishing simulation exercises
- Create security policies and procedures

---

### 6. Regular Security Assessments
**Priority:** High  
**Effort:** Medium  
**Cost:** Medium  

**Actions:**
- Schedule quarterly vulnerability scans
- Annual penetration testing
- Continuous security monitoring
```

---

## 8.4 Create Final Report

### Complete Report Example

**Create Report:**
```bash
cat > ~/final_report/PENETRATION_TEST_REPORT.md << 'EOF'
# Penetration Testing Report
## Network Security Assessment - Metasploitable 2

---

**Report Date:** [Current Date]  
**Testing Period:** [Start Date] - [End Date]  
**Prepared by:** [Your Name]  
**Version:** 1.0  

---

# Table of Contents

1. Executive Summary
2. Scope and Methodology
3. Findings Summary
4. Detailed Technical Findings
   - 4.1 Critical Findings
   - 4.2 High Severity Findings
   - 4.3 Medium Severity Findings
   - 4.4 Low Severity Findings
5. Risk Assessment
6. Recommendations
7. Conclusion
8. Appendices

---

# 1. Executive Summary

[Insert executive summary here]

---

# 2. Scope and Methodology

[Insert scope and methodology here]

---

# 3. Findings Summary

[Insert findings summary table here]

---

# 4. Detailed Technical Findings

## 4.1 Critical Findings

### Finding 1: vsftpd 2.3.4 Backdoor
[Insert detailed finding using template above]

### Finding 2: Samba Remote Code Execution
[Insert detailed finding]

## 4.2 High Severity Findings
[Continue with other findings...]

---

# 5. Risk Assessment

[Insert risk matrix and analysis]

---

# 6. Recommendations

[Insert recommendations section]

---

# 7. Conclusion

This penetration test identified [X] vulnerabilities across [Y] services. The most critical issues require immediate attention to prevent potential system compromise.

Key takeaways:
1. [Takeaway 1]
2. [Takeaway 2]
3. [Takeaway 3]

Regular security assessments and proactive patch management are essential for maintaining a strong security posture.

---

# 8. Appendices

## Appendix A: Raw Scan Results
[Link to scan files]

## Appendix B: Tool Outputs
[Link to tool outputs]

## Appendix C: Evidence Screenshots
[Link to screenshots]

## Appendix D: References
[List of references used]

EOF
```

---

## 🎯 LATIHAN SESI 8 - FINAL PROJECT

### **Final Project: Complete Penetration Test Report**

**Task:** Buat complete penetration test report

**Requirements:**

**1. Perform Complete Assessment:**
```bash
# Full network assessment
sudo nmap -sS -sV -sC -O -p- --script vuln 192.168.56.102 -oA final_comprehensive_scan

# Detailed enumeration on all open ports
# - NetBIOS/SMB enumeration
# - SNMP enumeration (if available)
# - DNS enumeration
# - SMTP enumeration
# - Any other services found

# Document everything
```

---

**2. Create Professional Report:**
```
✅ Executive Summary (1 page)
✅ Scope and Methodology (1-2 pages)
✅ Findings Summary with charts (1 page)
✅ Detailed findings (minimum 5 findings)
   - Each finding with complete template
   - Screenshots for evidence
   - Proof of concept commands
✅ Risk Assessment Matrix
✅ Prioritized Recommendations
✅ Conclusion
✅ Appendices
```

---

**3. Format and Delivery:**
```
- Format: Markdown or PDF
- Include: Table of contents
- Include: Page numbers (if PDF)
- Include: All screenshots labeled
- Include: Professional formatting
- Size: Minimum 10 pages
```

---

**4. Presentation (15 minutes):**
```
Prepare short presentation covering:
- 5 minutes: Key findings
- 5 minutes: Risk assessment
- 5 minutes: Recommendations
- Q&A
```

---

## 📊 RUBRIC PENILAIAN FINAL PROJECT

| Kriteria | Bobot | Deskripsi |
|----------|-------|-----------|
| **Technical Accuracy** | 30% | - Scanning completed correctly<br>- Enumeration thorough<br>- Vulnerabilities accurately identified |
| **Report Quality** | 30% | - Professional format<br>- Clear writing<br>- Complete sections<br>- Proper evidence |
| **Risk Assessment** | 20% | - Accurate CVSS scoring<br>- Proper prioritization<br>- Business impact analysis |
| **Recommendations** | 15% | - Actionable recommendations<br>- Prioritized properly<br>- Realistic solutions |
| **Presentation** | 5% | - Clear communication<br>- Time management<br>- Professional delivery |

---

## 📚 LAMPIRAN: CHEAT SHEETS

### Nmap Quick Reference

```bash
# HOST DISCOVERY
nmap -sn 192.168.56.0/24              # Ping scan
nmap -PS22,80,443 192.168.56.102     # TCP SYN ping
nmap -PA 192.168.56.102               # TCP ACK ping
nmap -PU 192.168.56.102               # UDP ping

# PORT SCANNING
nmap -sS 192.168.56.102               # TCP SYN scan (stealth)
nmap -sT 192.168.56.102               # TCP connect scan
nmap -sU 192.168.56.102               # UDP scan
nmap -sN 192.168.56.102               # NULL scan
nmap -sF 192.168.56.102               # FIN scan
nmap -sX 192.168.56.102               # Xmas scan

# SERVICE/VERSION DETECTION
nmap -sV 192.168.56.102               # Version detection
nmap -O 192.168.56.102                # OS detection
nmap -A 192.168.56.102                # Aggressive (all above)

# PORT SPECIFICATION
nmap -p 80 192.168.56.102             # Single port
nmap -p 22,80,443 192.168.56.102     # Multiple ports
nmap -p 1-1000 192.168.56.102         # Port range
nmap -p- 192.168.56.102               # All ports
nmap -F 192.168.56.102                # Fast (top 100)

# NSE SCRIPTS
nmap --script=vuln 192.168.56.102     # Vulnerability scan
nmap --script=default 192.168.56.102  # Default scripts
nmap --script=http-* 192.168.56.102   # All HTTP scripts

# TIMING
nmap -T0 192.168.56.102               # Paranoid
nmap -T1 192.168.56.102               # Sneaky
nmap -T2 192.168.56.102               # Polite
nmap -T3 192.168.56.102               # Normal (default)
nmap -T4 192.168.56.102               # Aggressive
nmap -T5 192.168.56.102               # Insane

# OUTPUT
nmap -oN output.txt 192.168.56.102    # Normal
nmap -oX output.xml 192.168.56.102    # XML
nmap -oG output.gnmap 192.168.56.102  # Grepable
nmap -oA output 192.168.56.102        # All formats

# FIREWALL EVASION
nmap -f 192.168.56.102                # Fragment packets
nmap -D RND:10 192.168.56.102         # Decoy scan
nmap --source-port 53 192.168.56.102  # Source port
```

---

### Enumeration Commands Quick Reference

```bash
# NETBIOS/SMB
enum4linux -a 192.168.56.102
smbclient -L //192.168.56.102 -N
smbmap -H 192.168.56.102
nbtscan 192.168.56.102

# SNMP
snmpwalk -v2c -c public 192.168.56.102
snmp-check 192.168.56.102
onesixtyone 192.168.56.102

# LDAP
ldapsearch -x -h 192.168.56.102 -b "dc=example,dc=com"
nmap -p 389 --script ldap-rootdse 192.168.56.102

# DNS
dig example.com ANY
dig axfr @ns1.example.com example.com
dnsrecon -d example.com
dnsenum example.com

# SMTP
smtp-user-enum -M VRFY -U users.txt -t 192.168.56.102
nmap -p 25 --script smtp-enum-users 192.168.56.102

# NTP
ntpq -c readlist 192.168.56.102
nmap -sU -p 123 --script ntp-info 192.168.56.102
```

---

### Vulnerability Research Commands

```bash
# SEARCHSPLOIT
searchsploit [term]                   # Search exploits
searchsploit -x [id]                  # Examine exploit
searchsploit -m [id]                  # Mirror/copy exploit
searchsploit -u                       # Update database
searchsploit -w [term]                # Show URLs

# ONLINE RESOURCES
# https://www.cvedetails.com
# https://www.exploit-db.com
# https://nvd.nist.gov
# https://www.rapid7.com/db/
```

---

## 🎓 SERTIFIKAT PENYELESAIAN

### Kriteria Kelulusan

**Untuk mendapatkan sertifikat:**
```
✅ Attendance: Minimal 90% (7.2 jam dari 8 jam)
✅ Praktik Lab: Menyelesaikan minimal 80% latihan
✅ Final Project: Score minimal 70/100
✅ Presentation: Completed
```

---

## 📖 RESOURCES TAMBAHAN

### Buku dan Dokumentasi
1. "Nmap Network Scanning" - Gordon Lyon
2. "The Web Application Hacker's Handbook" - Dafydd Stuttard
3. "Metasploit: The Penetration Tester's Guide"
4. Official Nmap Documentation: https://nmap.org/book/

### Online Platforms
1. HackTheBox - https://hackthebox.eu
2. TryHackMe - https://tryhackme.com
3. PentesterLab - https://pentesterlab.com
4. VulnHub - https://vulnhub.com

### Video Tutorials
1. YouTube: NetworkChuck
2. YouTube: John Hammond
3. YouTube: IppSec (HackTheBox walkthroughs)
4. Udemy: Ethical Hacking courses

---

## 📞 SUPPORT & FEEDBACK

**Questions?**
- Email: [instructor@example.com]
- Office Hours: [Schedule]
- Discord/Slack: [Community Link]

**Feedback Form:**
Please provide feedback untuk improvement modul ini.

---

## ⚖️ LEGAL & ETHICAL REMINDER

**PENTING - BACA DENGAN SEKSAMA:**

1. ⚠️ **HANYA gunakan skills ini pada:**
   - Lab environment Anda sendiri
   - Sistem dengan izin tertulis
   - Platform legal (HackTheBox, TryHackMe, dll)

2. ❌ **JANGAN PERNAH:**
   - Scan network tanpa izin
   - Access sistem tanpa otorisasi
   - Modify atau delete data orang lain
   - Share vulnerability details publicly tanpa responsible disclosure

3. 📜 **Konsekuensi Legal:**
   - UU ITE Pasal 30: Pidana max 8 tahun
   - Denda hingga ratusan juta rupiah
   - Criminal record

4. ✅ **Ethical Hacking:**
   - Selalu dapatkan izin tertulis
   - Document semua aktivitas
   - Report findings secara bertanggung jawab
   - Respect privacy dan confidentiality

---

**"With great power comes great responsibility"**

**Selamat Belajar dan Tetap Etis!** 🚀🔒

---

**End of Module - Version 1.0**  
**Last Updated:** 2025