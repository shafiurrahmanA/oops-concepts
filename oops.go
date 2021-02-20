package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) profession() {
	fmt.Println(p.firstName, "is an Engineer")
}

func main() {
	var rahman person
	rahman.firstName = "rahman"
	rahman.lastName = "khan"
	rahman.profession()

	var rahman2 person
	rahman2.firstName = "rahman2"
	rahman2.lastName = "khan"
	rahman2.profession()

}
