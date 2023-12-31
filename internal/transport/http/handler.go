package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("Gracefully Shutdown")
	return nil

}
