package main

import (
	"fmt"
	"time"
)

func main() {
	middleware1(middleware2(sumar))(5,10)
}

func sumar(a, b int) int {
	fmt.Println("Suma:", a+b)
	return a + b
}

func middleware1(f func(int, int) int) func(int, int) int {
	return func(a,b int) int {
		fmt.Println("1 antes")
		resultado := f(a,b)
		fmt.Println("1 despues")
		return resultado
	}
}

func middleware2(f func(int, int) int) func(int, int) int {
	return func(a,b int) int {
		fmt.Println("2 antes")
		resultado := f(a,b)
		fmt.Println("2 despues")
		return resultado
	}
}

func mostrarTiempo(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {

		fmt.Println(time.Now())

		return f(a,b)
	}
}

func mostarParametros(f func(int, int) int) func(int, int) int {
	return func(a,b int) int {
		fmt.Printf("primer parametro: %d - segundo parametro: %d \n",a,b)
		return f(a,b)
	}
}

/*
	fmt.Println("funcion1 antes")
		fmt.Println("funcion2 antes")
			resultado := f(a,b)
		fmt.Println("funcion2 despues")
	fmt.Println("funcion1 despues")
*/