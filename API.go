package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// event strcture
type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}

// ticket strctre
type Ticket struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Code    string `json:"code"`
}

// handler(definef ahead)
type EventHandler struct {
}

func RegisterRoutes(router *mux.Router, handler *EventHandler) {
	//create event
	router.HandleFunc("/api/events", handler.CreateEvent).
		Methods(http.MethodPost).
		Name("CreateEvent")

	//event by id
	router.HandleFunc("/api/events/{id}", handler.GetEventByID).
		Methods(http.MethodGet).
		Name("GetEventByID")

	//get all event
	router.HandleFunc("/api/events", handler.GetAllEvents).
		Methods(http.MethodGet).
		Name("GetAllEvents")

	//arequest
	router.HandleFunc("/api/events/{id}/approve", handler.ApproveRequest).
		Methods(http.MethodPut).
		Name("ApproveRequest")

	// rrequest
	router.HandleFunc("/api/events/{id}/reject", handler.RejectRequest).
		Methods(http.MethodPut).
		Name("RejectRequest")

	//  ticketg
	router.HandleFunc("/api/events/{id}/tickets", handler.GenerateTicket).
		Methods(http.MethodPost).
		Name("GenerateTicket")
}

// handler funs
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
}
func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
}
func (h *EventHandler) ApproveRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
}
func (h *EventHandler) RejectRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
}
func (h *EventHandler) GenerateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
}
func main() {
	router := mux.NewRouter()
	handler := &EventHandler{}
	RegisterRoutes(router, handler)
	http.ListenAndServe(":8080", router)
}
