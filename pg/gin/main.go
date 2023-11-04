package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int8   `json:"edad"`
}

var personas = []Persona{}

func main() {
	router := gin.Default()

	router.GET("/personas", getPersonas)
	router.POST("/personas", postPersonas)

	router.Run(":8080")
}

func getPersonas(c *gin.Context) {
	c.JSON(200, personas)
}

func postPersonas(c *gin.Context) {
	var persona Persona
	err := c.Bind(&persona)
	
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	personas = append(personas, persona)
	c.JSON(201, persona)
}

func postPersonas2(c *gin.Context) {
	var persona Persona
	requestBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(requestBody, &persona)

	personas = append(personas, persona)
	c.JSON(201, persona)
}





