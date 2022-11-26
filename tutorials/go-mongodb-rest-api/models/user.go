package models

type User struct {
	// Id     bson.ObjectId `json:"id" bson:"_id"`	// this data type causes duplicate key error
	// Id     primitive.ObjectID `json:"id" bson:"_id"`		// this data type causes duplicate key error
	// Id     string `json:"id,omitempty" bson:"_id,omitempty"`
	// this works, can insert unique key, but query id not work,
	// make it simple, use primitive
	Id     interface{} `json:"id" bson:"_id,omitempty"`
	Name   string      `json:"name" bson:"name"`
	Gender string      `json:"gender" bson:"gender"`
	Age    int         `json:"age" bson:"age"`
}
