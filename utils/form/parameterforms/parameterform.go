package parameterforms

import "time"

//define parameter variable

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Products struct {
	SKU          string `json:"SKU"`
	ItemName     string `json:"ItemName"`
	CurrentStock int    `json:"CurrentStock"`
}

type ProductIn struct {
	Id           int       `json:"Id"`
	Time         time.Time `json:"Time"`
	SKU          string    `json:"SKU"`
	ProductName  string    `json:"ProductName"`
	TotalStock   int       `json:"TotalStock"`
	ActualStock  int       `json:"ActualStock"`
	PricePerItem float64   `json:"PricePerItem"`
	TotalPrice   float64   `json:"TotalPrice"`
	Kwitansi     string    `json:"Kwitansi"`
	Remark       string    `json:"Remark"`
}

type ProductOut struct {
	Id          int       `json:"Id"`
	Time        time.Time `json:"Time"`
	ProductName string    `json:"ProductName"`
	Qty         int       `json:"Qty"`
	SellPrice   float64   `json:"SellPrice"`
	TotalPrice  float64   `json:"TotalPrice"`
	Remark      string    `json:"Remark"`
	SKU         string    `json:"SKU"`
}

type DateRange struct {
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}
