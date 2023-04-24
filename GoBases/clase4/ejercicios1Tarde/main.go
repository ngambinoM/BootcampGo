package main

import "fmt"

type Student struct {
	Nombre 		string
	Apellido 	string
	DNI 		int
	Fecha 		string
}

func  (s *Student) Detalle(){
	fmt.Printf("Nombre: %s \nApellido: %s \nDNI: %d \nFecha: %s",s.Nombre,s.Apellido,s.DNI,s.Fecha)
}

func main()  {
	student := Student{"Nicolas","Gambino",43418456,"16-07-2001"}
	student.Detalle()
}