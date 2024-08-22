package main

import (
	"log"
	"net/http"

	"95HW/config"
	"95HW/handlers"
	"95HW/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config yuklashda xatolik: %v", err)
	}
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Ma'lumotlar bazasiga ulanishda xatolik: %v", err)
	}

	http.HandleFunc("/webhook", handlers.WebhookHandler(db))

	log.Println("Server 8080-portda ishga tushdi...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}
}
