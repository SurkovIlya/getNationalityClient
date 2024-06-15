package handler

import (
	"encoding/json"
	"getNationalClient/internal/exception"
	"getNationalClient/internal/model"
	"getNationalClient/internal/service"
	"mime"
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
	mux.HandleFunc("GET /national", h.GetNationalName)
	mux.HandleFunc("POST /addexcention", h.AddExcention)

	return mux
}

func (h *Handler) GetNationalName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		answer model.User
		err    error
	)
	name := r.URL.Query().Get("name")
	answer, err = h.services.NationalName(name)
	if err != nil {
		json.NewEncoder(w).Encode(&model.User{})

		return
	}
	json.NewEncoder(w).Encode(answer)
}

func (h *Handler) AddExcention(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type ResponseId struct {
		Resp string `json:"resp"`
	}

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)

		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rp exception.ExcentionPerson
	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	h.services.Exception.AddExcStore(rp)
	json.NewEncoder(w).Encode(&ResponseId{Resp: "OK"})
}
