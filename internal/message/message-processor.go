package message

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/protogen/v1/message/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/all", GetAllMessages).Methods("GET")
	router.HandleFunc("/{id}", FindMessageById).Methods("GET")
}

func GetAllMessages(writer http.ResponseWriter, request *http.Request) {
	messages := getMessages()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(messages)
	if err != nil {
		log.Println("error occurred while fetching all messages", err)
		return
	}
	writer.Write(jsonResponse)
}

// Create some dummy messages
func getMessages() []proto.Message {
	messages := make([]proto.Message, 0)
	messages = append(messages, proto.Message{Id: "MSG00", Message: "hello world", GeneratorId: "sbaitmangalkar", CreateTime: timestamppb.New(time.Now())})
	messages = append(messages, proto.Message{Id: "MSG01", Message: "hello mars", GeneratorId: "emusk", CreateTime: timestamppb.New(time.Now())})
	return messages
}

func FindMessageById(writer http.ResponseWriter, request *http.Request) {
	messages := getMessages()
	vars := mux.Vars(request)
	messageId := vars["id"]
	var result proto.Message
	for _, message := range messages {
		if message.GetId() == messageId {
			result = message
			break
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Println("error occurred while fetching all messages", err)
		return
	}
	writer.Write(jsonResponse)
}
