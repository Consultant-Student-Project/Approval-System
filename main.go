package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Consultant-Student-Project/Approval-System/handlers"
)

func main() {

	l := log.New(os.Stdout, "Approval-System", log.LstdFlags)

	sm := mux.NewRouter()

	ch := handlers.NewApplications(l)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{postId:[0-9]+}", ch.GetApplications)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ch.AddApplication)
	postRouter.Use(ch.MiddlewareValidateApplication)
}
