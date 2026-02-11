package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/builders-lab/trailblazer-frontend/internal/models"
)

func (cfg *ApiConfig) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	
	// Obvious
	webhookSecret := cfg.WHSecret

	// Post only API
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Reading in memory
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	
	// Getting token
	signatureHeader := r.Header.Get("X-Hub-Signature-256")
	if signatureHeader == "" {
		http.Error(w, "Missing signature header", http.StatusForbidden)
		return
	}

	parts := strings.SplitN(signatureHeader, "=", 2)
	if len(parts) != 2 || parts[0] != "sha256" {
		http.Error(w, "Invalid signature format", http.StatusForbidden)
		return
	}
	receivedSigHex := parts[1]

	mac := hmac.New(sha256.New, []byte(webhookSecret))
	mac.Write(body)
	expectedSig := mac.Sum(nil)
	expectedSigHex := hex.EncodeToString(expectedSig)
	
	// comparing tokens
	if !hmac.Equal([]byte(receivedSigHex), []byte(expectedSigHex)) {
		log.Println("Alert! Invalid signature detected.")
		http.Error(w, "Invalid signature", http.StatusForbidden)
		return
	}

	var event models.PushEvent
	
	// Upon verification we parse it
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	
	// TODO: A common pretty printing function for all the structs 
	// Printing it
	fmt.Println("------------------------------------------------")
	fmt.Printf("⚡ Change Detected in Repo: %s\n", event.Repository.Name)
	fmt.Printf("📍 Branch: %s\n", event.Ref)
	fmt.Printf("👤 Pusher: %s\n", event.Pusher.Name)
	fmt.Println("📝 Commits:")
	for _, commit := range event.Commits {
		fmt.Printf("   - [%s] %s\n", commit.Author.Name, commit.Message)
	}
	fmt.Println("------------------------------------------------")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}
