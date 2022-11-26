package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"mongoapi/models"
)

type UserController struct {
	collection *mongo.Collection
}

func NewUserController(c *mongo.Collection) *UserController {
	return &UserController{c}
}

// below method needed
// GetUser
// CreateUser
// DeleteUser

// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
func (db *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var u models.User
	//objectID, _ := primitive.ObjectIDFromHex(vars["id"])	// somehow not work
	objectID := vars["id"] // use id directly to retrieve record, http://localhost:8070/user/1
	filter := bson.M{"_id": objectID}

	fmt.Println(filter, objectID[:])
	err := db.collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (db *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &u)
	fmt.Println(string(postBody[:]))
	result, err := db.collection.InsertOne(context.TODO(), u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}

func (db *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var u models.User

	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &u)

	objectID := vars["id"]
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": &u}
	result, err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}

func (db *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectID := vars["id"]
	filter := bson.M{"_id": objectID}

	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}
