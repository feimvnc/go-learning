package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"mongoapi/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// dont put this in real code. Nobody will speak to you
	credentials := options.Credential{
		Username: "admin",
		Password: "admin",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("appDB").Collection("users")
	uc := controllers.NewUserController(collection)

	r := mux.NewRouter()
	r.HandleFunc("/user", uc.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id:[a-zA-Z0-9]*}", uc.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id:[a-zA-Z0-9]*}", uc.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id:[a-zA-Z0-9]*}", uc.DeleteUser).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
