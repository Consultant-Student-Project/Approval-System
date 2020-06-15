package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Consultant-Student-Project/Approval-System/data"
)

type Applications struct {
	l *log.Logger
}

func NewApplications(l *log.Logger) *Applications {
	return &Applications{l}
}

func (a *Applications) GetApplications(rw http.ResponseWriter, r *http.Request) {
	// Get Applications
}

func (a *Applications) AddApplication(rw http.ResponseWriter, r *http.Request) {
	// TODO: Add Application
}

type KeyApplication struct{}

func (a Applications) MiddlewareValidateApplication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		application := data.Application{}

		err := application.FromJSON(r.Body)
		if err != nil {
			a.l.Println("[ERROR] deserializing comment", err)
			http.Error(rw, "Error reading comment", http.StatusBadRequest)
			return
		}

		err = application.Validate()
		if err != nil {
			a.l.Println("[ERROR] validating comment ", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating comment: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyApplication{}, application)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
