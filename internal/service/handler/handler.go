package handler

import (
	"encoding/json"
	"getNationalClient/internal/auth"
	"getNationalClient/internal/exception"
	"getNationalClient/internal/model"
	"getNationalClient/internal/service"
	"mime"
	"net/http"
	"time"
)

type Handler struct {
	auth     *auth.Auth
	services *service.Service
}

func NewHandler(services *service.Service, auth *auth.Auth) *Handler {
	return &Handler{
		services: services,
		auth:     auth,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /registration", h.Registration)
	mux.HandleFunc("POST /authorization", h.Authorization)
	mux.HandleFunc("GET /national", h.GetNationalName)
	mux.HandleFunc("POST /addexcention", h.AddExcention)
	mux.HandleFunc("POST /delexception", h.DelExcention)

	return mux
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")
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
	var rp model.Reg
	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	dataUser, er := h.auth.Registration(rp)
	endTime := time.Now()
	elepsedTime := endTime.Sub(startTime).String()
	if er != nil {
		json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Message: er.Error()})

		return
	}
	json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Result: dataUser})
}

func (h *Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")
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
	var rp model.Auth
	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	token, er := h.auth.Authorization(rp)
	endTime := time.Now()
	elepsedTime := endTime.Sub(startTime).String()
	if er != nil {
		json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Message: er.Error()})

		return
	}

	json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Result: token})
}

func (h *Handler) GetNationalName(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	// w.Header().Set("Content-Type", "application/json")
	var (
		answer model.Answer
		err    error
	)

	authorization := r.Header.Get("Authorization")
	err = h.auth.CheckToken(authorization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	name := r.URL.Query().Get("name")
	answer.Result, err = h.services.NationalName(name)
	if err != nil {
		endTime := time.Now()
		elepsedTime := endTime.Sub(startTime)
		answer.Time = elepsedTime.String()
		answer.Status = http.StatusText(200)
		answer.Message = err.Error()

		json.NewEncoder(w).Encode(answer)

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

	authorization := r.Header.Get("Authorization")
	err = h.auth.CheckToken(authorization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

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
		json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Message: er.Error()})

		return
	}

	json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime})
}

func (h *Handler) DelExcention(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")

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

	authorization := r.Header.Get("Authorization")
	err = h.auth.CheckToken(authorization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rp exception.ExceptionPerson
	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	er := h.services.Exception.DelException(rp.Name)
	endTime := time.Now()
	elepsedTime := endTime.Sub(startTime).String()
	if er != nil {
		json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime, Message: er.Error()})

		return
	}

	json.NewEncoder(w).Encode(&model.Answer{Status: http.StatusText(200), Time: elepsedTime})
}
