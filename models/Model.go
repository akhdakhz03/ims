package models

import (
	"api-inventory/utils/db"
	"api-inventory/utils/form/tableform"
)

type Model struct{}

func (m Model) GetBarangMasuk(limit int, offset int) (list []tableform.ProductIn) {
	db.GetDB().Table("product_in").Limit(limit).Offset(offset).Find(&list)
	return list
}

func (m Model) GetBarangKeluar(limit int, offset int) (list []tableform.ProductOut) {
	db.GetDB().Table("product_out").Limit(limit).Offset(offset).Find(&list)
	return list
}

func (m Model) GetStockBarang(limit int, offset int) (list []tableform.Products) {
	db.GetDB().Table("products").Limit(limit).Offset(offset).Find(&list)
	return list
}
