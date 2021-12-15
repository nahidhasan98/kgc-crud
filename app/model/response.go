package model

type Response struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Data    []Student `json:"data,omitempty"`
	Type    string    `json:"type,omitempty"`
}
