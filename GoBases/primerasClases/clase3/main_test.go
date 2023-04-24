package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuesto(t *testing.T){
	t.Run("Sueldo menor a 50000", func (t *testing.T)  {
		sueldo:=45000.00
	resultadoEsperado:= 0.00

	resultado:= impuestoSalario(sueldo)

	assert.Equal(t,resultadoEsperado,resultado)
	})
	t.Run("Sueldo mayor a 50000 menor a 150000", func (t *testing.T)  {
		sueldo:=55000.00
	resultadoEsperado:= sueldo*0.17

	resultado:= impuestoSalario(sueldo)

	assert.Equal(t,resultadoEsperado,resultado)
	})
	t.Run("Sueldo mayor  a 150000", func (t *testing.T)  {
		sueldo:=155000.00
	resultadoEsperado:= sueldo*0.27

	resultado:= impuestoSalario(sueldo)

	assert.Equal(t,resultadoEsperado,resultado)
	})
}

func TestCalcularPromedio(t *testing.T){
	t.Run("promedio de notas", func (t *testing.T)  {
		var nota1 float32=10
		var nota2 float32=8
		var nota3 float32=9
		var resultadoEsperado float32 =9.0

		resultado:= calcularPromedio(nota1,nota2,nota3)

		assert.Equal(t,resultadoEsperado,resultado)
	})
}

func TestCalcularSalario(t *testing.T){
	t.Run("Categoria A", func (t *testing.T)  {
		minutos:=60
		categoria:="A"
		resultadoEsperado:=4500.00

		resultado:=calcularSalario(minutos,categoria)

		assert.Equal(t,resultadoEsperado,resultado)
	})
	t.Run("Categoria B", func (t *testing.T)  {
		minutos:=60
		categoria:="B"
		resultadoEsperado:=1.0 * 1500 + (1.0 * 1500 * 0.2)

		resultado:=calcularSalario(minutos,categoria)

		assert.Equal(t,resultadoEsperado,resultado)
	})
	t.Run("Categoria C", func (t *testing.T)  {
		minutos:=60
		categoria:="C"
		resultadoEsperado:=1.0 * 1000

		resultado:=calcularSalario(minutos,categoria)

		assert.Equal(t,resultadoEsperado,resultado)
	})
}
func TestCalcularEstadisticas(t *testing.T){
	t.Run("Minima de notas", func (t *testing.T)  {
		nota1 :=10.0
		nota2 :=8.0
		nota3 :=9.0
		nota4 :=2.0
		resultadoEsperado :=2.0
		operacion:="minimo"
		ope,err := orquestador(operacion)
		if err != nil{
		resultado:=ope(nota1,nota2,nota3,nota4) 

		assert.Equal(t,resultadoEsperado,resultado)
		}
	})
	t.Run("Maxima de notas", func (t *testing.T)  {
		nota1 :=10.0
		nota2 :=8.0
		nota3 :=9.0
		nota4 :=2.0
		resultadoEsperado :=10.0
		operacion:="maximo"
		ope,err := orquestador(operacion)
		if err != nil{
		resultado:=ope(nota1,nota2,nota3,nota4) 

		assert.Equal(t,resultadoEsperado,resultado)
		}
	})
}
