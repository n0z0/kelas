# Panduan Lengkap Workshop Session Hijacking Security

**Penulis:** MiniMax Agent  
**Tanggal:** November 2025  
**Durasi:** 1 Hari Workshop

## 📋 Daftar Isi

1. [Pendahuluan](#pendahuluan)
2. [Struktur File Workshop](#struktur-file-workshop)
3. [Persiapan Lingkungan](#persiapan-lingkungan)
4. [Langkah-langkah Workshop](#langkah-langkah-workshop)
5. [Contoh Praktik dengan CLI](#contoh-praktik-dengan-cli)
6. [Implementasi Golang](#implementasi-golang)
7. [Skenario Serangan](#skenario-serangan)
8. [Metode Deteksi](#metode-deteksi)
9. [Strategi Pencegahan](#strategi-pencegahan)
10. [Troubleshooting](#troubleshooting)

## 🎯 Pendahuluan

Workshop ini dirancang untuk memberikan pemahaman komprehensif tentang session hijacking attacks dan countermeasure-nya melalui kombinasi teoria dan praktik. Peserta akan belajar mengidentifikasi berbagai teknik hijacking, mengimplementasikan kontrol keamanan, dan menggunakan tools untuk deteksi.

### Tujuan Pembelajaran:
- Memahami berbagai jenis session hijacking attacks
- Mengimplementasikan security controls untuk mencegah attacks
- Menggunakan tools deteksi untuk mengidentifikasi attempts
- Menerapkan Golang untuk membuat security tools
- Menguasai CLI tools untuk network analysis dan security testing

## 📁 Struktur File Workshop

```
session-security-workshop/
├── session_hijacking_presentation.html  # Slide presentasi utama
├── setup_lab.sh                        # Script setup otomatis
├── session_security_lab.go             # Aplikasi demo lengkap
├── README.md                           # Panduan ini
└── supplementary/                      # File pendukung
    ├── attack_scenarios.md             # Skenario serangan
    ├── cli_commands.md                 # Kumpulan command CLI
    └── golang_examples.md              # Contoh kode tambahan
```

## 🔧 Persiapan Lingkungan

### Prasyarat:
- Sistem operasi Linux (Ubuntu 20.04+ direkomendasikan)
- Akses sudo privileges
- Koneksi internet untuk download tools
- Browser web modern

### Setup Otomatis:
```bash
# 1. Buat direktori workshop
mkdir ~/session-security-workshop
cd ~/session-security-workshop

# 2. Copy semua file workshop ke direktori ini
# (File sudah disediakan dalam workspace)

# 3. Jalankan setup script
chmod +x setup_lab.sh
./setup_lab.sh
```

### Setup Manual (jika diperlukan):
```bash
# Install Go
curl -LO https://go.dev/dl/go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install tools network
sudo apt update
sudo apt install -y tcpdump wireshark-common nmap ettercap-common

# Install dependencies Go
go get github.com/google/gopacket
go get github.com/gorilla/mux
```

## 🚀 Langkah-langkah Workshop

### Fase 1: Teori (30 menit)
1. **Buka Presentasi:**
   ```bash
   # Buka di browser
   firefox session_hijacking_presentation.html
   # atau
   google-chrome session_hijacking_presentation.html
   ```

2. **Ikuti Slide 1-4:** Overview session hijacking dan tipos-tipos serangan

3. **Diskusi:** Real-world examples dan consequences

### Fase 2: Demo Praktik (45 menit)

#### Demo 1: Session Hijacking Basics
```bash
# Start target server
go run session_security_lab.go

# In another terminal, test session creation
curl -X POST http://localhost:8080/login

# Copy session ID and test access
curl --cookie "session_id=YOUR_SESSION_ID" http://localhost:8080/secure
```

#### Demo 2: Network Analysis
```bash
# Monitor network traffic (as root)
sudo tcpdump -i any -n 'tcp port 8080' -A

# Analyze with Wireshark (GUI)
wireshark &

# Use nmap for port scanning
nmap -sS -O localhost
```

### Fase 3: Hands-on Lab (90 menit)

#### Lab 1: Session Security Implementation
```bash
# 1. Start the demo application
go run session_security_lab.go

# 2. Open browser to http://localhost:8080

# 3. Test the following:
#    - Create session
#    - Access secure area
#    - View security events
#    - Simulate attacks
```

#### Lab 2: Attack Simulation
```bash
# Simulate session replay
curl -X POST http://localhost:8080/replay-simulate

# Simulate session fixation
curl -X POST http://localhost:8080/fixation-simulate

# Check events
curl http://localhost:8080/events
```

### Fase 4: Advanced Topics (60 menit)

#### Advanced 1: Network-level Hijacking
```bash
# Monitor TCP connections
sudo netstat -an | grep :8080

# Capture packets for analysis
sudo tcpdump -i any -w session_analysis.pcap 'tcp port 8080'

# Analyze with custom Go tool (see examples)
```

#### Advanced 2: Building Custom Tools
- Modify `session_security_lab.go`
- Add new detection algorithms
- Implement additional security measures

## 💻 Contoh Praktik dengan CLI

### Network Monitoring Commands:
```bash
# Basic packet capture
sudo tcpdump -i eth0 -n 'tcp[tcpflags] & tcp-push !=0'

# Capture HTTP sessions
sudo tcpdump -i any -A 'tcp port 80 or 443'

# Monitor specific interface
sudo tcpdump -i wlan0 -n 'host 192.168.1.100'

# Save to file for analysis
sudo tcpdump -i eth0 -w session_capture.pcap
```

### Session Analysis Commands:
```bash
# View active connections
sudo netstat -an | grep :80

# Monitor connection patterns
watch -n 1 'netstat -an | grep :8080 | wc -l'

# Check for unusual IPs
netstat -an | awk '{print $5}' | grep -v ':' | sort | uniq -c
```

### Security Testing Commands:
```bash
# Port scanning
nmap -sS -O target_ip

# Test SSL/TLS configuration
openssl s_client -connect target_ip:443

# Check HTTP headers
curl -I http://target_ip:80
```

## 🐹 Implementasi Golang

### Basic Session Management:
```go
// Generate secure session token
func generateSecureToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Validate session with security checks
func (sm *SessionManager) ValidateSession(r *http.Request) (bool, *Session, string) {
    // Implementation details in session_security_lab.go
}
```

### Detection Algorithms:
```go
// Detect session replay attacks
func (d *SessionHijackingDetector) detectSessionReplay(sessionID string, clientIP string) bool {
    if lastSeen, exists := d.replayTokens[sessionID]; exists {
        if time.Since(lastSeen) < time.Second*5 {
            // Potential replay attack detected
            return true
        }
    }
    d.replayTokens[sessionID] = time.Now()
    return false
}
```

### Network Monitoring:
```go
// Monitor network traffic for suspicious patterns
func (nm *NetworkMonitor) analyzePacket(packet gopacket.Packet) {
    // Extract and analyze packet data
    // Look for session-related patterns
    // Trigger security alerts
}
```

## 🎯 Skenario Serangan

### 1. Session Replay Attack
**Skenario:** Attacker captures valid session token dan menggunakan kembali untuk akses unauthorized.

**Demonstrasi:**
```bash
# Capture legitimate session
SESSION=$(curl -X POST http://localhost:8080/login | grep "Session ID:" | cut -d: -f2)

# Replay the same session multiple times
for i in {1..10}; do
    curl --cookie "session_id=$SESSION" http://localhost:8080/secure
    sleep 0.5  # Rapid reuse should trigger detection
done
```

**Expected Detection:** System mendeteksi rapid reuse dan mencatat security event.

### 2. Session Fixation Attack
**Skenario:** Attacker memaksa victim untuk menggunakan session ID yang sudah dikontrol attacker.

**Demonstrasi:**
```bash
# Attacker provides session ID
ATTACKER_SESSION="malicious_session_123"

# Victim uses attacker-controlled session
curl --cookie "session_id=$ATTACKER_SESSION" http://localhost:8080/secure
```

**Expected Detection:** System mendeteksi session fixation pattern dan menginvalidasi session.

### 3. Man-in-the-Browser (MITB)
**Skenario:** Malware di browser intercepts dan manipulates web traffic.

**Simulasi:** Browser extension atau proxy yang memodifikasi requests/responses.

**Detection:** Certificate pinning, browser integrity checks.

### 4. Network-level Hijacking
**Skenario:** Attacker intercepts network traffic dan mengambil kontrol TCP connection.

**Simulasi:** ARP spoofing atau packet interception.

**Detection:** Sequence number analysis, connection monitoring.

## 🔍 Metode Deteksi

### 1. Real-time Monitoring
```bash
# Start continuous monitoring
./scripts/network-monitor.sh &

# Check logs
tail -f ~/session-security-lab/data/logs/network-monitor.log
```

### 2. Anomaly Detection
```go
// Behavioral analysis
func (ba *BehavioralAnalyzer) DetectAnomaly(sessionID string) bool {
    activities := ba.sessionPatterns[sessionID]
    if len(activities) < 2 {
        return false
    }
    
    last := activities[len(activities)-1]
    secondLast := activities[len(activities)-2]
    
    // Detect rapid-fire requests
    if last.Sub(secondLast) < time.Second {
        return true
    }
    
    return false
}
```

### 3. Pattern Recognition
```bash
# Analyze traffic patterns
awk '/SESSION/ {print $1, $2, $3, $4, $5, $6}' /var/log/nginx/access.log | sort | uniq -c
```

### 4. Certificate Validation
```go
// Certificate pinning validation
func (d *MITBDetector) CheckCertificatePinning(url string) error {
    // Implement certificate verification logic
    // Detect unauthorized certificate changes
    return nil
}
```

## 🛡️ Strategi Pencegahan

### 1. Transport Layer Security
```go
// Set secure session cookies
func SetSecureSessionCookie(w http.ResponseWriter, name, value string) {
    cookie := &http.Cookie{
        Name:     name,
        Value:    value,
        Path:     "/",
        HttpOnly: true,
        Secure:   true,  // HTTPS only
        SameSite: http.SameSiteStrictMode,
        Expires:  time.Now().Add(30 * time.Minute),
    }
    http.SetCookie(w, cookie)
}
```

### 2. Session Management
```go
// Regenerate session ID after login
func RegenerateSessionID(w http.ResponseWriter, r *http.Request) {
    // Store old data
    // Create new session with new ID
    // Delete old session
}
```

### 3. Multi-factor Authentication
```go
// Implement 2FA
func (sm *SessionManager) ValidateWith2FA(sessionID string, code string) bool {
    // Validate session and 2FA code
    return isValidCode(sessionID, code)
}
```

### 4. IP Validation
```go
// Validate IP consistency
func validateIPConsistency(session Session, clientIP string) bool {
    if session.IPAddress != clientIP && clientIP != "" {
        // Log suspicious IP change
        return false
    }
    return true
}
```

## ⚠️ Troubleshooting

### Common Issues:

#### 1. Permission Errors
```bash
# Fix file permissions
chmod +x setup_lab.sh
chmod 755 ~/session-security-lab

# For network monitoring (require root)
sudo tcpdump -i eth0 -n 'tcp port 80'
```

#### 2. Go Dependencies
```bash
# Reinstall Go modules
go mod download
go get github.com/google/gopacket
go get github.com/gorilla/mux
```

#### 3. Port Conflicts
```bash
# Check what's using port 8080
sudo netstat -tlnp | grep :8080

# Kill process if needed
sudo kill -9 PID
```

#### 4. Network Interface Issues
```bash
# List available interfaces
ip link show

# Update network interface in code
# Change "eth0" to your interface name
```

### Debug Mode:
```bash
# Run with debug output
DEBUG=1 go run session_security_lab.go

# Enable detailed logging
export SESSION_DEBUG=1
```

## 📊 Evaluasi Workshop

### Checklist Pemahaman:
- [ ] Memahami berbagai tipos session hijacking attacks
- [ ] Dapat mengidentifikasi vulnerabilities dalam web applications
- [ ] Menguasai CLI tools untuk network analysis
- [ ] Dapat mengimplementasikan session security measures
- [ ] Memahami detection algorithms dan implementasinya
- [ ] Dapat membangun monitoring dan alerting systems

### Hands-on Assessment:
1. **Attack Simulation:** Peserta melakukan simulasi berbagai tipos serangan
2. **Detection Challenge:** Peserta mengimplementasikan custom detection algorithm
3. **Security Implementation:** Peserta membangun secure session management system

## 🎓 Referensi dan Sumber Belajar

### Dokumentasi:
- [OWASP Session Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html)
- [RFC 6265 - HTTP State Management Mechanism](https://tools.ietf.org/html/rfc6265)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)

### Tools References:
- [tcpdump Manual](https://www.tcpdump.org/manpages/tcpdump.1.html)
- [Wireshark User Guide](https://www.wireshark.org/docs/wsug_html/)
- [Go net/http Package Documentation](https://golang.org/pkg/net/http/)

### Books:
- "Web Application Security: A Beginner's Guide" by Bryan Sullivan
- "The Web Application Hacker's Handbook" by Dafydd Stuttard
- "Network Security Essentials" by William Stallings

## 📞 Support dan Kontak

**Instructor:** MiniMax Agent  
**Email:** support@minimax-agent.com  
**Workshop Materials:** Available in workshop directory  
**Lab Environment:** Ready-to-use virtual machines available on request

---

**Security Notice:** Workshop ini menggunakan simulasi attacks dalam environment yang isolated. Semua techniques yang dipelajari harus digunakan secara etis dan legal. Penggunaan untuk unauthorized access adalah illegal dan tidak ethic.

**Last Updated:** November 2025