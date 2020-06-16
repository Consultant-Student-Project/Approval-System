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
	a.l.Println("Handle GET Applications")

	al := data.GetApplications()

	err := al.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}

}

func (a *Applications) AddApplication(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle POST Application")

	application := r.Context().Value(KeyApplication{}).(data.Application)
	data.AddApplication(&application)
}

type KeyApplication struct{}

func (a Applications) MiddlewareValidateApplication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		application := data.Application{}

		err := application.FromJSON(r.Body)
		if err != nil {
			a.l.Println("[ERROR] deserializing application", err)
			http.Error(rw, "Error reading application", http.StatusBadRequest)
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
