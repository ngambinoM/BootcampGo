package main

// Vamos a hacer que nuestro programa sea un poco más complejo y útil.
// Desarrollá las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número
// negativo, la función debe devolver un error. El mismo tendrá que indicar
// “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

import (
	"errors"
	"fmt"
)

func main()  {
	horasTrabajadas:=10
	valorHora := 1510
	salario,error := calcularSalarioMensual(horasTrabajadas,valorHora)
	if error!=nil {
		fmt.Println(error)
		return 
	}
	fmt.Println(salario)
}
var (
	ErrlessThan80 = errors.New("la cantidad de horas mensuales ingresadas sea menor a 80")
)
type WorkLess struct{

}
func (e *WorkLess) Error() string{
	return ("la cantidad de horas mensuales ingresadas sea menor a 80")
}


func calcularSalarioMensual(horasTrabajadas int, valorHora int)  (float32,error){
	var salarioMensual float32
	if horasTrabajadas<80 {
		return 0,fmt.Errorf("Error: %w",ErrlessThan80)
	}
	salarioMensual = float32(horasTrabajadas) * float32(valorHora)
	if salarioMensual>= 150000 {
		salarioMensual*=0.9
	}
	return salarioMensual,nil
}