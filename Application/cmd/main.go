package main

import (
	"github.com/plaenkler/gost/pkg/database"
	"github.com/plaenkler/gost/pkg/importer"
	"github.com/plaenkler/gost/pkg/router"
)

var (
	dbManager *database.Manager
	ipManager *importer.Manager
	rtManager *router.Manager
)

func init() {
	dbManager = database.GetManager()
	dbManager.Start()

	ipManager = importer.GetManager()
	ipManager.Start()
}

func main() {
	rtManager = router.GetManager()
	rtManager.Start()
}
