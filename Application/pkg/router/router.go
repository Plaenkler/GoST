package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/plaenkler/gost/pkg/config"
	"github.com/plaenkler/gost/pkg/handler"
	"github.com/plaenkler/gost/pkg/router/routes"
)

var (
	//go:embed routes/static
	static   embed.FS
	instance *Manager
)

type Manager struct {
	Router *http.ServeMux
	config *config.Config
}

func GetManager() *Manager {
	defer handler.HandlePanic("router")

	if instance == nil {
		instance = &Manager{
			config: config.GetConfig(),
		}
	}

	return instance
}

func (manager *Manager) Start() {
	defer handler.HandlePanic("router")

	manager.Router = http.NewServeMux()

	manager.Router.HandleFunc("/",
		routes.ProvideHomePage)

	manager.Router.HandleFunc("/time/",
		routes.ProvideTimePage)

	manager.Router.HandleFunc("/count/",
		routes.ProvideCountPage)

	if err := manager.provideFiles(); err != nil {
		log.Fatalf("[router] could not provide files - error: %s", err)
	}

	if err := http.ListenAndServe(manager.config.Port, manager.Router); err != nil {
		log.Fatalf("[router] could not listen and serve - error: %s", err)
	}
}

func (manager *Manager) provideFiles() error {
	defer handler.HandlePanic("router")

	css, err := fs.Sub(static, "routes/static")
	if err != nil {
		return err
	}

	manager.Router.Handle("/css/", http.FileServer(http.FS(css)))

	img, err := fs.Sub(static, "routes/static")
	if err != nil {
		return err
	}

	manager.Router.Handle("/img/", http.FileServer(http.FS(img)))
	return nil
}
