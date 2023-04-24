package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)




func main() {
	tickets.ExtractData()
	data,err := tickets.ExtractData()
	if err!=nil {
		panic(err)
	}
	//people := tickets.GetTotalTickets("China", data)

	time,err := tickets.AverageDestination("China",data)
	if err!=nil {
		panic(err)
	}
	fmt.Println(time,"%")
}
