package importer

import (
	"log"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
	"github.com/plaenkler/gost/pkg/config"
	"github.com/plaenkler/gost/pkg/handler"
)

var instance *Manager

type Manager struct {
	config *config.Config
}

func GetManager() *Manager {
	defer handler.HandlePanic("importer")

	if instance == nil {
		instance = &Manager{
			config: config.GetConfig(),
		}
	}

	return instance
}

func (manager *Manager) Start() {
	defer handler.HandlePanic("importer")
}

func (manager *Manager) Import() {
	defer handler.HandlePanic("importer")

	// Open file
	dbf, err := dbase.Open(manager.config.PathToDB, new(dbase.Win1250Converter))
	if err != nil {
		log.Fatalf("[importer] connection failed - error: %s", err.Error())
	}
	defer dbf.Close()
}
