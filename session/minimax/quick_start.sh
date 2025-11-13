#!/bin/bash

# Quick Start Script for Session Hijacking Security Workshop
# Author: MiniMax Agent
# Date: November 2025

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# ASCII Art Banner
cat << 'EOF'
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║   🔐 SESSION HIJACKING SECURITY WORKSHOP 🔐                 ║
║                                                              ║
║   Comprehensive Security Training with Hands-on Practice   ║
║                                                              ║
║   ✨ Theory + CLI Tools + Golang Implementation ✨          ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
EOF

print_banner() {
    echo -e "${PURPLE}===================================================${NC}"
    echo -e "${PURPLE} $1 ${NC}"
    echo -e "${PURPLE}===================================================${NC}"
}

print_step() {
    echo -e "${CYAN}[STEP $1]${NC} $2"
}

print_success() {
    echo -e "${GREEN}[✓]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[⚠]${NC} $1"
}

print_error() {
    echo -e "${RED}[✗]${NC} $1"
}

# Check prerequisites
print_banner "Checking Prerequisites"

# Check OS
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    print_success "Linux OS detected"
else
    print_error "This workshop is designed for Linux. Current OS: $OSTYPE"
    exit 1
fi

# Check if Go is installed
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_success "Go is installed (version: $GO_VERSION)"
else
    print_warning "Go is not installed. Installing..."
    
    # Install Go
    cd /tmp
    wget -q https://go.dev/dl/go1.21.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    
    print_success "Go installed successfully"
fi

# Check if sudo access is available
if sudo -n true 2>/dev/null; then
    print_success "Sudo access available"
else
    print_warning "Sudo access will be requested when needed"
fi

# Create workspace
print_banner "Setting Up Workspace"

WORKSPACE_DIR="$HOME/session-security-workshop"
if [ -d "$WORKSPACE_DIR" ]; then
    print_warning "Workspace directory already exists. Backing up..."
    mv "$WORKSPACE_DIR" "${WORKSPACE_DIR}_backup_$(date +%Y%m%d_%H%M%S)"
fi

mkdir -p "$WORKSPACE_DIR"
cd "$WORKSPACE_DIR"
print_success "Workspace created at: $WORKSPACE_DIR"

# Install required tools
print_banner "Installing Required Tools"

print_step "1" "Installing network analysis tools..."
sudo apt update -qq
sudo apt install -y -qq tcpdump wireshark-common nmap netcat-openbsd curl wget git > /dev/null 2>&1
print_success "Network tools installed"

print_step "2" "Installing Go dependencies..."
export GOPATH="$WORKSPACE_DIR"
export PATH=$PATH:$GOPATH/bin
mkdir -p "$GOPATH"

# Create a temporary go.mod for dependency installation
echo 'module temp
go 1.21
require (
    github.com/google/gopacket v1.1.17
    github.com/gorilla/mux v1.8.1
)' > "$GOPATH/go.mod"

cd "$GOPATH"
go get github.com/google/gopacket > /dev/null 2>&1
go get github.com/gorilla/mux > /dev/null 2>&1
print_success "Go dependencies installed"

# Setup project structure
print_banner "Creating Project Structure"

mkdir -p {cmd/{server,client,monitor},pkg/{hijack,detect,prevent},configs,scripts,data/{logs,certificates},tests}
print_success "Project structure created"

# Generate certificates
print_step "3" "Generating SSL certificates..."
mkdir -p data/certificates
cd data/certificates

openssl genrsa -out server.key 2048 2>/dev/null
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=Test/L=Test/O=Security Lab/OU=IT/CN=localhost" 2>/dev/null
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt 2>/dev/null

cd "$WORKSPACE_DIR"
print_success "SSL certificates generated"

# Copy files
print_banner "Setting Up Workshop Files"

print_step "4" "Setting up main application..."
# The main application file should be copied here by the user
if [ -f "../session_security_lab.go" ]; then
    cp "../session_security_lab.go" cmd/server/main.go
    print_success "Main application configured"
else
    print_warning "Main application file not found. Please copy session_security_lab.go to cmd/server/main.go"
fi

# Create go.mod for the project
cat > go.mod << 'EOF'
module session-security-workshop

go 1.21

require (
    github.com/google/gopacket v1.1.17
    github.com/gorilla/mux v1.8.1
    golang.org/x/crypto v0.17.0
)
EOF

print_success "Project configuration complete"

# Set permissions
chmod -R 755 "$WORKSPACE_DIR"
chmod 600 "$WORKSPACE_DIR/data/certificates/*"

# Create helpful scripts
print_banner "Creating Helper Scripts"

# Start server script
cat > scripts/start-server.sh << 'EOF'
#!/bin/bash
echo "🚀 Starting Session Security Lab Server..."
echo "📍 Server will be available at: http://localhost:8080"
echo "📱 Open http://localhost:8080 in your browser"
echo "📊 Security events will be displayed in console"
echo ""
echo "⚠️  Press Ctrl+C to stop the server"
echo ""

cd "$(dirname "$0")/.."
go run cmd/server/main.go
EOF

chmod +x scripts/start-server.sh

# Network monitor script
cat > scripts/start-monitor.sh << 'EOF'
#!/bin/bash
echo "🔍 Starting Network Monitor..."
echo "📊 Monitoring network traffic for suspicious activity"
echo ""

INTERFACE="eth0"
if [ ! -z "$1" ]; then
    INTERFACE="$1"
fi

echo "Using interface: $INTERFACE"
echo "Press Ctrl+C to stop monitoring"
echo ""

sudo tcpdump -i "$INTERFACE" -n 'tcp port 80 or 443' -A
EOF

chmod +x scripts/start-monitor.sh

# Attack simulator script
cat > scripts/simulate-attacks.sh << 'EOF'
#!/bin/bash
echo "🎯 Session Hijacking Attack Simulator"
echo "======================================"
echo ""

# Check if server is running
if ! curl -s http://localhost:8080 > /dev/null; then
    echo "❌ Server is not running. Please start the server first:"
    echo "   ./scripts/start-server.sh"
    exit 1
fi

echo "1. Simulating Session Replay Attack..."
curl -X POST http://localhost:8080/replay-simulate -s
echo ""

echo "2. Simulating Session Fixation Attack..."
curl -X POST http://localhost:8080/fixation-simulate -s
echo ""

echo "3. Checking Security Events..."
echo "Recent Security Events:"
curl -s http://localhost:8080/events | jq -r '.[] | "\(.severity): \(.event_type) - \(.description)"' 2>/dev/null || \
curl -s http://localhost:8080/events | grep -o '"description":"[^"]*"' | cut -d'"' -f4

echo ""
echo "✅ Attack simulation complete!"
echo "🔍 Check the server console for security alerts"
EOF

chmod +x scripts/simulate-attacks.sh

print_success "Helper scripts created"

# Final setup and instructions
print_banner "Setup Complete! 🎉"

echo -e "${GREEN}"
cat << 'EOF'
╔══════════════════════════════════════════════════════════════╗
║                   🎊 WORKSHOP READY! 🎊                     ║
╚══════════════════════════════════════════════════════════════╝
EOF
echo -e "${NC}"

echo -e "${CYAN}📋 NEXT STEPS:${NC}"
echo ""
echo -e "${YELLOW}1. Start the Workshop:${NC}"
echo -e "   ${GREEN}firefox session_hijacking_presentation.html${NC}"
echo -e "   ${GREEN}# or${NC}"
echo -e "   ${GREEN}google-chrome session_hijacking_presentation.html${NC}"
echo ""
echo -e "${YELLOW}2. Start the Lab Server:${NC}"
echo -e "   ${GREEN}./scripts/start-server.sh${NC}"
echo ""
echo -e "${YELLOW}3. In another terminal, try these commands:${NC}"
echo -e "   ${GREEN}# Monitor network traffic${NC}"
echo -e "   ${GREEN}./scripts/start-monitor.sh eth0${NC}"
echo ""
echo -e "   ${GREEN}# Simulate attacks${NC}"
echo -e "   ${GREEN}./scripts/simulate-attacks.sh${NC}"
echo ""
echo -e "${YELLOW}4. Open Browser:${NC}"
echo -e "   ${GREEN}http://localhost:8080${NC}"
echo ""

echo -e "${CYAN}📁 WORKSPACE STRUCTURE:${NC}"
echo -e "${WORKSPACE_DIR}/"
echo -e "├── 📄 session_hijacking_presentation.html  # Main presentation"
echo -e "├── 📖 WORKSHOP_GUIDE.md                   # Comprehensive guide"
echo -e "├── 🚀 scripts/"
echo -e "│   ├── start-server.sh                    # Start lab server"
echo -e "│   ├── start-monitor.sh                   # Network monitoring"
echo -e "│   └── simulate-attacks.sh                # Attack simulation"
echo -e "├── 🐹 cmd/server/"
echo -e "│   └── main.go                           # Main application"
echo -e "└── 📊 data/"
echo -e "    ├── certificates/                     # SSL certificates"
echo -e "    └── logs/                            # Log files"
echo ""

echo -e "${PURPLE}🎯 QUICK COMMANDS:${NC}"
echo -e "${GREEN}# Start everything${NC}"
echo -e "./scripts/start-server.sh &"
echo -e "./scripts/start-monitor.sh eth0 &"
echo ""
echo -e "${GREEN}# Test the system${NC}"
echo -e "curl -X POST http://localhost:8080/login"
echo -e "curl http://localhost:8080/events"
echo ""

echo -e "${RED}⚠️  SECURITY WARNING:${NC}"
echo -e "• This is for educational purposes only"
echo -e "• Only test on systems you own or have permission to test"
echo -e "• All attacks are simulated in isolated lab environment"
echo ""

echo -e "${BLUE}💡 TIPS:${NC}"
echo -e "• Read WORKSHOP_GUIDE.md for detailed instructions"
echo -e "• Watch the console for security event alerts"
echo -e "• Experiment with different attack scenarios"
echo -e "• Try modifying the code to add new features"
echo ""

print_success "Happy Learning! 🚀"

# Ask if user wants to start the server now
echo -e "${CYAN}🤔 Would you like to start the server now? (y/n)${NC}"
read -r response

if [[ "$response" =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}🚀 Starting server...${NC}"
    ./scripts/start-server.sh
else
    echo -e "${YELLOW}👍 To start later, run: ./scripts/start-server.sh${NC}"
fi