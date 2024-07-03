package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cassioglay/arquitetura-hexagonal/adapters/web/handler"
	"github.com/cassioglay/arquitetura-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewEwbServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
