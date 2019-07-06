package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/form/parameterforms"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductStockController struct{}

var productStockObject = new(models.Model)
var productStockParam parameterforms.Pagination

func (w *ProductStockController) GetStockBarang() gin.HandlerFunc {
	return func(c *gin.Context) {
		productStockParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productStockParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productStockParam.Offset *= productStockParam.Limit
		r := productStockObject.GetStockBarang(productStockParam.Limit, productStockParam.Offset)
		if len(r) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Data Not Found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"data":    r,
			})

		}
	}
}
