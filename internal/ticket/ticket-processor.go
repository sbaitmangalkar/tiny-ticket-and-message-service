package ticket

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/protogen/v1/ticket/proto"
	"log"
	"net/http"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/all", GetAllTickets).Methods("GET")
	router.HandleFunc("/{id}", FindTicketById).Methods("GET")
}

func GetAllTickets(writer http.ResponseWriter, request *http.Request) {
	tickets := createTickets()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(tickets)
	if err != nil {
		log.Println("error occurred while fetching all messages", err)
		return
	}
	writer.Write(jsonResponse)
}

// Create some dummy tickets
func createTickets() []proto.Ticket {
	tickets := make([]proto.Ticket, 0)
	issuesList0 := make([]*proto.Ticket_Issue, 0)
	issuesList0 = append(issuesList0, &proto.Ticket_Issue{Id: uuid.New().String(), Priority: proto.Ticket_HIGH, Description: "unable to access Google Drive", Assignee: "rrawat"})
	issuesList0 = append(issuesList0, &proto.Ticket_Issue{Id: uuid.New().String(), Priority: proto.Ticket_MEDIUM, Description: "unable to access Dropbox", Assignee: "rrawat"})
	tickets = append(tickets, proto.Ticket{Id: "TIK00", CreatorId: "sbaitman", CreatorName: "sbaitmangalkar", Issues: issuesList0})

	issuesList1 := make([]*proto.Ticket_Issue, 0)
	issuesList1 = append(issuesList1, &proto.Ticket_Issue{Id: uuid.New().String(), Priority: proto.Ticket_LOW, Description: "unable to access Twitter", Assignee: "emusk"})
	issuesList1 = append(issuesList1, &proto.Ticket_Issue{Id: uuid.New().String(), Priority: proto.Ticket_MEDIUM, Description: "unable to access Dropbox", Assignee: "bgates"})
	tickets = append(tickets, proto.Ticket{Id: "TIK01", CreatorId: "rrawat", CreatorName: "rrawat", Issues: issuesList1})
	return tickets
}

func FindTicketById(writer http.ResponseWriter, request *http.Request) {
	tickets := createTickets()
	vars := mux.Vars(request)
	messageId := vars["id"]
	var result proto.Ticket
	for _, message := range tickets {
		if message.GetId() == messageId {
			result = message
			break
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Println("error occurred while fetching all tickets", err)
		return
	}
	writer.Write(jsonResponse)
}
