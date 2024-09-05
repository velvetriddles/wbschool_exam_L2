package app

import (
	"dev11/internal/handler"
	"log"
	"net/http"
)

func Run() error {
	h := handler.NewHandler()
	h.InitRouter()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	return nil
}
