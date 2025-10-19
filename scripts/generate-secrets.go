// SECURITY_NOTE: This is a template file for generating production secrets
// This file does NOT contain actual credentials - it only shows how to generate them
// All generated values should be stored securely and never committed to version control

package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

// generateRandomHex creates a cryptographically secure random hex string
func generateRandomHex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic("Failed to generate random bytes: " + err.Error())
	}
	return hex.EncodeToString(bytes)
}

func main() {
	fmt.Println("# MCP Ultra Secret Generator")
	fmt.Printf("# Generated at: %s\n", time.Now().Format(time.RFC3339))
	fmt.Println("# SECURITY: These are GENERATED values - store securely!")
	fmt.Println()

	// Environment variable names (no actual values here)
	envVars := []string{
		"JWT_SECRET",
		"ENCRYPTION_KEY",
		"AUDIT_LOG_ENCRYPTION_KEY",
		"GRPC_SERVER_TOKEN",
		"GRPC_CLIENT_TOKEN",
		"DB_PASSWORD",
		"NATS_TOKEN",
	}

	// Generate and output each secret
	for _, envVar := range envVars {
		value := generateRandomHex(32) // Generate 32-byte hex string
		fmt.Printf("export %s=%s\n", envVar, value)
	}

	fmt.Println()
	fmt.Println("# SECURITY REMINDER:")
	fmt.Println("# 1. Save this output to a secure location")
	fmt.Println("# 2. Never commit these values to source control")
	fmt.Println("# 3. Use environment variables or secret management")
	fmt.Println("# 4. Rotate secrets every 90 days")

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Println("\nUsage: go run generate-secrets.go")
		fmt.Println("Generates cryptographically secure random secrets for MCP Ultra")
	}
}
