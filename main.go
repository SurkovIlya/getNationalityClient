package main

import (
	"getNationalClient/internal/nationalpredict"
	"getNationalClient/internal/nationalsource"
	"getNationalClient/internal/service"
	"log"
)

const host = "https://api.nationalize.io"

func main() {
	ns := nationalsource.New(host)

	cl, err := nationalpredict.GetCountryList()
	if err != nil {
		log.Println("CountryList error:", err)
	}

	np := nationalpredict.New(cl, ns)

	sv := service.New(np)

	sv.Start()

}
