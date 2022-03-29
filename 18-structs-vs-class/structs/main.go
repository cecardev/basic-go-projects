package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id

}

func (e *Employee) setName(name string) {
	e.name = name
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{id, name}
}

func main() {
	var emp Employee
	emp.id = 1
	emp.name = "John"
	fmt.Println(emp)
	emp.setId(2)
	fmt.Println(emp)
	emp.setName("Jane")
	fmt.Println(emp)

	//Constructor 1
	emp2 := Employee{}
	fmt.Println(emp2)
	//Constructor 2
	var emp3 = Employee{id: 3, name: "John"}
	fmt.Println(emp3)
	//Constructor 3
	var emp4 = new(Employee)
	fmt.Println(emp4)
	//Constructor 4
	emp5 := NewEmployee(5, "Jane")
	fmt.Println(emp5)

}
