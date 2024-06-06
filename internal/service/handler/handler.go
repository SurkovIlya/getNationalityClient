package handler

import (
	"encoding/json"
	"getNationalClient/internal/model"
	"getNationalClient/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/national", h.GetNationalName)

	return mux
}

func (h *Handler) GetNationalName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var answer model.UserRespons
	var err error
	name := r.URL.Query().Get("name")
	answer.Name = name
	answer.National, err = h.services.NationalName(name)

	if err != nil {
		json.NewEncoder(w).Encode(&model.User{})

		return
	}
	json.NewEncoder(w).Encode(answer)

}
