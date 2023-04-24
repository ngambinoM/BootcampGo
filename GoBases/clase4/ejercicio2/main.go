package main

import "fmt"

func main()  {
	employee1 := Employee{1,"Jefe",Person{1,"Juan","16-07-2001"}}
	employee1.PrintEmployee()
}

type Person struct{
	IDPerson 	int
	Name 		string
	DateOfBirth string
}

type Employee struct{
	ID 			int
	Position 	string
	Person
}

func (e Employee) PrintEmployee(){
	fmt.Printf("Id: %d, Position: %s, IDPerson: %d, Name: %s, Date Of Birth: %s \n",e.ID,e.Position,e.IDPerson,e.Name,e.DateOfBirth)
}