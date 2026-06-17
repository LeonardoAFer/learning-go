package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       int
	createdAt time.Time
}

func main() {
	var appUser User

	appUser = User{
		firstName: "John",
		lastName:  "Doe",
		age:       42,
		createdAt: time.Now(),
	}

	fmt.Println(appUser)
}
