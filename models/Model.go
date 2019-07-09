package models

import (
	"api-inventory/utils/db"
	"api-inventory/utils/form/tableform"

	"github.com/jinzhu/gorm"
)

type Model struct{}

func (m Model) InsertProduct(data tableform.Products) tableform.Products {
	db.GetDB().Table("product_in").Create(&data)
	return data
}

func (m Model) GetBarangMasuk(limit int, offset int) (list []tableform.ProductIn) {
	db.GetDB().Table("product_in").Limit(limit).Offset(offset).Find(&list)
	return list
}

func (m Model) GetAllBarangMasuk() (list []tableform.ProductIn) {
	db.GetDB().Table("product_in").Find(&list)
	return list
}

func (m Model) GetBarangKeluar(limit int, offset int) (list []tableform.ProductOut) {
	db.GetDB().Table("product_out").Limit(limit).Offset(offset).Find(&list)
	return list
}

func (m Model) GetAllBarangKeluar() (list []tableform.ProductOut) {
	db.GetDB().Table("product_out").Find(&list)
	return list
}

func (m Model) GetStockBarang(limit int, offset int) (list []tableform.Products) {
	db.GetDB().Table("products").Limit(limit).Offset(offset).Find(&list)
	return list
}

func (m Model) InsertBarangMasuk(data tableform.ProductIn) tableform.ProductIn {
	db.GetDB().Table("product_in").Create(&data)
	return data
}

func (m Model) UpdateBarangMasuk(data tableform.ProductIn) (res tableform.ProductIn) {
	var total float64 = data.PricePerItem * float64(data.TotalStock)
	db.GetDB().Table("product_in").Where("id = ?", data.Id).First(&res).Updates(map[string]interface{}{"TotalStock": data.TotalStock, "ActualStock": data.ActualStock, "PricePerItem": data.PricePerItem, "Time": data.Time, "TotalPrice": total, "Kwitansi": data.Kwitansi, "Remark": data.Remark})
	return res
}

func (m Model) GetBarangMasukById(id int) (res tableform.ProductIn) {
	db.GetDB().Table("product_in").Where("id = ?", id).First(&res)
	return res
}

func (m Model) DeleteBarangMasuk(id int) bool {
	var res tableform.ProductIn
	db.GetDB().Table("product_in").Where("id = ?", id).Delete(&res)
	return true
}

func (m Model) InsertBarangKeluar(data tableform.ProductOut) tableform.ProductOut {
	db.GetDB().Table("product_out").Create(&data)
	return data
}

func (m Model) UpdateBarangKeluar(data tableform.ProductOut) (res tableform.ProductOut) {
	db.GetDB().Table("product_out").Where("id = ?", data.Id).First(&res).Updates(map[string]interface{}{"SKU": data.SKU, "SellPrice": data.SellPrice, "TotalPrice": (float64(data.Qty) * data.SellPrice), "Qty": data.Qty, "Time": data.Time, "Remark": data.Remark})
	return res
}

func (m Model) GetBarangKeluarById(id int) (res tableform.ProductIn) {
	db.GetDB().Table("product_out").Where("id = ?", id).First(&res)
	return res
}

func (m Model) DeleteBarangKeluar(id int) bool {
	var res tableform.ProductOut
	db.GetDB().Table("product_out").Where("id = ?", id).Delete(&res)
	return true
}

func (m Model) UpdateStock(sku string, stock int) (list tableform.ProductIn) {
	db.GetDB().Table("products").Where("SKU = ?", sku).First(&list).Update("CurrentStock", stock)
	return list
}

func (m Model) DeleteProduct(id int) bool {
	db.GetDB().Raw("Delete from products where id = ?", id)
	return true
}

func (m Model) GetQtyBarang(sku string) (list tableform.Products) {
	db.GetDB().Table("products").Where("SKU = ?", sku).First(&list)
	return list
}

func (m Model) GetLaporanNilaiBarangDetail() (result []tableform.LaporanNilaiBarang) {
	db.GetDB().Raw("select a.SKU as SKU, a.ItemName as ItemName, a.CurrentStock as CurrentStock, round(sum(pi.TotalPrice)/sum(pi.ActualStock),0) as avgprice, (round(sum(pi.TotalPrice)/sum(pi.ActualStock),0) * a.CurrentStock) as total  from products a left join product_in pi on a.SKU = pi.SKU group by a.SKU, a.ItemName, a.CurrentStock").Scan(&result)
	return result
}

func (m Model) GetAllLaporanPenjualan() (result []tableform.LaporanPenjualan) {
	db.GetDB().Raw("select REPLACE(Remark,'Pesanan ', '') as IdPesanan, po.SKU, Time as Waktu, ProductName as NamaBarang, Qty as Jumlah, (SellPrice/1) as HargaJual, TotalPrice as Total, hb.HargaBeli, SellPrice - (Qty*hb.HargaBeli) as Laba from product_out po left join (select SKU, round(sum(TotalPrice)/sum(ActualStock),0) as HargaBeli from product_in group by SKU) hb on po.SKU = hb.SKU WHERE Remark NOT IN ('Barang Hilang', 'Barang Rusak')").Scan(&result)
	return result
}

func (m Model) GetLaporanPenjualan(date_start string, date_end string) (result []tableform.LaporanPenjualan) {
	db.GetDB().Raw("select REPLACE(Remark,'Pesanan ', '') as IdPesanan, po.SKU, Time as Waktu, ProductName as NamaBarang, Qty as Jumlah, (SellPrice/1) as HargaJual, TotalPrice as Total, hb.HargaBeli, SellPrice - (Qty*hb.HargaBeli) as Laba from product_out po left join (select SKU, round(sum(TotalPrice)/sum(ActualStock),0) as HargaBeli from product_in group by SKU) hb on po.SKU = hb.SKU WHERE Remark NOT IN ('Barang Hilang', 'Barang Rusak') and date(Time) between ? and ?", date_start, date_end).Scan(&result)
	return result
}

func (m Model) GetTotalPenjualan(date_start string, date_end string) (count int64) {
	db.GetDB().Raw("select count(distinct Remark) from product_out WHERE Remark NOT IN ('Barang Hilang', 'Barang Rusak') and date(Time) between ? and ?", date_start, date_end).Count(&count)
	return count
}

func (m Model) GetNamaBarang(sku string) (result tableform.Products) {
	db.GetDB().Table("products").Where("sku =?", sku).Select("ItemName").Find(&result)
	return result
}

func (m Model) GetJumlahSKU() (count *gorm.DB) {
	db.GetDB().Table("products").Count(&count)
	return count
}

func (m Model) InsertHistory(data tableform.ProductHistory) tableform.ProductHistory {
	db.GetDB().Table("product_history").Create(&data)
	return data
}
