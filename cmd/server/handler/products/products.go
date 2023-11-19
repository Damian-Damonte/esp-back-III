package products

import (
	"net/http"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
	"github.com/aldogayaladh/go-web-1598/internal/products"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service products.Service
}

func NewControllerProducts(service products.Service) *Controller {
	return &Controller{service: service}
}

// Doc ...
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		listProducts, err := c.service.GetAll(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": listProducts,
		})
	}
}

// Doc ...
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		producto, err := c.service.GetByID(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Doc ...
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productRequest domain.Producto

		err := ctx.Bind(&productRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err.Error(),
			})
			return
		}

		producto, err := c.service.Create(ctx, productRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})

	}
}

// Doc ...
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		var productRequest domain.Producto

		err := ctx.Bind(&productRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err.Error(),
			})
			return
		}

		producto, err := c.service.Update(ctx, productRequest, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Doc ...
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		err := c.service.Delete(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Producto eliminado",
		})
	}
}
