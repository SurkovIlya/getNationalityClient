package model

import "time"

type User struct {
	ID             uint32
	Name           string
	National       string
	Lastusedgetime time.Time
}
