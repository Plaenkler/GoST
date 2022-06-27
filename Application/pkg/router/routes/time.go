package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/plaenkler/gost/pkg/handler"
)

var instance *Session

type timePeriod struct {
	start time.Time
	end   time.Time
}

type Session struct {
	timePeriod *timePeriod
}

func GetSession() *Session {
	defer handler.HandlePanic("routes")

	if instance == nil {
		instance = &Session{
			timePeriod: &timePeriod{
				start: time.Now().Add(-8760 * time.Hour),
				end:   time.Now(),
			},
		}
	}

	return instance
}

func ProvideTimePage(writer http.ResponseWriter, request *http.Request) {
	defer handler.HandlePanic("routes")

	log.Printf("[providetimepage] request on /time/")

	if request.Method != "POST" {
		writer.WriteHeader(403)
		writer.Write([]byte("405 - Method not allowed"))
	}

	if err := request.ParseForm(); err != nil {
		log.Printf("[providetimepage] could not parse form err: %s", err.Error())
		ProvideHomePage(writer, request)
	}

	start := request.FormValue("start")
	if len(start) == 0 {
		log.Printf("[providetimepage] value start is not present")
		ProvideHomePage(writer, request)
	}

	end := request.FormValue("end")
	if len(end) == 0 {
		log.Printf("[providetimepage] value end is not present")
		ProvideHomePage(writer, request)
	}

	dateStart, err := time.Parse("2006-01-02", start)
	if err != nil {
		log.Printf("[providetimepage] could not convert start to time - error: %s", err.Error())
		ProvideHomePage(writer, request)
	}

	dateEnd, err := time.Parse("2006-01-02", end)
	if err != nil {
		log.Printf("[providetimepage] could not convert end to time - error: %s", err.Error())
		ProvideHomePage(writer, request)
	}

	instance.timePeriod.start = dateStart
	instance.timePeriod.end = dateEnd
	ProvideHomePage(writer, request)
}
