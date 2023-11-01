package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var input string
	channelDevoluciones := make(chan string)
	channelPrestamos := make(chan string)

	go Devolucion(channelDevoluciones)

	go Prestamo(channelPrestamos)

	go PrintDevolucion(channelDevoluciones)

	go PrintPrestamo(channelPrestamos)

	
	fmt.Scan(&input)
	if input != "" {
		fmt.Println("Saliendo del programa...")
		os.Exit(0)
	}
}

func Devolucion(channelDevolucion chan string) {
	defer close(channelDevolucion)

	for {
		time.Sleep(time.Second * 1)
		channelDevolucion <- "Devolucion procesada"
	}
}

func Prestamo(channelPretamo chan string) {
	defer close(channelPretamo)
	for {
		time.Sleep(time.Second / 2)
		channelPretamo <- "Prestamo procesado"
	}
}

func PrintDevolucion(channelDevolucion chan string) {
	for devolucion := range channelDevolucion {
		fmt.Println(devolucion)
	}
}

func PrintPrestamo(channelPrestamo chan string) {
	for prestamo := range channelPrestamo {
		fmt.Println(prestamo)
	}
}
