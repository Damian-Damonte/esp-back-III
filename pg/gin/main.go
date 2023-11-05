package main

import (
	"fmt"
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

	router.GET("/hello", helloWorld)

	personas := router.Group("/personas")
	personas.GET("", getPersonas)
	personas.GET(":nombre", getPersonaByNombre)
	personas.GET("/query", getPersonaQuerys)
	personas.POST("", postPersonas)

	router.Run(":8080")
}

func helloWorld(c *gin.Context) {
	c.Writer.WriteString("Hello world")
}

func getPersonas(c *gin.Context) {
	c.JSON(200, personas)
}

func getPersonaByNombre(c *gin.Context) {
	paramNombre := c.Param("nombre")
	var persona Persona
	fmt.Println(paramNombre)
	for _,v := range personas {
		if v.Nombre == paramNombre {
			persona = v
			break
		}
	}

	if persona.Nombre != "" {
		c.JSON(200, persona)
	} else {
		c.Status(404)
	}
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

func getPersonaQuerys(c *gin.Context) {
	queryNombre := c.Query("nombre")
	queryApellido := c.Query("apellido")
	var persona Persona

	for _,v := range personas {
		if v.Nombre == queryNombre && v.Apellido == queryApellido {
			persona = v
			break
		}
	}

	if persona.Nombre != "" {
		c.JSON(200, persona)
	} else {
		c.Status(404)
	}
}