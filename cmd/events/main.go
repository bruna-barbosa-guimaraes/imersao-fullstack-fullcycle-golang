package main

import (
	"database/sql"
	"net/http"

	"github.com/bruna-barbosa-guimaraes/imersao-fullstack-fullcycle-golang/internal/events/infra/service"
	"github.com/bruna-barbosa-guimaraes/imersao-fullstack-fullcycle-golang/internal/events/infra/service/repository"
	"github.com/bruna-barbosa-guimaraes/imersao-fullstack-fullcycle-golang/internal/events/usecase"

	httpHandler "github.com/bruna-barbosa-guimaraes/imersao-fullstack-fullcycle-golang/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "htttp://localhost:9080/api1",
		2: "htttp://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotUseCase,
		getEventUseCase,
		buyTicketUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTicket)

	http.ListenAndServe(":8080", r)
}
