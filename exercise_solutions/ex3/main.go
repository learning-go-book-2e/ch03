package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"Person", "Personson", 12345}
	emp2 := Employee{
		firstName: "Mary",
		lastName:  "Jones",
		id:        65432,
	}
	var emp3 Employee
	emp3.firstName = "Pat"
	emp3.lastName = "Carter"
	emp3.id = 98765

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
