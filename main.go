package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type healthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("starting %s on %s", serviceName(), addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func serviceName() string {
	if name := os.Getenv("SERVICE_NAME"); name != "" {
		return name
	}
	return "konflux-github-sample"
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(
		w,
		"Hello from %s!\n\nThis repo is a minimal Docker app for learning Konflux.\nTry GET /health for a JSON health check.\n",
		serviceName(),
	)
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(healthResponse{
		Status:    "ok",
		Service:   serviceName(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
