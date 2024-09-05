package handler

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouter() {
	http.HandleFunc("/create_event", h.Create)
	http.HandleFunc("/update_event", h.Update)
	http.HandleFunc("/delete_event", h.Delete)
	http.HandleFunc("/events_for_day", h.GetForDay)
	http.HandleFunc("/events_for_week", h.GetForWeek)
	http.HandleFunc("/events_for_month", h.GetForMonth)
}
