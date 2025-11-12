package model

import "time"

type Completion struct {
	ID               int       `json:"id"`
	Firstname        string    `json:"firstname"`
	Lastname         string    `json:"lastname"`
	IDPassportNumber *string   `json:"id_passport_number,omitempty"`
	Company          *string   `json:"company,omitempty"`
	Date             time.Time `json:"date"`
}
