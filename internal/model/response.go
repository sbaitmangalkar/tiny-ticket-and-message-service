package model

import "time"

type GenericMessage struct {
	Id        string    `json:"id"`
	Message   string    `json:"message"`
	Generator string    `json:"generator"`
	CreatedAt time.Time `json:"createdAt"`
}
