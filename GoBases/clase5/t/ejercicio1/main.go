package main

import (
	"fmt"
	"os"
)

func main()  {
	defer func () {
		fmt.Println("ejecución finalizada")
		
	}()
	file,err:= os.Open("customers.txt")
	if err != nil{
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	defer file.Close()
	fmt.Println(file)
	
}
