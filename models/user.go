package models

import "fmt"

// User - model
type User struct {
	ID    string `json:"id"`
	FName string `json:"fname"`
	Email string `json:"email"`
}

var yannick = User{"1", "Yanick", "Yanick@gmail.com"}
var gabriel = User{"1", "Giovanni", "gabriel@gmail.com"}

var users []User

func int() {
	users = append(users, yannick)
	users = append(users, gabriel)
}

// GetUsers - objects
func GetUsers() []User {
	users = append(users, yannick)
	users = append(users, gabriel)
	fmt.Println(users)
	return users
}

// SignUp - register new user
func SignUp(email string, fname string) *User {
	num := len(users) + 1

	newUser := User{
		ID:    string(num),
		Email: email,
		FName: fname,
	}
	users = append(users, newUser)
	return &newUser
}

// GetUser - return a user
func GetUser(id string) User {
	for _, user := range users {
		if id == user.ID {
			return user
		}
	}

	return User{}
}
