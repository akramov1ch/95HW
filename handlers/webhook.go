package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"95HW/model"
	"95HW/repository"
	"gorm.io/gorm"
)

func WebhookHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Xatolik: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		eventType := r.Header.Get("X-GitHub-Event")
		if eventType != "repository" {
			log.Printf("Tinglanmagan event turi: %s", eventType)
			http.Error(w, "Unsupported event type", http.StatusNotImplemented)
			return
		}

		var payload model.RepositoryEvent
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("JSONni unmarshallingda xatolik: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		log.Printf("Yangi repository event qabul qilindi: %v", payload)

		if err := repository.SaveRepositoryEvent(db, &payload); err != nil {
			log.Printf("Ma'lumotlar bazasiga saqlashda xatolik: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Repository event qabul qilindi")
	}
}
