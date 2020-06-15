package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type Application struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	ImageURL   string `json:"imageURL"`
	CreatedOn  string `json:"-"`
}

func (a *Application) Validate() error {
	validate := validator.New()

	return validate.Struct(a)
}

func (a *Application) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

type Applications []*Application

func (a *Applications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}
