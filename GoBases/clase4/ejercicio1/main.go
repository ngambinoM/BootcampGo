package main

import "fmt"

func main()  {
	 producto4 := Product{2,"Pizza",30.0,"Pizza con queso y jamon, extra crujiente","Comida"}
	producto4.Save()
	producto4.GetAll()
}

type Product struct{
	ID 			int
	Name 		string
	Price 		float64
	Description string
	Category	string
}
var Products = []Product{
	{1,"Papas",12.0,"Papas fritas muy crujientes","Comida"},
	{2,"Pizza",30.0,"Pizza con queso y jamon, extra crujiente","Comida"},
}

func  (p *Product) Save(){
	Products = append(Products, *p)
}

func  (p *Product) GetAll() {
	for _,p:= range Products{
		fmt.Printf("Product con Id: %d, Name: %s, Price: %f, Description: %s, Category: %s\n",p.ID,p.Name,p.Price,p.Description,p.Category)
	}
}

func getById (Id int) *Product{
	for i := range Products {
        if Products[i].ID == Id {
            return &Products[i]
        }
    }
    return nil
}