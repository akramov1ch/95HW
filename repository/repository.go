package repository

import (
	"log"

	"95HW/config"
	"95HW/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Avtomatik migratsiya
	if err := db.AutoMigrate(&model.RepositoryEvent{}); err != nil {
		return nil, err
	}

	log.Println("Ma'lumotlar bazasi ulanishi muvaffaqiyatli!")
	return db, nil
}

func SaveRepositoryEvent(db *gorm.DB, event *model.RepositoryEvent) error {
	return db.Create(event).Error
}
