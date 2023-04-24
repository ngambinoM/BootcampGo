package main

import "fmt"
const (
	Small = "pequeño"
	Median = "mediano"
	Big="grande"
)

type Producto interface{
	precio() float64
}

type Pequeño struct{
	PrecioProducto float32
}
type Mediano struct{
	PrecioProducto float32
}
type Grande struct{
	PrecioProducto float32
}

func  (p Pequeño) precio() float32{
	return p.PrecioProducto
}

func  (m Mediano) precio() float32{
	return m.PrecioProducto + m.PrecioProducto*0.03
}
func  (g Grande) precio() float32{
	return g.PrecioProducto + g.PrecioProducto*0.03 +25000
}

func FactoryProduct (tipo string, precio float32) float32{
	switch tipo {
	case "pequeño":
		p := Pequeño{PrecioProducto: precio}
		return p.precio()
	case "mediano":
		m := Mediano{PrecioProducto: precio}
		return m.precio()
	case "grande":
		g := Grande{PrecioProducto: precio}
		return g.precio()
	default:
		return 0.0
		
	}
}

func main()  {
	precio := FactoryProduct(Big,100.0)
	fmt.Println("Precio: ",precio)
}