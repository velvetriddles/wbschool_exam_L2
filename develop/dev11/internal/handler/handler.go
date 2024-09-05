package handler

import (
	"dev11/internal/service"
	"encoding/json"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, date, err := parseCreate(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.Create(userID, date)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)
	newResultResponse(w, string(res), http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	eventID, userID, date, err := parseUpdate(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.Update(eventID, userID, date)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)
	newResultResponse(w, string(res), http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	eventID, err := parseDelete(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Delete(eventID)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newResultResponse(w, "Event deleted successfully", http.StatusOK)
}

func (h *Handler) GetForDay(w http.ResponseWriter, r *http.Request) {
	date, err := parseForDay(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.GetDay(date)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)
	newResultResponse(w, string(res), http.StatusOK)
}

func (h *Handler) GetForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := parseForWeek(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.GetWeek(date)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)
	newResultResponse(w, string(res), http.StatusOK)
}

func (h *Handler) GetForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := parseForMonth(r)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.GetMonth(date)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(result)
	newResultResponse(w, string(res), http.StatusOK)
}
