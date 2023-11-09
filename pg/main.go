package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Empleado struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Activo bool   `json:"activo"`
}

var Empleados []Empleado

func main() {
	Empleados = generarEmpleados()

	router := gin.Default()

	router.GET("", bienvenida)
	groupEmployees := router.Group("/employees")
	groupEmployees.GET("", getAllEmployees)
	groupEmployees.GET(":id", getEmployeeById)
	router.POST("/employeeparams", postEmployeeParams)
	router.GET("/employeesactives", getEmployeesActives)

	router.Run(":8080")
}

func generarEmpleados() []Empleado {
	return []Empleado{
		{1, "Damian", true},
		{2, "Juan", true},
		{3, "Matias", false},
		{4, "Marcelo", true},
		{5, "Santiago", true},
	}
}

func bienvenida(c *gin.Context) {
	c.Writer.WriteString("¡Bienvenido a la empresa Gophers!")
}

func getAllEmployees(c *gin.Context) {
	c.JSON(200, Empleados)
}

func getEmployeeById(c *gin.Context) {
	var employee Empleado
	pathVariableId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
		c.Writer.WriteString("El id debe ser numérico")
	}

	for _, v := range Empleados {
		if v.Id == pathVariableId {
			employee = v
			break
		}
	}

	if employee.Id != 0 {
		c.JSON(200, employee)
	} else {
		c.Status(404)
	}
}

func postEmployeeParams(c *gin.Context) {
	paramId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.Status(400)
		c.Writer.WriteString("Param id debe ser númerico")
		return
	}
	paramNombre := c.Query("nombre")
	paramActivo, err := strconv.ParseBool(c.Query("activo"))
	if err != nil {
		c.Status(400)
		c.Writer.WriteString("Param activo debe ser true o false")
		return
	}

	empleado := Empleado{paramId, paramNombre, paramActivo}
	Empleados = append(Empleados, empleado)
	c.JSON(201, empleado)
}

func getEmployeesActives(c *gin.Context) {
	active := true
	if c.Query("active") == "false" {
		active = false
	}
	var empleadosActive []Empleado

	for _, v := range Empleados {
		if v.Activo == active {
			empleadosActive = append(empleadosActive, v)
		}
	}

	c.JSON(200, empleadosActive)
}
