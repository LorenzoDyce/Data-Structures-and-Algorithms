package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// //Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Steve"] = "steve@gmail.com"

	// Declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails["Mike"] = "Mike@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
