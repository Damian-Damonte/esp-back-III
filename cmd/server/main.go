package main

import (
	"log"

	handlerPing "github.com/aldogayaladh/go-web-1598/cmd/server/handler/ping"
	handlerProducto "github.com/aldogayaladh/go-web-1598/cmd/server/handler/products"
	"github.com/aldogayaladh/go-web-1598/internal/products"
	"github.com/aldogayaladh/go-web-1598/pkg/jsonstorage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	jsonStorage := jsonstorage.NewJsonStorage()
	jsonStorage.Inicializacion()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	repostory := products.NewMemoryRepository(jsonStorage)
	service := products.NewServiceProduct(repostory)
	controllerProduct := handlerProducto.NewControllerProducts(service)

	engine := gin.Default()

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		grupoProducto := group.Group("/productos")
		{
			grupoProducto.GET("", controllerProduct.HandlerGetAll())
			grupoProducto.GET("/:id", controllerProduct.HandlerGetByID())
			grupoProducto.POST("", controllerProduct.HandlerCreate())
			grupoProducto.PUT("/:id", controllerProduct.HandlerUpdate())
			grupoProducto.DELETE("/:id", controllerProduct.HandlerDelete())
		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
