package certgen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

// Config holds the configuration for certificate generation
type Config struct {
	KeyType      string   // "rsa" or "ecdsa"
	KeySize      int      // 2048, 4096 for RSA, 256, 384 for ECDSA
	ValidityDays int      // Number of days the certificate should be valid
	OutputDir    string   // Directory to save certificates
	Domains      []string // List of domains to include in the certificate
}

// DefaultConfig returns a default configuration
func DefaultConfig() Config {
	return Config{
		KeyType:      "rsa",
		KeySize:      2048,
		ValidityDays: 365,
		OutputDir:    "./certs",
		Domains:      []string{"localhost", "127.0.0.1", "::1"},
	}
}

// GenerateKeyPair generates a private key based on the configuration
func GenerateKeyPair(config Config) (interface{}, error) {
	switch config.KeyType {
	case "rsa":
		if config.KeySize < 512 {
			return nil, fmt.Errorf("RSA key size must be at least 512 bits")
		}
		return rsa.GenerateKey(rand.Reader, config.KeySize)
	case "ecdsa":
		var curve elliptic.Curve
		switch config.KeySize {
		case 256:
			curve = elliptic.P256()
		case 384:
			curve = elliptic.P384()
		case 521:
			curve = elliptic.P521()
		default:
			return nil, fmt.Errorf("unsupported ECDSA key size: %d. Use 256, 384, or 521", config.KeySize)
		}
		return ecdsa.GenerateKey(curve, rand.Reader)
	default:
		return nil, fmt.Errorf("unsupported key type: %s. Use 'rsa' or 'ecdsa'", config.KeyType)
	}
}

// GenerateCertificate generates a self-signed certificate
func GenerateCertificate(privateKey interface{}, config Config) (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Local Development"},
			CommonName:   "Local Development Certificate",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(0, 0, config.ValidityDays),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Add IP addresses and domains
	for _, domain := range config.Domains {
		if ip := net.ParseIP(domain); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, domain)
		}
	}

	// Self-sign the certificate
	_, err = x509.CreateCertificate(rand.Reader, &template, &template, publicKey(privateKey), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	return &template, nil
}

// publicKey extracts the public key from the private key
func publicKey(privateKey interface{}) interface{} {
	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// SaveCertificate saves the certificate and private key to files
func SaveCertificate(cert *x509.Certificate, privateKey interface{}, config Config) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Save certificate
	certPath := filepath.Join(config.OutputDir, "localhost.crt")
	certFile, err := os.Create(certPath)
	if err != nil {
		return fmt.Errorf("failed to create certificate file: %v", err)
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}); err != nil {
		return fmt.Errorf("failed to write certificate: %v", err)
	}

	// Save private key
	keyPath := filepath.Join(config.OutputDir, "localhost.key")
	keyFile, err := os.Create(keyPath)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %v", err)
	}
	defer keyFile.Close()

	var keyBytes []byte
	var blockType string

	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		keyBytes = x509.MarshalPKCS1PrivateKey(k)
		blockType = "RSA PRIVATE KEY"
	case *ecdsa.PrivateKey:
		keyBytes, err = x509.MarshalECPrivateKey(k)
		if err != nil {
			return fmt.Errorf("failed to marshal ECDSA private key: %v", err)
		}
		blockType = "EC PRIVATE KEY"
	default:
		return fmt.Errorf("unsupported private key type")
	}

	if err := pem.Encode(keyFile, &pem.Block{
		Type:  blockType,
		Bytes: keyBytes,
	}); err != nil {
		return fmt.Errorf("failed to write private key: %v", err)
	}

	return nil
}

// PrintOutput displays information about the generated certificate
func PrintOutput(cert *x509.Certificate, privateKey interface{}, config Config) {
	fmt.Println("=== Certificate Generated Successfully ===")
	fmt.Printf("Key Type: %s\n", config.KeyType)
	if config.KeyType == "rsa" {
		fmt.Printf("Key Size: %d bits\n", config.KeySize)
	} else {
		fmt.Printf("Key Size: %d\n", config.KeySize)
	}
	fmt.Printf("Validity Period: %d days\n", config.ValidityDays)
	fmt.Printf("Output Directory: %s\n", config.OutputDir)
	fmt.Println("\nGenerated Files:")
	fmt.Printf("  Certificate: %s\n", filepath.Join(config.OutputDir, "localhost.crt"))
	fmt.Printf("  Private Key: %s\n", filepath.Join(config.OutputDir, "localhost.key"))

	fmt.Println("\nCertificate Details:")
	fmt.Printf("  Serial Number: %s\n", cert.SerialNumber.String())
	fmt.Printf("  Common Name: %s\n", cert.Subject.CommonName)
	fmt.Printf("  Organization: %s\n", cert.Subject.Organization[0])
	fmt.Printf("  Valid From: %s\n", cert.NotBefore.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Valid Until: %s\n", cert.NotAfter.Format("2006-01-02 15:04:05"))

	fmt.Println("\nDomains/IP Addresses:")
	for _, name := range cert.DNSNames {
		fmt.Printf("  DNS: %s\n", name)
	}
	for _, ip := range cert.IPAddresses {
		fmt.Printf("  IP: %s\n", ip.String())
	}

	fmt.Println("\nNote: This certificate is self-signed and intended for local development only.")
	fmt.Println("      Your browser may show a security warning when using this certificate.")
}

// GenerateCertificateWithConfig generates a certificate using the provided configuration
func GenerateCertificateWithConfig(config Config) error {
	// Generate key pair
	fmt.Println("Generating key pair...")
	privateKey, err := GenerateKeyPair(config)
	if err != nil {
		return fmt.Errorf("failed to generate key pair: %v", err)
	}

	// Generate certificate
	fmt.Println("Generating certificate...")
	cert, err := GenerateCertificate(privateKey, config)
	if err != nil {
		return fmt.Errorf("failed to generate certificate: %v", err)
	}

	// Save certificate
	fmt.Println("Saving certificate files...")
	if err := SaveCertificate(cert, privateKey, config); err != nil {
		return fmt.Errorf("failed to save certificate: %v", err)
	}

	// Print output
	PrintOutput(cert, privateKey, config)
	return nil
}

// GenerateInteractiveCertificate generates a certificate with interactive command-line interface
func GenerateInteractiveCertificate() {
	config := DefaultConfig()

	// Simple command-line interface
	fmt.Println("=== HTTPS Certificate Generator for Local Development ===")
	fmt.Println()

	// Get key type
	fmt.Print("Enter key type (rsa/ecdsa) [default: rsa]: ")
	var keyTypeInput string
	fmt.Scanln(&keyTypeInput)
	if keyTypeInput != "" {
		config.KeyType = keyTypeInput
	}

	// Get key size
	fmt.Printf("Enter key size (%d for %s, %d for %s) [default: %d]: ",
		2048, "RSA", 256, "ECDSA", config.KeySize)
	var keySizeInput string
	fmt.Scanln(&keySizeInput)
	if keySizeInput != "" {
		fmt.Sscanf(keySizeInput, "%d", &config.KeySize)
	}

	// Get validity period
	fmt.Printf("Enter validity period in days [default: %d]: ", config.ValidityDays)
	var validityInput string
	fmt.Scanln(&validityInput)
	if validityInput != "" {
		fmt.Sscanf(validityInput, "%d", &config.ValidityDays)
	}

	// Get output directory
	fmt.Printf("Enter output directory [default: %s]: ", config.OutputDir)
	var outputDirInput string
	fmt.Scanln(&outputDirInput)
	if outputDirInput != "" {
		config.OutputDir = outputDirInput
	}

	fmt.Println()

	// Generate certificate using the helper function
	if err := GenerateCertificateWithConfig(config); err != nil {
		log.Fatalf("Failed to generate certificate: %v", err)
	}
}
