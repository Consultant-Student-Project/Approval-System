package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"

	"github.com/Approval-System/handlers"
)

func main() {

	l := log.New(os.Stdout, "Approval-System", log.LstdFlags)

	sm := mux.NewRouter()

	ch := handlers.NewApplications(l)
}
