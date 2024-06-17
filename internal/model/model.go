package model

import "time"

type User struct {
	Name           string
	National       string
	Lastusedgetime time.Time
}

type Answer struct {
	Result User          `json:"result"`
	Time   time.Duration `json:"time"`
}
