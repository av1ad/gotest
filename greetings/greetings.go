package main

import (
	"bufio"
	"fmt"
	"os"
)

func Hello() string {

	var age int

	fmt.Print("Enter full name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your location: ")
	location, _ := reader.ReadString('\n')
	location = location[:len(location)-1]

	message := fmt.Sprintf("Greetings %s! You are %d years old and located in %s", name, age, location)
	fmt.Println(message)

	return message

}

func main() {
	Hello()
}
