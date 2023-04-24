package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
)

func main()  {

	defer func () {
		fmt.Println("\n ejecución finalizada")
	}()
	fmt.Print(ReadFile(("h.txt"))) 
	// file,err:= openFile("customer.txt")
	// if err != nil{
	// 	panic("el archivo indicado no fue encontrado o está dañado")
	// }
	// 	scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

}

func ReadFile(nameFile string) string {
	file,err := os.Open(nameFile)
	if err != nil{
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	
	defer file.Close()
	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}
	return text

	// file1,err1 := io.ReadAll(file)
	// if err1 != nil{
	// 	panic("No se puede leer")
	// }
	// return string(file1)
	// file,err := os.ReadFile(nameFile)
	// if err != nil{
	// 	panic("el archivo indicado no fue encontrado o está dañado")
	// }

	// fmt.Printf("%s", file)
}