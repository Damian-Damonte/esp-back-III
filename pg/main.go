package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	Code_value   string  `json:"code_value" binding:"required"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

var productos []Producto

func main() {
	cargarProductos()
	engine := gin.Default()

	productos := engine.Group("/productos")
	productos.GET("", getProductos)
	productos.POST("", saveProducto)

	engine.Run(":8080")
}

func getProductos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, productos)
}

func saveProducto(ctx *gin.Context) {
	var producto Producto
	if err := ctx.ShouldBindJSON(&producto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	for _,p := range productos {
		if p.Code_value == producto.Code_value {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ya existe un producto con el code value %s", p.Code_value)})
			return
		}
	}

	_, err := time.Parse("02/01/2006", producto.Expiration)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "formato de fecha invalido"})
		return
	}

	id := len(productos) +1
	producto.Id = id

	productos = append(productos, producto)

	ctx.JSON(http.StatusCreated, producto)
}

func cargarProductos() {
	producto1 := Producto{1, "Coca Cola", 8, "C45FSF", true, "16/01/2021", 250}
	producto2 := Producto{2, "Papas fritas", 3, "DF253D5", true, "19/11/2022", 750}
	producto3 := Producto{3, "Galletitas", 15, "GE352D", false, "02/05/2022", 500}
	productos = append(productos, producto1, producto2, producto3)
}