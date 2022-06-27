package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/plaenkler/gost/pkg/database"
	"github.com/plaenkler/gost/pkg/database/model"
	"github.com/plaenkler/gost/pkg/handler"
	"gorm.io/gorm/clause"
)

var (
	//go:embed static
	static embed.FS
)

type homePage struct {
	Title                   string
	FormatStartDate         string
	FormatEndDate           string
	WarrantyInvoicesCounter int
	PayableInvoicesCounter  int
	WarrantyInvoiceCompany  map[string]int
	PayableInvoiceCompany   map[string]int
}

func ProvideHomePage(writer http.ResponseWriter, request *http.Request) {
	defer handler.HandlePanic("routes")

	log.Printf("[providehomepage] request on /")

	session := GetSession()

	dbManager := database.GetManager()
	invoices := []model.WorkshopInvoice{}
	err := dbManager.DB.Preload(clause.Associations).Preload("ReferedArticle.Manufacturer").Find(&invoices).Error
	if err != nil {
		log.Printf("[providehomepage] could not get invoices - error: %s", err.Error())
	}

	warrantyInvoices := []model.WorkshopInvoice{}
	payableInvoices := []model.WorkshopInvoice{}

	for _, invoice := range invoices {
		if invoice.Date.Before(session.timePeriod.start) {
			continue
		}
		if invoice.Date.After(session.timePeriod.end) {
			continue
		}

		if invoice.PaymentCondition >= 26 && invoice.PaymentCondition <= 28 {
			warrantyInvoices = append(warrantyInvoices, invoice)
			continue
		}

		payableInvoices = append(payableInvoices, invoice)
	}

	homePage := homePage{
		Title:           "Homepage",
		FormatStartDate: session.timePeriod.start.Format("2006-01-02"),
		FormatEndDate:   session.timePeriod.end.Format("2006-01-02"),
	}

	homePage.WarrantyInvoicesCounter = len(warrantyInvoices)
	homePage.PayableInvoicesCounter = len(payableInvoices)

	homePage.WarrantyInvoiceCompany = make(map[string]int)
	for _, warrantyInvoice := range warrantyInvoices {
		homePage.WarrantyInvoiceCompany[warrantyInvoice.ReferedArticle.Manufacturer.Name]++
	}

	homePage.PayableInvoiceCompany = make(map[string]int)
	for _, payableInvoice := range payableInvoices {
		homePage.PayableInvoiceCompany[payableInvoice.ReferedArticle.Manufacturer.Name]++
	}

	template, err := template.New("home").ParseFS(static, "static/html/home.html")
	if err != nil {
		log.Fatalf("[providehomepage] could not provide template - error: %s", err)
	}

	writer.Header().Add("Content-Type", "text/html")
	template.Execute(writer, homePage)
}
