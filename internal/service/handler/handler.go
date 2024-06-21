package handler

import (
	"encoding/json"
	"getNationalClient/internal/exception"
	"getNationalClient/internal/model"
	"getNationalClient/internal/service"
	"mime"
	"net/http"
	"time"
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
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")
	var (
		answer model.Answer
		err    error
	)
	name := r.URL.Query().Get("name")
	answer.Result, err = h.services.NationalName(name)
	if err != nil {
		json.NewEncoder(w).Encode(&model.User{})

		return
	}
	endTime := time.Now()
	elepsedTime := endTime.Sub(startTime)
	answer.Time = elepsedTime.String()
	answer.Status = http.StatusText(200)

	json.NewEncoder(w).Encode(answer)
}

func (h *Handler) AddExcention(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")

	type Response struct {
		Resp    string `json:"resp"`
		Time    string `json:"time"`
		Message string `json:"message"`
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
	var rp exception.ExceptionPerson
	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	er := h.services.Exception.AddExcStore(rp)
	endTime := time.Now()
	elepsedTime := endTime.Sub(startTime).String()
	if er != nil {
		json.NewEncoder(w).Encode(&Response{Resp: http.StatusText(200), Time: elepsedTime, Message: er.Error()})

		return
	}

	json.NewEncoder(w).Encode(&Response{Resp: http.StatusText(200), Time: elepsedTime})
}
