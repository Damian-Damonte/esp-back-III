package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DetalleCompra struct {
	NombreProducto string  `json:"nombre_producto"`
	Cantidad       int     `json:"cantidad"`
	PrecioTotal    float64 `json:"precio_total"`
}

type Producto struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type Storage struct {
	Productos []Producto
}

func (s *Storage) completarProductos() {
	producto1 := Producto{754, "Coca Cola", 8, "C45FSF", true, "16/01/2021", 250}
	s.Productos = append(s.Productos, producto1)
	producto2 := Producto{253, "Papas fritas", 3, "DF253D5", true, "19/11/2022", 750}
	s.Productos = append(s.Productos, producto2)
	producto3 := Producto{126, "Galletitas", 15, "GE352D", false, "02/05/2022", 500}
	s.Productos = append(s.Productos, producto3)
}

func (s *Storage) agregarProducto(p Producto) {
	s.Productos = append(s.Productos, p)
}

var storage = Storage{}

func main() {
	storage.completarProductos()
	engine := gin.Default()

	apiBase := engine.Group("/api/v1")
	{
		apiBase.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"mensaje": "pong"}) })
		productos := apiBase.Group("/productos")
		{
			productos.GET("/:id", getProductById)
			productos.GET("productparams", productParams)
			productos.GET("/searchbyquantity", searchProductByQuantity)
			productos.GET("/buy", buyProduct)
		}
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func getProductById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	idCast, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro id incorrecto"})
		return
	}

	var product Producto

	for _, p := range storage.Productos {
		if p.Id == idCast {
			product = p
		}
	}

	if product.Id == 0 {
		errMsj := fmt.Sprintf("producto con id %d no encontrado", idCast)
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": errMsj})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func productParams(ctx *gin.Context) {
	queryId := ctx.Query("id")
	idCast, err := strconv.Atoi(queryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro id incorrecto"})
		return
	}
	queryName := ctx.Query("name")
	queryQuantity := ctx.Query("quantity")
	quantityCast, err := strconv.Atoi(queryQuantity)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro quantity incorrecto"})
		return
	}
	queryCodeValue := ctx.Query("code_value")
	queryIsPublished := ctx.Query("is_published")
	isPublishedCast, err := strconv.ParseBool(queryIsPublished)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro is_published incorrecto"})
		return
	}
	queryExpiration := ctx.Query("expiration")
	queryPrice := ctx.Query("price")
	priceCast, err := strconv.ParseFloat(queryPrice, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro price incorrecto"})
		return
	}

	producto := Producto{
		Id:           idCast,
		Name:         queryName,
		Quantity:     quantityCast,
		Code_value:   queryCodeValue,
		Is_published: isPublishedCast,
		Expiration:   queryExpiration,
		Price:        priceCast,
	}

	storage.agregarProducto(producto)

	ctx.JSON(http.StatusOK, producto)
}

func searchProductByQuantity(ctx *gin.Context) {
	queryMin := ctx.Query("min")
	minCast, err := strconv.Atoi(queryMin)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro min incorrecto"})
		return
	}
	queryMax := ctx.Query("max")
	maxCast, err := strconv.Atoi(queryMax)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro max incorrecto"})
		return
	}

	listaProductos := []Producto{}

	for _, p := range storage.Productos {
		if p.Quantity > minCast && p.Quantity < maxCast {
			listaProductos = append(listaProductos, p)
		}
	}

	ctx.JSON(http.StatusOK, listaProductos)
}

func buyProduct(ctx *gin.Context) {
	queryCode := ctx.Query("code")
	queryCantidad := ctx.Query("cantidad")
	cantidadCast, err := strconv.Atoi(queryCantidad)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mensaje": "parametro cantidad incorrecto"})
		return
	}

	var productoByCode Producto

	for _,v := range storage.Productos {
		if v.Code_value == queryCode {
			productoByCode = v
			break
		}
	}

	if productoByCode.Id == 0 {
		errMsj := fmt.Sprintf("producto con codigo '%s' no encontrado", queryCode)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"mensaje": errMsj})
		return
	}

	detalle := DetalleCompra {
		NombreProducto: productoByCode.Name,
		Cantidad: cantidadCast,
		PrecioTotal: productoByCode.Price * float64(cantidadCast),
	}
	
	ctx.JSON(http.StatusOK, detalle)
}