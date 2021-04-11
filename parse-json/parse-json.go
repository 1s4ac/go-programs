package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Users struct, array of users
type Users struct {
	Users []User `json:"users"`
}

//User struct
type User struct {
	Forename string   `json:"forename"`
	Surname  string   `json:"surname"`
	Age      int      `json:"age"`
	Location Location `json:"Location"`
}

//Location struct
type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	jsonContent, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonContent.Close()
	//read file as byte array
	content, _ := ioutil.ReadAll(jsonContent)
	//Initialize users array
	var user Users
	//unmarshal content into users
	json.Unmarshal(content, &user)
	//Iterate through each user
	for i := 0; i < len(user.Users); i++ {
		fmt.Printf(user.Users[i].Surname)
		fmt.Printf(user.Users[i].Forename)
		fmt.Printf(strconv.Itoa(user.Users[i].Age))
		fmt.Printf(user.Users[i].Location.Country + "\n")

	}

}
