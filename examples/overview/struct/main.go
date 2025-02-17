package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName string
	LastName  string
	Phone     string
	Age       int
	Birth     time.Time
}

func main() {
	user := User{
		FirstName: "Lucca",
		LastName:  "Mariano",
	}

	log.Println(user.FirstName, user.LastName)
}
