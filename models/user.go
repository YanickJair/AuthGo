package models

// User - model
type User struct {
	ID    string `json:"id"`
	FName string `json:"fname"`
	Email string `json:"email"`
}

var yannick = &User{"1", "Yanick", "Yanick@gmail.com"}
var gabriel = &User{"1", "Giovanni", "gabriel@gmail.com"}

var users = map[string]*User{"1": yannick, "2": gabriel}

// GetUsers - objects
func GetUsers() *User {
	return users["1"]
}

// SignUp - register new user
func SignUp(email string, fname string) *User {
	num := len(users) + 1

	newUser := &User{
		ID:    string(num),
		Email: email,
		FName: fname,
	}
	users[newUser.ID] = newUser
	return newUser
}

// GetUser - return a user
func GetUser(id string) *User {
	if user, ok := users[id]; ok {
		return user
	}
	return nil
}
