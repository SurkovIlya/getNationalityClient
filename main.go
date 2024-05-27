package main

import (
	"ezclient/internal/nationalpredict"
	"ezclient/internal/nationalsource"
)

const host = "https://api.nationalize.io"

func main() {
	ns := nationalsource.New(host)

	np := nationalpredict.New(nil, ns)

	service := service.New(np)

	service.Start()
}
