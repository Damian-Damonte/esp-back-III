package main

// Ejercicio PG clase 7 A usando waitGroup y mutex

import (
	"fmt"
	"sync"
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
	productos       = []Producto{{"Monitor", 75000.00, 3}, {"Teclado", 20000.00, 5}, {"Mouse", 8000.00, 10}}
	servicios       = []Servicio{{"Armado de PC", 4500, 40}, {"Limpieza de PC", 3000, 25}}
	mantenimientos  = []Mantenimiento{{"Mantenimiento de PC", 3500}, {"Mantenimiento de auto", 6000}}
	totalCategorias = 0.0
)

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(3)
	go sumarProductos(productos, wg, mutex)
	go sumarServicios(servicios, wg, mutex)
	go sumarMantenimiento(mantenimientos, wg, mutex)

	wg.Wait()
	fmt.Println(totalCategorias)
}

func sumarProductos(productos []Producto, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	total := 0.0
	for _, prod := range productos {
		total += prod.Precio * float64(prod.Cantidad)
	}
	mutex.Lock()
	totalCategorias += total
	mutex.Unlock()
}

func sumarServicios(servicios []Servicio, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	total := 0.0
	for _, serv := range servicios {
		if serv.MinsTrabajados > 30 {
			total += serv.Precio * (float64(serv.MinsTrabajados) / 30)
		} else {
			total += serv.Precio
		}
	}
	mutex.Lock()
	totalCategorias += total
	mutex.Unlock()
}

func sumarMantenimiento(mantenimientos []Mantenimiento, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	total := 0.0
	for _, matn := range mantenimientos {
		total += matn.Precio
	}
	mutex.Lock()
	totalCategorias += total
	mutex.Unlock()
}
