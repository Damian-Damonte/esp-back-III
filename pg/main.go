package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	variable1 := os.Getenv("VARIABLE_1")
	fmt.Println(variable1)
}