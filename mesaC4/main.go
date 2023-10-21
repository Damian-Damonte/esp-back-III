package main

import "fmt"

type Pequenio struct {
	Costo  float32
}

func (p Pequenio) Precio() float32 {
	return p.Costo
}

type Mediano struct {
	Costo  float32
}

func (m Mediano) Precio() float32 {
	return m.Costo * 1.03
}

type Grande struct {
	Costo  float32
}

func (g Grande) Precio() float32 {
	return g.Costo*1.06 + 2500
}

type Producto interface{
	Precio() float32
}

func ProductoFactory(tipo string, precio float32) Producto {
	switch tipo {
	case "pequeño":
		return Pequenio{precio}
	case "mediano":
		return Mediano{precio}
	case "grande":
		return Grande{precio}
	default:
		return nil
	}
}

func main() {
	productoPequenio := ProductoFactory("pequeño", 100)
	productoMediano := ProductoFactory("mediano", 100)
	productoGrande := ProductoFactory("grande", 100)

	fmt.Printf("Precio producto pequeño: %.2f \n", productoPequenio.Precio())
	fmt.Printf("Precio producto mediano: %.2f \n", productoMediano.Precio())
	fmt.Printf("Precio producto grande: %.2f \n", productoGrande.Precio())
}