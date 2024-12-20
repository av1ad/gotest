package main

import "fmt"

func Hello(name string, age int, location string) string {

	message := fmt.Sprintf("Greetings, %v! You are %v years old and located in %v", name, age, location)
	fmt.Println(message)
	return message

}

func main() {
	Hello("Aviad", 21, "Florida")
}
