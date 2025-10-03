package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Version string `json:"version"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	ver := os.Getenv("APP_VERSION")
	if ver == "" {
		ver = "dev"
	}
	res := response{
		Message: "Hello from Go + CI/CD!",
		Time:    time.Now().Format(time.RFC3339),
		Version: ver,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
