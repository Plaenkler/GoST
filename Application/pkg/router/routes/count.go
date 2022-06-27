package routes

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/plaenkler/gost/pkg/database"
	"github.com/plaenkler/gost/pkg/database/model"
	"github.com/plaenkler/gost/pkg/handler"
	"gorm.io/gorm/clause"
)

type countPage struct {
	Title           string
	FormatStartDate string
	FormatEndDate   string
	Revenue         uint64
	Count           uint64
	Invoices        []model.WorkshopInvoice
}

func ProvideCountPage(writer http.ResponseWriter, request *http.Request) {
	defer handler.HandlePanic("routes")

	log.Printf("[providecountpage] request on /count/")

	if request.Method != "GET" {
		writer.WriteHeader(403)
		writer.Write([]byte("405 - Method not allowed"))
	}

	session := GetSession()

	keys := request.URL.Query()
	WIType := keys["type"]
	if WIType == nil {
		log.Printf("[providecountPage] key type not provided")
	}

	WIMFR := keys["mfr"]
	if WIMFR == nil {
		log.Printf("[providecountPage] manufacturer not provided")
	}

	dbManager := database.GetManager()
	Invoices := []model.WorkshopInvoice{}
	err := dbManager.DB.Preload(clause.Associations).Preload("ReferedArticle.Manufacturer").Find(&Invoices).Error
	if err != nil {
		log.Printf("[providecountPage] could not get invoices - error: %s", err.Error())
	}

	countPage := countPage{
		Title:           "Countpage",
		FormatStartDate: session.timePeriod.start.Format("2006-01-02"),
		FormatEndDate:   session.timePeriod.end.Format("2006-01-02"),
	}

	for _, invoice := range Invoices {
		if invoice.Date.Before(session.timePeriod.start) || invoice.Date.After(session.timePeriod.end) {
			continue
		}

		if WIMFR != nil {
			MFGName := strings.ToLower(strings.TrimSpace(invoice.ReferedArticle.Manufacturer.Name))
			if MFGName != WIMFR[0] {
				continue
			}
		}

		if WIType[0] == "gr" && invoice.PaymentCondition >= 26 && invoice.PaymentCondition <= 28 {
			countPage.Invoices = append(countPage.Invoices, invoice)
			countPage.Revenue += uint64(invoice.SellPrice)
			countPage.Count++
			continue
		}

		if WIType[0] == "wr" && invoice.PaymentCondition < 26 || invoice.PaymentCondition > 28 {
			countPage.Invoices = append(countPage.Invoices, invoice)
			countPage.Revenue += uint64(invoice.SellPrice)
			countPage.Count++
		}
	}

	template, err := template.New("count").ParseFS(static, "static/html/count.html")
	if err != nil {
		log.Fatalf("[countpage] could not provide template - error: %s", err)
	}

	writer.Header().Add("Content-Type", "text/html")
	template.Execute(writer, countPage)
}
