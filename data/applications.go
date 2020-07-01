package data

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
)

type Application struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	ImageURL   string `json:"imageURL"`
	Done       bool   `json:"done"`
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
func SetApplicationDone(postID int) {
	_, err := Collection.UpdateOne(
		context.TODO(),
		bson.M{"id": postID},
		bson.D{
			{"$set", bson.D{{"done", true}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
func FindApplications() Applications {
	var results Applications
	cur, err := Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var app Application
		err := cur.Decode(&app)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &app)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results
}

// GetComments : Find every comment in database
func GetApplications() Applications {
	applications := FindApplications()

	return applications
}

// AddComment : Add comment to database
func AddApplication(a *Application) {
	a.ID = getNextID()
	a.CreatedOn = time.Now().UTC().String()
	a.Done = false
	_, err := Collection.InsertOne(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
	}
}
func getNextID() int {
	dbSize, err := Collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return (int)(dbSize + 1)
}
