package main

import (
	"strconv"
)

//Address is a struct representing an ssh address
type Address struct {
	username string `json:"username"`
	ip       string `json:"ip"`
	port     int    `json:"port"`
	desc     string `json:"desc"`
}

func (a Address) express() string {
	argString := ""
	if a.username != "" {
		argString = argString + a.username + "@"
	}
	argString = argString + a.ip
	if a.port != 22 {
		argString = argString + ":" + strconv.Itoa(a.port)
	}
	return argString
}

//AddressBook is a struct that stores many Address instances
type AddressBook struct {
	list []Address `json:"list"`
}

//Add appends an Address to the AddressBook
func (a *AddressBook) Add(addr Address) {
	a.list = append(a.list, addr)
}
func main() {}
