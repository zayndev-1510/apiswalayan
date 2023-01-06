package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

var dbprodukmasuk *gorm.DB

type ProdukMasukModelRelasi struct {
	Id          int32  `json:"id"`
	IdBarang    int32  `json:"id_barang"`
	IdSuplier   int32  `json:"id_suplier"`
	NamaBarang  string `json:"nama_barang"`
	NamaSuplier string `json:"nama_suplier"`
	Jumlah      int32  `json:"jumlah"`
	Tgl         string `json:"tgl"`
	StokLama    int32  `json:"stok_lama"`
	Harga       int32  `json:"harga"`
}
type ProdukMasukModel struct {
	Id        int32  `json:"id"`
	IdBarang  int32  `json:"id_barang"`
	IdSuplier int32  `json:"id_suplier"`
	Jumlah    int32  `json:"jumlah"`
	Tgl       string `json:"tgl"`
	StokLama  int32  `json:"stok_lama"`
}

func (ProdukMasukModel) TableName() string {
	return "tbl_produk_masuk"
}

func init() {
	dbprodukmasuk = setup.ConnectDB()
}
func JumlahProdukMasuk()(int32){
	var stok int32
	row:=db.Table("tbl_produk_masuk").Select("sum(jumlah) as stok").Row()
	row.Scan(&stok)
	return stok

}

func GetDataProdukMasuk(Tgl string) ([]ProdukMasukModelRelasi, *gorm.DB) {
	var result []ProdukMasukModelRelasi
	if Tgl == "all" {
		queryall := dbprodukmasuk.Table("tbl_produk_masuk").Joins("JOIN tbl_barang ON tbl_barang.id=tbl_produk_masuk.id_barang").
			Joins("JOIN tbl_suplier ON tbl_suplier.id=tbl_produk_masuk.id_suplier").Select("tbl_barang.id as id_barang", "tbl_produk_masuk.id", "tbl_produk_masuk.id_suplier",
			"tbl_barang.nama_barang", "tbl_produk_masuk.jumlah", "tbl_produk_masuk.tgl", "tbl_produk_masuk.stok_lama",
			"tbl_suplier.nama_suplier", "tbl_barang.harga").Find(&result)
		return result, queryall
	} else {
		query := dbprodukmasuk.Table("tbl_produk_masuk").Joins("JOIN tbl_barang ON tbl_barang.id=tbl_produk_masuk.id_barang").
			Joins("JOIN tbl_suplier ON tbl_suplier.id=tbl_produk_masuk.id_suplier").Select("tbl_barang.id as id_barang", "tbl_produk_masuk.id", "tbl_produk_masuk.id_suplier",
			"tbl_barang.nama_barang", "tbl_produk_masuk.jumlah", "tbl_produk_masuk.tgl", "tbl_produk_masuk.stok_lama",
			"tbl_suplier.nama_suplier", "tbl_barang.harga").Where("tbl_produk_masuk.tgl=?", Tgl).Find(&result)
		return result, query

	}

}

func GetDataProdukMasukFilterMonth(a string, b string) ([]ProdukMasukModelRelasi, *gorm.DB) {
	var result []ProdukMasukModelRelasi
	query := dbprodukmasuk.Table("tbl_produk_masuk").Joins("JOIN tbl_barang ON tbl_barang.id=tbl_produk_masuk.id_barang").
		Joins("JOIN tbl_suplier ON tbl_suplier.id=tbl_produk_masuk.id_suplier").Select("tbl_barang.id as id_barang", "tbl_produk_masuk.id", "tbl_produk_masuk.id_suplier",
		"tbl_barang.nama_barang", "tbl_produk_masuk.jumlah", "tbl_produk_masuk.tgl", "tbl_produk_masuk.stok_lama",
		"tbl_suplier.nama_suplier", "tbl_barang.harga").Where("tbl_produk_masuk.tgl between ? and ?", a, b).Find(&result)
	return result, query
}

func GetDataProdukMasukFilterYear(a string, b string) ([]ProdukMasukModelRelasi, *gorm.DB) {
	var result []ProdukMasukModelRelasi
	query := dbprodukmasuk.Table("tbl_produk_masuk").Joins("JOIN tbl_barang ON tbl_barang.id=tbl_produk_masuk.id_barang").
		Joins("JOIN tbl_suplier ON tbl_suplier.id=tbl_produk_masuk.id_suplier").Select("tbl_barang.id as id_barang", "tbl_produk_masuk.id", "tbl_produk_masuk.id_suplier",
		"tbl_barang.nama_barang", "tbl_produk_masuk.jumlah", "tbl_produk_masuk.tgl", "tbl_produk_masuk.stok_lama",
		"tbl_suplier.nama_suplier", "tbl_barang.harga").Where("YEAR(tbl_produk_masuk.tgl) between ? and ?", a, b).Find(&result)
	return result, query
}

func SaveDataProdukMasuk(produk []ProdukMasukModel) *gorm.DB {

	query := dbprodukmasuk.Create(&produk)

	return query
}
