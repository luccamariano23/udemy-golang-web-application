package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println(i)

	whatWasSaid, otherThingSaid := saySomething()
	fmt.Println(whatWasSaid, otherThingSaid)
}

func saySomething() (string, string) {
	return "something", "another"
}
