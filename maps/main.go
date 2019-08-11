package main

import "fmt"

func main() {

	// Define maps for email in two ways
	emailList := make(map[string]string)
	newEmailLits := map[string]string{"new": "new@new.com", "old": "old@old.com"}

	// Add data
	emailList["Dhan"] = "me@dhanvsagar.in"
	emailList["DVS"] = "dvs@dhanvsagar.in"

	// Retreive Data
	fmt.Println(len(emailList))
	fmt.Println(emailList)
	fmt.Println(emailList["Dhan"])

	fmt.Println(newEmailLits)

	// Using range() to print all

	for key, value := range emailList {
		fmt.Printf("%s : %s \n", key, value)
	}

	for key := range emailList {
		fmt.Println("Name : " + key)
	}

}
