package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

type CommentService interface {
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.MapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) MapRoutes() {
	h.Router.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello world")
		},
	)

}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
