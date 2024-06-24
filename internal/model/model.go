package model

import "time"

type User struct {
	Name           string
	National       string
	Lastusedgetime time.Time
}

type Answer struct {
	Status  string      `json:"status"`
	Time    string      `json:"time"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Reg struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserData struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	Hash  []byte `json:"hash"`
}
