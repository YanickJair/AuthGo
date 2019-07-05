package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/bsontype"
	"github.com/YanickJair/AuthGo/lib"
)

// User - model
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

// SignUp - register a new user
func (user User) SignUp() *User {
	client := lib.DB()

	collection := client.Database("authGo").Collection("user")
	// yannick := models.User{"Yannick", "yanick.jair.ta@gmail.com"}
	res, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}

	user.ID = res.InsertedID
	fmt.Println(user.ID)
	return &user
}
