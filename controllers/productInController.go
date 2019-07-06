package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/form/parameterforms"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductInController struct{}

var productInObject = new(models.Model)
var productInParam parameterforms.Pagination

func (w *ProductInController) GetBarangMasuk() gin.HandlerFunc {
	return func(c *gin.Context) {
		productInParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productInParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productInParam.Offset *= productInParam.Limit
		r := productInObject.GetBarangMasuk(productInParam.Limit, productInParam.Offset)
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
