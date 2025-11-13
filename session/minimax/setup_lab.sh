#!/bin/bash

# Session Hijacking Security Lab Setup Script
# Author: MiniMax Agent
# Date: November 2025

echo "========================================="
echo "Session Hijacking Security Lab Setup"
echo "========================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root
if [[ $EUID -eq 0 ]]; then
   print_error "This script should not be run as root for security reasons"
   print_status "Please run as a regular user with sudo privileges"
   exit 1
fi

print_status "Starting lab environment setup..."

# Update system packages
print_status "Updating system packages..."
sudo apt update && sudo apt upgrade -y

# Install required tools
print_status "Installing network analysis tools..."
sudo apt install -y \
    tcpdump \
    wireshark-common \
    ettercap-common \
    nmap \
    netcat-openbsd \
    htop \
    curl \
    wget \
    git

# Install Go
print_status "Installing Go programming language..."
if ! command -v go &> /dev/null; then
    cd /tmp
    wget -q https://go.dev/dl/go1.21.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    print_success "Go installed successfully"
else
    print_success "Go is already installed"
fi

# Create project structure
print_status "Creating project directory structure..."
mkdir -p ~/session-security-lab/{cmd/{client,server,monitor},pkg/{hijack,detect,prevent},configs,scripts,data/{logs,certificates},tests}

# Set up Go workspace
export GOPATH=~/session-security-lab
export PATH=$PATH:$GOPATH/bin
print_status "Go workspace set up at $GOPATH"

# Install Go dependencies
print_status "Installing Go dependencies..."
cd ~/session-security-lab

# Initialize Go modules
go mod init session-security-lab

# Install required packages
go get github.com/google/gopacket
go get github.com/gorilla/mux
go get golang.org/x/crypto
go get github.com/gin-gonic/gin
go get github.com/stretchr/testify

print_success "Dependencies installed"

# Create certificate directory and generate self-signed certificates
print_status "Generating self-signed certificates for HTTPS testing..."
cd ~/session-security-lab/data/certificates

# Generate private key
openssl genrsa -out server.key 2048

# Generate certificate signing request
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=Test/L=Test/O=Security Lab/OU=IT/CN=localhost"

# Generate self-signed certificate
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

print_success "Certificates generated"

# Set up monitoring scripts
print_status "Creating monitoring scripts..."

cat > ~/session-security-lab/scripts/network-monitor.sh << 'EOF'
#!/bin/bash
# Network monitoring script for session hijacking detection

INTERFACE="eth0"
LOG_FILE="/home/$USER/session-security-lab/data/logs/network-monitor.log"

echo "$(date): Starting network monitoring on $INTERFACE" >> $LOG_FILE

# Monitor TCP traffic for suspicious patterns
sudo tcpdump -i $INTERFACE -n 'tcp[tcpflags] & tcp-push !=0' -l | while read line; do
    echo "$(date): $line" >> $LOG_FILE
    
    # Check for rapid-fire requests (potential automated attack)
    if echo "$line" | grep -q "Flags \[P\]"; then
        TIMESTAMP=$(date +%s)
        echo "SUSPICIOUS TCP ACTIVITY DETECTED: $line" >> $LOG_FILE
    fi
done
EOF

chmod +x ~/session-security-lab/scripts/network-monitor.sh

# Create session analysis script
cat > ~/session-security-lab/scripts/session-analyzer.sh << 'EOF'
#!/bin/bash
# Session analysis script

LOG_FILE="/home/$USER/session-security-lab/data/logs/session-analysis.log"

echo "$(date): Analyzing session patterns..."

# Analyze HTTP requests for session patterns
sudo tcpdump -i any -n 'tcp port 80 or 443' -l | while read line; do
    if echo "$line" | grep -q "GET\|POST"; then
        SESSION_ID=$(echo "$line" | grep -oP 'session_id=\K[^;]+')
        if [ ! -z "$SESSION_ID" ]; then
            echo "$(date): Session request - $SESSION_ID" >> $LOG_FILE
        fi
    fi
done
EOF

chmod +x ~/session-security-lab/scripts/session-analyzer.sh

# Create attack simulation script
cat > ~/session-security-lab/scripts/simulate-attacks.sh << 'EOF'
#!/bin/bash
# Session hijacking attack simulation script

TARGET_IP="127.0.0.1"
TARGET_PORT="8080"
LOG_FILE="/home/$USER/session-security-lab/data/logs/attack-simulation.log"

echo "$(date): Starting attack simulations..." | tee -a $LOG_FILE

# Simulate session replay attack
echo "Simulating session replay attack..."
for i in {1..10}; do
    curl -X POST http://$TARGET_IP:$TARGET_PORT/login \
         -H "Content-Type: application/json" \
         -d '{"username":"test","password":"test"}' \
         --cookie "session_id=replayed_session_$i" \
         --silent >> $LOG_FILE 2>&1
    sleep 1
done

# Simulate session fixation
echo "Simulating session fixation..."
for i in {1..5}; do
    curl -X GET http://$TARGET_IP:$TARGET_PORT/ \
         -H "Cookie: session_id=fixed_session_$i" \
         --silent >> $LOG_FILE 2>&1
    sleep 1
done

echo "$(date): Attack simulations completed" | tee -a $LOG_FILE
EOF

chmod +x ~/session-security-lab/scripts/simulate-attacks.sh

# Create test configuration files
print_status "Creating configuration files..."

# Nginx configuration for reverse proxy
cat > ~/session-security-lab/configs/nginx.conf << 'EOF'
server {
    listen 80;
    server_name localhost;
    
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

server {
    listen 443 ssl;
    server_name localhost;
    
    ssl_certificate /home/$(whoami)/session-security-lab/data/certificates/server.crt;
    ssl_certificate_key /home/$(whoami)/session-security-lab/data/certificates/server.key;
    
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
EOF

# Create iptables rules for lab
cat > ~/session-security-lab/configs/iptables-rules.sh << 'EOF'
#!/bin/bash
# iptables rules for lab environment

# Allow incoming SSH
sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT

# Allow incoming HTTP/HTTPS
sudo iptables -A INPUT -p tcp --dport 80 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 443 -j ACCEPT

# Allow local loopback
sudo iptables -A INPUT -i lo -j ACCEPT

# Drop all other incoming traffic
sudo iptables -A INPUT -j DROP

# Log dropped packets
sudo iptables -A INPUT -j LOG --log-prefix "DROPPED INPUT: "

echo "iptables rules configured for lab environment"
EOF

chmod +x ~/session-security-lab/configs/iptables-rules.sh

# Create Go example files
print_status "Creating Go example implementations..."

# Basic session server
cat > ~/session-security-lab/cmd/server/main.go << 'EOF'
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "net/http"
    "sync"
    "time"
)

type Session struct {
    ID        string
    UserID    string
    Created   time.Time
    LastUsed  time.Time
    IPAddress string
    UserAgent string
}

type SessionManager struct {
    sessions sync.Map
}

func generateSecureToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

func (sm *SessionManager) CreateSession(w http.ResponseWriter, userID string) (string, error) {
    sessionID, err := generateSecureToken()
    if err != nil {
        return "", err
    }

    session := Session{
        ID:        sessionID,
        UserID:    userID,
        Created:   time.Now(),
        LastUsed:  time.Now(),
        IPAddress: "127.0.0.1",
        UserAgent: "Test Client",
    }

    sm.sessions.Store(sessionID, session)
    
    cookie := &http.Cookie{
        Name:     "session_id",
        Value:    sessionID,
        Path:     "/",
        HttpOnly: true,
        Secure:   false, // Set to true in production with HTTPS
        SameSite: http.SameSiteStrictMode,
        Expires:  time.Now().Add(30 * time.Minute),
    }
    
    http.SetCookie(w, cookie)
    return sessionID, nil
}

func main() {
    sm := &SessionManager{}
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Session Security Lab Server\n")
        fmt.Fprintf(w, "Available endpoints:\n")
        fmt.Fprintf(w, "GET / - This page\n")
        fmt.Fprintf(w, "POST /login - Create session\n")
        fmt.Fprintf(w, "GET /secure - Access protected resource\n")
    })
    
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            sessionID, err := sm.CreateSession(w, "test_user")
            if err != nil {
                http.Error(w, "Failed to create session", http.StatusInternalServerError)
                return
            }
            fmt.Fprintf(w, "Session created: %s\n", sessionID)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    http.HandleFunc("/secure", func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_id")
        if err != nil {
            http.Error(w, "No session found", http.StatusUnauthorized)
            return
        }
        
        value, exists := sm.sessions.Load(cookie.Value)
        if !exists {
            http.Error(w, "Invalid session", http.StatusUnauthorized)
            return
        }
        
        session := value.(Session)
        fmt.Fprintf(w, "Welcome! Session ID: %s\n", session.ID)
        fmt.Fprintf(w, "User: %s\n", session.UserID)
    })
    
    fmt.Println("Starting session security lab server on :8080")
    http.ListenAndServe(":8080", nil)
}
EOF

# Session hijacking detection tool
cat > ~/session-security-lab/cmd/monitor/main.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "time"
)

type SessionEvent struct {
    Timestamp time.Time
    SessionID string
    IPAddress string
    UserAgent string
    EventType string
}

type SessionDetector struct {
    events []SessionEvent
}

func (sd *SessionDetector) RecordEvent(event SessionEvent) {
    sd.events = append(sd.events, event)
}

func (sd *SessionDetector) DetectReplay() []SessionEvent {
    var replayAttacks []SessionEvent
    sessionMap := make(map[string][]SessionEvent)
    
    for _, event := range sd.events {
        sessionMap[event.SessionID] = append(sessionMap[event.SessionID], event)
    }
    
    for sessionID, events := range sessionMap {
        if len(events) > 1 {
            // Check for rapid replay attempts
            for i := 1; i < len(events); i++ {
                if events[i].Timestamp.Sub(events[i-1].Timestamp) < time.Second {
                    replayAttacks = append(replayAttacks, events[i])
                    fmt.Printf("REPLAY DETECTED: Session %s reused rapidly\n", sessionID)
                }
            }
        }
    }
    
    return replayAttacks
}

func main() {
    detector := &SessionDetector{}
    
    // Simulate some events
    now := time.Now()
    detector.RecordEvent(SessionEvent{
        Timestamp: now,
        SessionID: "session1",
        IPAddress: "192.168.1.100",
        EventType: "login",
    })
    
    detector.RecordEvent(SessionEvent{
        Timestamp: now.Add(time.Millisecond * 500),
        SessionID: "session1",
        IPAddress: "192.168.1.101", // Different IP
        EventType: "replay",
    })
    
    replayAttacks := detector.DetectReplay()
    fmt.Printf("Detected %d replay attacks\n", len(replayAttacks))
    
    if len(replayAttacks) > 0 {
        os.Exit(1) // Attack detected
    }
}
EOF

# Create go.mod file
cat > ~/session-security-lab/go.mod << 'EOF'
module session-security-lab

go 1.21

require (
    github.com/google/gopacket v1.1.17
    github.com/gorilla/mux v1.8.1
)
EOF

# Set proper permissions
chmod -R 755 ~/session-security-lab
chmod 600 ~/session-security-lab/data/certificates/*

# Create README
cat > ~/session-security-lab/README.md << 'EOF'
# Session Hijacking Security Lab

This lab environment provides hands-on experience with session hijacking attacks and countermeasures.

## Setup Complete!

Your session security lab environment is now ready. Here's what has been installed:

### Tools Installed:
- tcpdump - Network packet capture
- wireshark - Network protocol analyzer  
- nmap - Network scanner
- ettercap - Network sniffer
- Go programming language

### Project Structure:
```
session-security-lab/
├── cmd/
│   ├── server/          # Target server application
│   ├── monitor/         # Security monitoring tool
│   └── client/          # Client test application
├── pkg/
│   ├── hijack/          # Session hijacking implementations
│   ├── detect/          # Detection algorithms
│   └── prevent/         # Prevention mechanisms
├── configs/             # Configuration files
├── scripts/             # Monitoring and analysis scripts
├── data/
│   ├── certificates/    # SSL/TLS certificates
│   └── logs/           # Log files
└── tests/              # Test cases
```

### Quick Start:

1. **Start the target server:**
   ```bash
   cd ~/session-security-lab
   go run cmd/server/main.go
   ```

2. **In another terminal, start monitoring:**
   ```bash
   cd ~/session-security-lab
   go run cmd/monitor/main.go
   ```

3. **Test session creation:**
   ```bash
   curl -X POST http://localhost:8080/login
   ```

4. **Access protected resource:**
   ```bash
   curl --cookie "session_id=YOUR_SESSION_ID" http://localhost:8080/secure
   ```

### Monitoring Scripts:

- `scripts/network-monitor.sh` - Network traffic monitoring
- `scripts/session-analyzer.sh` - Session pattern analysis  
- `scripts/simulate-attacks.sh` - Attack simulation

### Learning Objectives:

1. Understand different types of session hijacking attacks
2. Implement session security measures
3. Use network analysis tools for detection
4. Build monitoring and alerting systems
5. Practice defense-in-depth strategies

### Security Warnings:

⚠️ **IMPORTANT**: This lab environment is for educational purposes only.
- Only test on systems you own or have explicit permission to test
- All attack simulations should be performed in isolated lab environments
- Remember to follow responsible disclosure practices

### Next Steps:

1. Review the presentation slides in `session_hijacking_presentation.html`
2. Run through the lab exercises step by step
3. Experiment with different attack scenarios
4. Implement additional security measures
5. Build your own detection algorithms

Happy learning!
EOF

print_success "Lab environment setup completed!"
print_status "Next steps:"
echo "1. Review the presentation: ~/session_hijacking_presentation.html"
echo "2. Start learning: cd ~/session-security-lab && go run cmd/server/main.go"
echo "3. Check README: cat ~/session-security-lab/README.md"
echo ""
print_warning "Remember: Only test on systems you own or have explicit permission to test!"
echo ""
print_success "Setup script completed successfully!"