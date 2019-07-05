package main

import (
	"fmt"

	"github.com/YanickJair/AuthGo/models"
)

func main() {
	yannick := models.User{
		FirstName: "Yannick",
		LastName:  "Andrade",
		Password:  "newPassw0rd",
		Email:     "yanick.jair.ta@gmail.com",
	}
	yannick.SignUp()
	fmt.Println(yannick.ID)
}
