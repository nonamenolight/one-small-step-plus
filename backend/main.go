package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../config/config.yaml"
	}

	config, err := loadConfig(configPath)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/config", configHandler(config))

	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "../frontend"
	}

	http.Handle("/", http.FileServer(http.Dir(frontendPath)))

	fmt.Println("server listening on :8081")

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
