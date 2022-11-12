package product

import (
	"net/http"
	"tokomania/internal/usecase/product"

	"github.com/gin-gonic/gin"
)

func GetProducts(useCase product.ProductUseCase) gin.HandlerFunc{
	return func(c *gin.Context){
		var page uint
		page = 1
		limit := 10

		result, err := useCase.GetProducts(c.Copy().Request.Context(), page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
} 
func GetProduct(useCase product.ProductUseCase) gin.HandlerFunc{
	return func(c *gin.Context){
		var id uint64

		result, err := useCase.GetProduct(c.Copy().Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
} 