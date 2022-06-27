package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/plaenkler/gost/pkg/config"
	"github.com/plaenkler/gost/pkg/database/model"
	"github.com/plaenkler/gost/pkg/handler"
)

var instance *Manager

type Manager struct {
	DB     *gorm.DB
	config *config.Config
}

func GetManager() *Manager {
	defer handler.HandlePanic("database")

	if instance == nil {
		instance = &Manager{
			config: config.GetConfig(),
		}
	}

	return instance
}

func (manager *Manager) Start() {
	defer handler.HandlePanic("database")

	db, err := gorm.Open(sqlite.Open(manager.config.PathForDB))
	if err != nil {
		log.Fatalf("[database] connection failed - error: %s", err.Error())
	} else {
		log.Printf("[database] successfully connected")
		manager.DB = db
		err := manager.DB.AutoMigrate(
			&model.Agent{},
			&model.Article{},
			&model.Technician{},
			&model.Manufacturer{},
			&model.CustomerError{},
			&model.DetectedError{},
			&model.WorkshopInvoice{},
			&model.WorkshopInvoicePosition{},
			&model.WorkshopInvoiceCustomerError{},
			&model.WorkshopInvoiceDetectedError{},
		)
		if err != nil {
			log.Fatalf("[database] migration failed - error: %s", err.Error())
		}
		log.Printf("[database] successfully migrated models")
	}
}
