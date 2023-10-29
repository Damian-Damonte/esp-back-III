package main

// Ejercicio PG clase 7 A usando canales

import (
	"fmt"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre         string
	Precio         float64
	MinsTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

var (
	productos      = []Producto{{"Monitor", 75000.00, 3}, {"Teclado", 20000.00, 5}, {"Mouse", 8000.00, 10}}
	servicios      = []Servicio{{"Armado de PC", 4500, 40}, {"Limpieza de PC", 3000, 25}}
	mantenimientos = []Mantenimiento{{"Mantenimiento de PC", 3500}, {"Mantenimiento de auto", 6000}}
)

func main() {
	total := 0.0
	canal := make(chan float64)

	go sumarProductos(productos, canal)
	go sumarServicios(servicios, canal)
	go sumarMantenimiento(mantenimientos, canal)

	total += <-canal
	total += <-canal
	total += <-canal
	fmt.Println(total)
}

func sumarProductos(productos []Producto, canal chan float64) {
	total := 0.0
	for _, prod := range productos {
		total += prod.Precio * float64(prod.Cantidad)
	}
	canal <- total
}

func sumarServicios(servicios []Servicio, canal chan float64) {
	total := 0.0
	for _, serv := range servicios {
		if serv.MinsTrabajados > 30 {
			total += serv.Precio * (float64(serv.MinsTrabajados) / 30)
		} else {
			total += serv.Precio
		}
	}
	canal <- total
}

func sumarMantenimiento(mantenimientos []Mantenimiento, canal chan float64) {
	total := 0.0
	for _, matn := range mantenimientos {
		total += matn.Precio
	}
	canal <- total
}
