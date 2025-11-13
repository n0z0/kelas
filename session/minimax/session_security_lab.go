package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Session represents a user session
type Session struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Created     time.Time `json:"created"`
	LastUsed    time.Time `json:"last_used"`
	IPAddress   string    `json:"ip_address"`
	UserAgent   string    `json:"user_agent"`
	Fingerprint string    `json:"fingerprint"`
}

// SecurityEvent represents a security event
type SecurityEvent struct {
	Timestamp   time.Time `json:"timestamp"`
	EventType   string    `json:"event_type"`
	SessionID   string    `json:"session_id"`
	SourceIP    string    `json:"source_ip"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"` // LOW, MEDIUM, HIGH, CRITICAL
}

// SessionManager handles session lifecycle and security
type SessionManager struct {
	sessions    sync.Map
	events      []SecurityEvent
	eventMutex  sync.Mutex
	config      *SecurityConfig
}

// SecurityConfig holds security configuration
type SecurityConfig struct {
	SessionTimeout    time.Duration
	MaxSessionsPerIP  int
	EnableFingerprint bool
	EnableAnomalyDetection bool
}

// SessionHijackingDetector detects various types of session hijacking
type SessionHijackingDetector struct {
	manager      *SessionManager
	knownIPs     map[string]time.Time
	replayTokens map[string]time.Time
}

// NewSessionManager creates a new session manager
func NewSessionManager(config *SecurityConfig) *SessionManager {
	return &SessionManager{
		config: config,
		events: make([]SecurityEvent, 0),
	}
}

// generateSecureToken generates a cryptographically secure session token
func generateSecureToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// generateFingerprint creates a session fingerprint
func generateFingerprint(sessionID string, userAgent string) string {
	data := sessionID + "|" + userAgent + "|salt"
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:16])
}

// CreateSession creates a new session with security measures
func (sm *SessionManager) CreateSession(w http.ResponseWriter, userID string, r *http.Request) (string, error) {
	sessionID, err := generateSecureToken()
	if err != nil {
		return "", err
	}

	clientIP := sm.getClientIP(r)
	userAgent := r.Header.Get("User-Agent")

	session := Session{
		ID:          sessionID,
		UserID:      userID,
		Created:     time.Now(),
		LastUsed:    time.Now(),
		IPAddress:   clientIP,
		UserAgent:   userAgent,
		Fingerprint: generateFingerprint(sessionID, userAgent),
	}

	sm.sessions.Store(sessionID, session)
	sm.recordEvent("SESSION_CREATED", sessionID, clientIP, "New session created", "LOW")

	// Set secure cookie
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(sm.config.SessionTimeout),
	}

	http.SetCookie(w, cookie)
	return sessionID, nil
}

// ValidateSession validates session with security checks
func (sm *SessionManager) ValidateSession(r *http.Request) (bool, *Session, string) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, nil, "No session cookie found"
	}

	value, exists := sm.sessions.Load(cookie.Value)
	if !exists {
		return false, nil, "Session not found"
	}

	session := value.(Session)
	clientIP := sm.getClientIP(r)
	userAgent := r.Header.Get("User-Agent")

	// Check session timeout
	if time.Since(session.LastUsed) > sm.config.SessionTimeout {
		sm.sessions.Delete(cookie.Value)
		sm.recordEvent("SESSION_EXPIRED", session.ID, clientIP, "Session expired", "MEDIUM")
		return false, nil, "Session expired"
	}

	// Update last used time
	session.LastUsed = time.Now()
	sm.sessions.Store(cookie.Value, session)

	// Security validations
	if sm.config.EnableFingerprint {
		expectedFingerprint := generateFingerprint(session.ID, userAgent)
		if session.Fingerprint != expectedFingerprint {
			sm.recordEvent("FINGERPRINT_MISMATCH", session.ID, clientIP, 
				fmt.Sprintf("Fingerprint mismatch: expected %s, got %s", 
				session.Fingerprint, expectedFingerprint), "HIGH")
			return false, nil, "Session fingerprint mismatch"
		}
	}

	// Check IP consistency (optional, can be disabled for mobile users)
	if session.IPAddress != clientIP && clientIP != "" {
		sm.recordEvent("IP_CHANGE", session.ID, clientIP, 
			fmt.Sprintf("IP changed from %s to %s", session.IPAddress, clientIP), "MEDIUM")
		// In production, you might want to re-authenticate here
	}

	return true, &session, ""
}

// getClientIP extracts client IP from request
func (sm *SessionManager) getClientIP(r *http.Request) string {
	// Check for forwarded headers (load balancers, proxies)
	if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
		return xForwardedFor
	}
	if xRealIP := r.Header.Get("X-Real-IP"); xRealIP != "" {
		return xRealIP
	}
	
	// Fallback to direct connection
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// recordEvent records a security event
func (sm *SessionManager) recordEvent(eventType, sessionID, sourceIP, description, severity string) {
	sm.eventMutex.Lock()
	defer sm.eventMutex.Unlock()

	event := SecurityEvent{
		Timestamp:   time.Now(),
		EventType:   eventType,
		SessionID:   sessionID,
		SourceIP:    sourceIP,
		Description: description,
		Severity:    severity,
	}

	sm.events = append(sm.events, event)
	
	// Log to console
	fmt.Printf("[%s] %s: %s (Session: %s, IP: %s)\n", 
		severity, eventType, description, sessionID, sourceIP)
}

// NewSessionHijackingDetector creates a new detector
func NewSessionHijackingDetector(manager *SessionManager) *SessionHijackingDetector {
	return &SessionHijackingDetector{
		manager:     manager,
		knownIPs:    make(map[string]time.Time),
		replayTokens: make(map[string]time.Time),
	}
}

// detectSessionReplay detects session replay attacks
func (d *SessionHijackingDetector) detectSessionReplay(sessionID string, clientIP string) bool {
	if lastSeen, exists := d.replayTokens[sessionID]; exists {
		if time.Since(lastSeen) < time.Second*5 {
			d.manager.recordEvent("SESSION_REPLAY_DETECTED", sessionID, clientIP,
				"Rapid session token reuse detected", "CRITICAL")
			return true
		}
	}
	
	d.replayTokens[sessionID] = time.Now()
	return false
}

// detectBruteForceSessionAccess detects brute force access attempts
func (d *SessionHijackingDetector) detectBruteForceSessionAccess(clientIP string) bool {
	if lastSeen, exists := d.knownIPs[clientIP]; exists {
		if time.Since(lastSeen) < time.Second*10 {
			d.manager.recordEvent("BRUTE_FORCE_DETECTED", "", clientIP,
				"Rapid session access attempts from same IP", "HIGH")
			return true
		}
	}
	
	d.knownIPs[clientIP] = time.Now()
	return false
}

// NetworkMonitor monitors network traffic for suspicious activity
type NetworkMonitor struct {
	device     string
	snaplen    int
	promiscuous bool
	timeout    time.Duration
	detector   *SessionHijackingDetector
}

// NewNetworkMonitor creates a new network monitor
func NewNetworkMonitor(device string, detector *SessionHijackingDetector) *NetworkMonitor {
	return &NetworkMonitor{
		device:      device,
		snaplen:     1600,
		promiscuous: false,
		timeout:     pcap.BlockForever,
		detector:    detector,
	}
}

// StartMonitoring starts packet capture and analysis
func (nm *NetworkMonitor) StartMonitoring() error {
	handle, err := pcap.OpenLive(nm.device, nm.snaplen, nm.promiscuous, nm.timeout)
	if err != nil {
		return fmt.Errorf("failed to open device %s: %v", nm.device, err)
	}
	defer handle.Close()

	// Set filter to capture only HTTP/HTTPS traffic
	err = handle.SetBPFFilter("tcp port 80 or 443")
	if err != nil {
		return fmt.Errorf("failed to set BPF filter: %v", err)
	}

	fmt.Printf("Starting network monitoring on interface: %s\n", nm.device)
	
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		nm.analyzePacket(packet)
	}
	
	return nil
}

// analyzePacket analyzes a captured packet
func (nm *NetworkMonitor) analyzePacket(packet gopacket.Packet) {
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer == nil {
		return
	}
	
	payload := applicationLayer.Payload()
	if len(payload) == 0 {
		return
	}
	
	// Look for session-related patterns in HTTP traffic
	payloadStr := string(payload)
	
	// Check for session ID patterns
	if contains(payloadStr, "session_id=") || contains(payloadStr, "cookie:") {
		fmt.Printf("[NETWORK] Detected session-related traffic\n")
		
		// Extract potential session information
		if nm.detector != nil {
			// This would need more sophisticated parsing in a real implementation
			nm.detector.detectBruteForceSessionAccess("unknown")
		}
	}
}

// contains checks if a string contains a substring (simple version)
func contains(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// HTTP Handlers for the demo application

// homeHandler shows the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Session Security Lab</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			.container { max-width: 800px; margin: 0 auto; }
			.card { border: 1px solid #ddd; padding: 20px; margin: 20px 0; border-radius: 8px; }
			button { background: #007cba; color: white; padding: 10px 20px; border: none; border-radius: 4px; cursor: pointer; }
			button:hover { background: #005a87; }
			.result { background: #f5f5f5; padding: 15px; border-radius: 4px; margin: 10px 0; }
			.security-event { border-left: 4px solid #ff4444; background: #ffeeee; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Session Security Lab</h1>
			<div class="card">
				<h3>🔐 Create Session</h3>
				<button onclick="createSession()">Create Test Session</button>
				<div id="sessionResult" class="result" style="display:none;"></div>
			</div>
			
			<div class="card">
				<h3>🔍 Access Protected Resource</h3>
				<button onclick="accessSecure()">Access Secure Area</button>
				<div id="secureResult" class="result" style="display:none;"></div>
			</div>
			
			<div class="card">
				<h3>⚠️ Security Events</h3>
				<button onclick="getSecurityEvents()">View Security Events</button>
				<div id="eventsResult" class="result" style="display:none;"></div>
			</div>
			
			<div class="card">
				<h3>🎯 Attack Simulations</h3>
				<button onclick="simulateReplay()">Simulate Session Replay</button>
				<button onclick="simulateFixation()">Simulate Session Fixation</button>
				<div id="attackResult" class="result" style="display:none;"></div>
			</div>
		</div>
		
		<script>
			async function createSession() {
				try {
					const response = await fetch('/login', { method: 'POST' });
					const data = await response.text();
					document.getElementById('sessionResult').innerHTML = data;
					document.getElementById('sessionResult').style.display = 'block';
				} catch (error) {
					document.getElementById('sessionResult').innerHTML = 'Error: ' + error.message;
					document.getElementById('sessionResult').style.display = 'block';
				}
			}
			
			async function accessSecure() {
				try {
					const response = await fetch('/secure');
					const data = await response.text();
					document.getElementById('secureResult').innerHTML = data;
					document.getElementById('secureResult').style.display = 'block';
				} catch (error) {
					document.getElementById('secureResult').innerHTML = 'Error: ' + error.message;
					document.getElementById('secureResult').style.display = 'block';
				}
			}
			
			async function getSecurityEvents() {
				try {
					const response = await fetch('/events');
					const data = await response.json();
					let html = '<h4>Recent Security Events:</h4>';
					data.forEach(event => {
						html += '<div class="security-event"><strong>' + event.severity + '</strong>: ' + 
							   event.event_type + ' - ' + event.description + 
							   ' (' + new Date(event.timestamp).toLocaleString() + ')</div>';
					});
					document.getElementById('eventsResult').innerHTML = html;
					document.getElementById('eventsResult').style.display = 'block';
				} catch (error) {
					document.getElementById('eventsResult').innerHTML = 'Error: ' + error.message;
					document.getElementById('eventsResult').style.display = 'block';
				}
			}
			
			async function simulateReplay() {
				try {
					// First create a session
					await fetch('/login', { method: 'POST' });
					
					// Simulate replay attack
					const response = await fetch('/replay-simulate', { method: 'POST' });
					const data = await response.text();
					document.getElementById('attackResult').innerHTML = data;
					document.getElementById('attackResult').style.display = 'block';
				} catch (error) {
					document.getElementById('attackResult').innerHTML = 'Error: ' + error.message;
					document.getElementById('attackResult').style.display = 'block';
				}
			}
			
			async function simulateFixation() {
				try {
					const response = await fetch('/fixation-simulate', { method: 'POST' });
					const data = await response.text();
					document.getElementById('attackResult').innerHTML = data;
					document.getElementById('attackResult').style.display = 'block';
				} catch (error) {
					document.getElementById('attackResult').innerHTML = 'Error: ' + error.message;
					document.getElementById('attackResult').style.display = 'block';
				}
			}
		</script>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

// loginHandler creates a new session
func (sm *SessionManager) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID, err := sm.CreateSession(w, "demo_user", r)
	if err != nil {
		http.Error(w, "Failed to create session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "✅ Session created successfully!\nSession ID: %s\n", sessionID)
}

// secureHandler validates session and shows secure content
func (sm *SessionManager) secureHandler(w http.ResponseWriter, r *http.Request) {
	valid, session, error := sm.ValidateSession(r)
	if !valid {
		http.Error(w, "❌ Access denied: "+error, http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "🔐 Welcome to the secure area!\n")
	fmt.Fprintf(w, "Session ID: %s\n", session.ID)
	fmt.Fprintf(w, "User ID: %s\n", session.UserID)
	fmt.Fprintf(w, "IP Address: %s\n", session.IPAddress)
	fmt.Fprintf(w, "User Agent: %s\n", session.UserAgent)
	fmt.Fprintf(w, "Last Used: %s\n", session.LastUsed.Format("2006-01-02 15:04:05"))
}

// eventsHandler returns security events
func (sm *SessionManager) eventsHandler(w http.ResponseWriter, r *http.Request) {
	sm.eventMutex.Lock()
	defer sm.eventMutex.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sm.events)
}

// replaySimulateHandler simulates a session replay attack
func (d *SessionHijackingDetector) replaySimulateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simulate replay attack
	d.detectSessionReplay("fake_session_123", "192.168.1.100")
	d.detectSessionReplay("fake_session_123", "192.168.1.100") // Rapid reuse
	
	fmt.Fprintf(w, "🎯 Session replay attack simulated!\n")
	fmt.Fprintf(w, "Check the console and events endpoint for detected security events.\n")
}

// fixationSimulateHandler simulates a session fixation attack
func (sm *SessionManager) fixationSimulateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simulate session fixation attack by setting a predetermined session ID
	attackSessionID := "attacker_controlled_session"
	
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    attackSessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	
	http.SetCookie(w, cookie)
	sm.recordEvent("SESSION_FIXATION_SIMULATED", attackSessionID, 
		"192.168.1.200", "Session fixation attack simulated", "HIGH")
	
	fmt.Fprintf(w, "🎯 Session fixation attack simulated!\n")
	fmt.Fprintf(w, "Attacker-controlled session ID set: %s\n", attackSessionID)
	fmt.Fprintf(w, "Check the events endpoint for security alerts.\n")
}

// main function sets up and starts the application
func main() {
	// Configuration
	config := &SecurityConfig{
		SessionTimeout:             30 * time.Minute,
		MaxSessionsPerIP:           5,
		EnableFingerprint:          true,
		EnableAnomalyDetection:     true,
	}

	// Initialize components
	manager := NewSessionManager(config)
	detector := NewSessionHijackingDetector(manager)
	monitor := NewNetworkMonitor("eth0", detector) // Change interface as needed

	// Set up HTTP handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", manager.loginHandler)
	http.HandleFunc("/secure", manager.secureHandler)
	http.HandleFunc("/events", manager.eventsHandler)
	http.HandleFunc("/replay-simulate", detector.replaySimulateHandler)
	http.HandleFunc("/fixation-simulate", manager.fixationSimulateHandler)

	// Start network monitoring in background
	go func() {
		fmt.Println("Starting network monitoring...")
		if err := monitor.StartMonitoring(); err != nil {
			fmt.Printf("Network monitoring error: %v\n", err)
		}
	}()

	port := ":8080"
	fmt.Printf("🚀 Session Security Lab Server starting on port %s\n", port)
	fmt.Printf("📱 Open http://localhost%s in your browser\n", port)
	fmt.Printf("🔍 Security events will be logged to console\n")
	fmt.Printf("📊 View events at http://localhost%s/events\n", port)
	fmt.Println("\n⚠️  Security Warnings:")
	fmt.Println("   - This is a demonstration server")
	fmt.Println("   - Use only in isolated lab environments")
	fmt.Println("   - Monitor console for security events")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}