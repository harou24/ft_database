package main

import (
	"encoding/json"
	"fmt"
)

const Version = "1.0"

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func New(dir, o) {

}

func main() {
	fmt.Println("HELLO")
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	employees := []User{
		{"John", "23", "888888888", "Yahoo", Address{"Amsterdam", "North Holland", "the Netherlands", "1018"}},
		{"Paul", "23", "777777777", "Google", Address{"Paris", "Ile de france", "France", "75002"}},
		{"Harou", "23", "4444444444", "Tesla", Address{"New York", "Mannathan", "USA", "1624"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Records->", records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error: ", err)
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println("All->", allusers)

	if err := db.Delete("user", "john"); err != nil {
		fmt.Println("Error: ", err)
	}

	if err := db.Delete("user", ""); err != nil {
		fmt.Println("Error->", err)
	}
}
