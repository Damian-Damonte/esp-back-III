package main

import (
	"fmt"
	"sync"
)

/*
Al ejecutar podríamos esperar que el contador sea 0 ya que sumar() suma un millón y restar() resta un millón.
Si ejecutamos varias veces podremos notar que en algunos casos el resultado no es 0. Eso se debe a que se está produciendo una condición de carrera entre las 2 goroutines.
Básicamente se producen situaciones en las que a ambas goroutines modifican o leen la variable contador al mismo tiempo.
*/

func main() {
	contador := 0
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go sumar(&contador, wg)
	go restar(&contador, wg)
	wg.Wait()

	fmt.Printf("Contador: %d", contador)
}

// Suma 1 al contador un millon de veces
func sumar(contador *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=0; i < 1000000; i++ {
		*contador += 1
	}
}

// Resta 1 al contador un millon de veces
func restar(contador *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=0; i < 1000000; i++ {
		*contador -= 1
	}
}

/*
============================ USANDO MUTEX ============================
En este caso no importa cuantas veces ejecutemos, contador siempre dará 0.
En este caso nos aseguramos de que las goroutines no modifiquen o lean la variable contador a la vez utilizando mutex.

Comentar el código de arriba y descomentar el de abajo
*/

// func main() {
// 	contador := 0
// 	wg := &sync.WaitGroup{}
// 	mutex := &sync.Mutex{}

// 	wg.Add(2)
// 	go sumar(&contador, wg, mutex)
// 	go restar(&contador, wg, mutex)
// 	wg.Wait()

// 	fmt.Printf("Contador: %d\n", contador)
// }

// func sumar(contador *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	for i:=0; i < 1000000; i++ {
// 		mutex.Lock()
// 		*contador += 1
// 		mutex.Unlock()
// 	}
// }

// func restar(contador *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	for i:=0; i < 1000000; i++ {
// 		mutex.Lock()
// 		*contador -= 1
// 		mutex.Unlock()
// 	}
// }