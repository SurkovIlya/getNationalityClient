package model

import "time"

type User struct {
	Name           string
	National       string
	Lastusedgetime time.Time
}

type Answer struct {
	Status string `json:"status"`
	Time   string `json:"time"`
	Result User   `json:"result"`
}
