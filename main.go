package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
func run(ctx context.Context) error {
	log.Print("Starting Copr-Hook at port 7070\n")
	return http.ListenAndServe(":7070", http.HandlerFunc(handler))

}
func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	payload := make(map[string]interface{})
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Fatal(err)
	}
	if _, has := payload["ref"]; !has {
		log.Fatal("Payload not received!")
	}
	ref, ok := payload["ref"].(string)
	if !ok {
		log.Fatal("Expected string")
	}
	refs := strings.Split(ref, "/")
	switch refs[2] {
	case "master":
		url := os.Getenv("MASTER")
		res, err := http.Post(url, "application/json", nil)
		if err != nil {
			log.Fatal("Adress Not Reachable")
		}
		if res.Status == "200" {
			log.Println("Success")
		}

	case "v0.10":
		url := os.Getenv("V10")
		res, err := http.Post(url, "application/json", nil)
		if err != nil {
			log.Fatal("Adress Not Reachable")
		}
		if res.Status == "200" {
			log.Println("Success")
		}
	}

}
