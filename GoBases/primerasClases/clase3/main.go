package main

import (
	"fmt"
	"github.com/pkg/errors"
)





func main (){
	fmt.Println("El impuesto es: ", impuestoSalario(155000))
	//fmt.Println("Promedio: ", calcularPromedio(10,8,9))
	//ope,err := orquestador("promedio")
	
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println("Resultado: ", ope(10,8,8))
	//}
	animalType := "hamstger"
	animal, err := calcularAlimento(animalType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("La cantidad de alimento para %s es de %f kg \n", animalType, float64(animal(10)) / float64(1000))
	}
}
//....................................... Primer ejercicio .......................................
func impuestoSalario (sueldo float64) float64 {
	var impuesto float64;
	if sueldo < 50000 {
		impuesto = 0;
	}else{
		impuesto=(sueldo*0.17);
		if sueldo >150000 {
			impuesto+= sueldo*0.1
		}
	}
	return impuesto;
}
//....................................... Segundo ejercicio .......................................
func calcularPromedio(notas ...float32) float32 {
	var promedio,suma float32;
	

	for _, nota := range notas{
		if nota >= 0 {
		suma+=nota
	}
	}
	promedio=suma/float32(len(notas));
	return promedio;
}

//....................................... Tercer ejercicio .......................................
func convertirMinutosAHoras(minutos int) float64 {
	return float64(minutos / 60)
}


func calcularSalario (minutosMes int, categoria string) float64 {
	horas := convertirMinutosAHoras(minutosMes)
	switch categoria{
		case "A":
			return horas * 3000 + (horas * 3000 * 0.5)
		case "B":
			return horas * 1500 + (horas * 1500 * 0.2)
		case "C": 
			return horas * 1000

	}
	return 0
}

//....................................... Cuarto ejercicio .......................................


func buscarMinimo (numeros ...float64) float64{
	minimo := numeros[0]

    for _, valor := range numeros {
        if valor < minimo {
            minimo = valor
        }
    }
	return minimo;
}

func buscarMaximo (numeros ...float64) float64{
	maximo := numeros[0]

    for _, valor := range numeros {
        if valor > maximo {
            maximo = valor
        }
    }
	return maximo;
}

func buscarPromedio (numeros ...float64) float64{
	var suma float64;
	for _, valor := range numeros {
        suma+=valor;
    }
	return suma/float64(len(numeros))
}

func orquestador(tipoFuncion string) (func(notas ...float64) float64, error){
	switch tipoFuncion {
	case "minimo":
				return buscarMinimo, nil
	case "promedio":
				return buscarPromedio, nil
	case "maximo":
				return buscarMaximo, nil
	default:
		return nil,errors.New("No hay funcion de ese tipo")
	}
 
}
//....................................... Quinto ejercicio .......................................
func perroCantidad (cantidad int) int {
	return cantidad * 10000  
}

func gatoCantidad (cantidad int) int {
	return cantidad * 5000
}

func hamsterCantidad (cantidad int) int {
	return cantidad * 250 
}

func tarantulaCantidad (cantidad int) int {
	return cantidad * 150
}

func calcularAlimento(animal string) (func(cantidad int) int, error) {
	switch animal {
		case "perro": 
			return perroCantidad, nil
		case "gato":
			return gatoCantidad, nil
		case "hamster":
			return hamsterCantidad, nil
		case "tarantula":
			return tarantulaCantidad, nil
		default:
			return nil, errors.New("Ha ingresado una opción no válida")
	}
}