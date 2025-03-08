package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
)

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}

type Ticket struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Code    string `json:"code"`
}

type EventHandler struct{}

func RegisterRoutes(router *mux.Router, handler *EventHandler) {
	router.HandleFunc("/api/events", handler.CreateEvent).
		Methods(http.MethodPost).
		Name("CreateEvent")
	router.HandleFunc("/api/events/{id}", handler.GetEventByID).
		Methods(http.MethodGet).
		Name("GetEventByID")
	router.HandleFunc("/api/getevents", handler.GetAllEvents).
		Methods(http.MethodGet).
		Name("GetAllEvents")
	router.HandleFunc("/api/events/{id}/approve", handler.ApproveRequest).
		Methods(http.MethodPut).
		Name("ApproveRequest")
	router.HandleFunc("/api/events/{id}/reject", handler.RejectRequest).
		Methods(http.MethodPut).
		Name("RejectRequest")
	router.HandleFunc("/api/events/{id}/tickets", handler.GenerateTicket).
		Methods(http.MethodPost).
		Name("GenerateTicket")
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Event endpoint\n"))
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Event by ID endpoint\n"))
}

func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All Events endpoint\n"))
}

func (h *EventHandler) ApproveRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Approve Request endpoint\n"))
}

func (h *EventHandler) RejectRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reject Request endpoint\n"))
}

func (h *EventHandler) GenerateTicket(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Generate Ticket endpoint\n"))
}

func main() {
	router := mux.NewRouter()
	handler := &EventHandler{}
	RegisterRoutes(router, handler)

	// start server goroutine
	go func() {
		fmt.Println("Server started at http://localhost:8080/api/getevents")
		if err := http.ListenAndServe(":8080", router); err != nil {
			fmt.Println("Server failed:", err)
		}
	}()

	openBrowser("http://localhost:8080/api/getevents")

	// prevent main from exiting
	select {}
}

// opens at default browser
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	}

	if err := exec.Command(cmd, args...).Start(); err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}
