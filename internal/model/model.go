package model

import "time"

type User struct {
	Name           string
	National       string
	Lastusedgetime time.Time
}

type UserRespons struct {
	Name     string `json:"name"`
	National string `json:"national"`
}
