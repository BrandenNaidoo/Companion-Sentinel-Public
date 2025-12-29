package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Config
const API_URL = "https://the-companion.com/api/agent/heartbeat" // Replace with localhost for dev
const HEARTBEAT_INTERVAL = 60 * time.Second

type Check struct {
	Type       string `json:"type"`
	Status     string `json:"status"` // pass, fail, warn
	Details    string `json:"details"`
	FixCommand string `json:"fix_command"`
}

type Payload struct {
	OSInfo      map[string]string `json:"os_info"`
	Checks      []Check           `json:"checks"`
	HealthScore int               `json:"health_score"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: companion-agent <API_KEY>")
		return
	}
	apiKey := os.Args[1]

	fmt.Println("Starting Companion Agent...")
	fmt.Printf("API Key: %s...\n", apiKey[:5])

	for {
		// 1. Gather Data
		checks, score := runChecks()
		
		payload := Payload{
			OSInfo: map[string]string{
				"os":   runtime.GOOS,
				"arch": runtime.GOARCH,
				"hostname": getHostname(),
			},
			Checks:      checks,
			HealthScore: score,
		}

		// 2. Send Heartbeat
		err := sendHeartbeat(apiKey, payload)
		if err != nil {
			fmt.Printf("Error sending heartbeat: %v\n", err)
		} else {
			fmt.Println("Heartbeat sent successfully.")
		}

		time.Sleep(HEARTBEAT_INTERVAL)
	}
}

func runChecks() ([]Check, int) {
	var checks []Check
	points := 100

	// Check 1: SSH Root Login
	if isRootLoginEnabled() {
		checks = append(checks, Check{
			Type:       "ssh_root",
			Status:     "fail",
			Details:    "PermitRootLogin is enabled. This allows attackers to brute-force the root user directly.",
			FixCommand: "sudo sed -i 's/PermitRootLogin yes/PermitRootLogin no/g' /etc/ssh/sshd_config && sudo systemctl restart ssh",
		})
		points -= 30
	} else {
		checks = append(checks, Check{
			Type:    "ssh_root",
			Status:  "pass",
			Details: "Root login is disabled or restricted.",
		})
	}

	// Check 2: SSH Password Authentication
	if isPasswordAuthEnabled() {
		checks = append(checks, Check{
			Type:       "ssh_password_auth",
			Status:     "warn",
			Details:    "Password authentication is enabled. Consider using SSH keys only for better security.",
			FixCommand: "sudo sed -i 's/PasswordAuthentication yes/PasswordAuthentication no/g' /etc/ssh/sshd_config && sudo systemctl restart ssh",
		})
		points -= 20
	} else {
		checks = append(checks, Check{
			Type:    "ssh_password_auth",
			Status:  "pass",
			Details: "Password authentication is disabled.",
		})
	}

	// Check 3: UFW Firewall
	if !isUfwActive() {
		checks = append(checks, Check{
			Type:       "firewall",
			Status:     "fail",
			Details:    "UFW (Uncomplicated Firewall) is inactive or not installed. Your server might be exposing unnecessary ports.",
			FixCommand: "sudo ufw allow 22/tcp && sudo ufw --force enable",
		})
		points -= 30
	} else {
		checks = append(checks, Check{
			Type:    "firewall",
			Status:  "pass",
			Details: "UFW is active.",
		})
	}

	if points < 0 {
		points = 0
	}

	return checks, points
}

func isRootLoginEnabled() bool {
	// Common locations for sshd_config
	content, err := os.ReadFile("/etc/ssh/sshd_config")
	if err != nil {
		return false // Can't read, assume safe or handle error
	}
	
	// Very simple check
	return bytes.Contains(content, []byte("PermitRootLogin yes"))
}

func isPasswordAuthEnabled() bool {
	content, err := os.ReadFile("/etc/ssh/sshd_config")
	if err != nil {
		return true // Default is often yes
	}
	
	return bytes.Contains(content, []byte("PasswordAuthentication yes"))
}

func isUfwActive() bool {
	cmd := exec.Command("ufw", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}
	return bytes.Contains(out.Bytes(), []byte("Status: active"))
}


func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}

func sendHeartbeat(apiKey string, payload Payload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned status: %s", resp.Status)
	}

	return nil
}
