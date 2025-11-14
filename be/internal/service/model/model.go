package model

import "time"

type Completion struct {
	ID          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	PhoneNumber string    `json:"phone_number"`
	Date        time.Time `json:"date"`
}
