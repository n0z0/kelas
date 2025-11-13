// Package certgen provides certificate generation functionality for local development
package main

import (
	"fmt"
	"os"
	certgen "session/z/z"
)

// GenerateCertificates generates self-signed SSL/TLS certificates for local development
func GenerateCertificates(outputDir string, keyType string, keySize int, validityDays int) error {
	config := certgen_module.DefaultConfig()
	config.OutputDir = outputDir
	config.KeyType = keyType
	config.KeySize = keySize
	config.ValidityDays = validityDays

	return certgen_module.GenerateCertificateWithConfig(config)
}

// GenerateCertificatesInteractive provides an interactive command-line interface for certificate generation
func GenerateCertificatesInteractive() {
	config := certgen.DefaultConfig()

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
	if err := certgen.GenerateCertificateWithConfig(config); err != nil {
		fmt.Printf("Failed to generate certificate: %v\n", err)
		os.Exit(1)
	}
}

// GenerateCertificatesForServer generates certificates suitable for the HTTPS server
func GenerateCertificatesForServer() error {
	config := certgen.DefaultConfig()
	config.OutputDir = "."
	config.Domains = []string{"localhost", "127.0.0.1"}

	if err := certgen.GenerateCertificateWithConfig(config); err != nil {
		return fmt.Errorf("failed to generate certificates: %v", err)
	}

	// Rename files to match expected names for the server
	if err := os.Rename("localhost.crt", "cert.pem"); err != nil {
		return fmt.Errorf("failed to rename certificate: %v", err)
	}

	if err := os.Rename("localhost.key", "key.pem"); err != nil {
		return fmt.Errorf("failed to rename private key: %v", err)
	}

	return nil
}
