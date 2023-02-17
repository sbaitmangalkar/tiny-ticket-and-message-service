package message

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/internal/model"
	"log"
	"net/http"
	"time"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/all", GetAllMessages).Methods("GET")
	//router.HandleFunc("/{id}", FindMessageById).Methods("GET")
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := getMessages()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(messages)
	if err != nil {
		log.Println("error occurred while fetching all messages", err)
		return
	}
	w.Write(jsonResponse)
}

func getMessages() []model.GenericMessage {
	messages := make([]model.GenericMessage, 0)
	messages = append(messages, model.GenericMessage{Id: "MSG00", Message: "hello world", Generator: "sbaitmangalkar", CreatedAt: time.Now()})
	messages = append(messages, model.GenericMessage{Id: "MSG01", Message: "hello mars", Generator: "emusk", CreatedAt: time.Now()})
	return messages
}

func FindMessageById(w http.ResponseWriter, r *http.Request) {

}
