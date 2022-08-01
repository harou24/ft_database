package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

const Version = "1.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warning(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
	}
)

type Options struct {
	Logger
}

func New(dir, o) {

}

func Write() error {

}

func Read() error {

}

func ReadAll() {

}

func Delete() error {

}

func getOrCreateMutex() *sync.Mutex {

}

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
