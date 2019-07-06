package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/form/parameterforms"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductOutController struct{}

var productOutObject = new(models.Model)
var productOutParam parameterforms.Pagination

func (w *ProductOutController) GetBarangKeluar() gin.HandlerFunc {
	return func(c *gin.Context) {
		productOutParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productOutParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productOutParam.Offset *= productOutParam.Limit
		r := productOutObject.GetBarangKeluar(productOutParam.Limit, productOutParam.Offset)
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
