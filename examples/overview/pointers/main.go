package main

import "log"

func main() {
	var myString string
	myString = "Blue"

	log.Println(myString)

	changePointer(&myString)

	log.Println(myString)
}

func changePointer(s *string) {
	newValue := "Green"
	*s = newValue
}
