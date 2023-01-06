package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

// Deklarasi Variabel Database

// Installasi Struktru Tabel
var conntransaksi *gorm.DB

type TblTranasksi struct {
	Id              int32           `json:"id"`
	IdPengguna      int32           `json:"id_pengguna"`
	NomorTransaksi  string          `json:"nomor_transaksi"`
	NomorKeranjang  string          `json:"nomor_keranjang"`
	Total           float32         `json:"total"`
	Jumlah          float32         `json:"jumlah"`
	Kasir           string          `json:"kasir"`
	Status          int32           `json:"status"`
	Tgl             string          `gorm:"type:date" json:"tgl"`
	RegistrasiModel RegistrasiModel `gorm:"Foreignkey:IdPengguna;association_foreignkey:Id;" json:"pembeli"`
}

func (TblTranasksi) TableName() string {
	return "tbl_transaksi"
}

func init() {
	conntransaksi = setup.ConnectDB()
}

func GetDataTransaksiAll() ([]TblTranasksi, *gorm.DB) {
	var result []TblTranasksi
	query := conntransaksi.Preload("RegistrasiModel").Find(&result)
	return result, query
}

func GetDataTransaksiMonth(start string,end string) ([]TblTranasksi, *gorm.DB) {
	var result []TblTranasksi
	query := conntransaksi.Preload("RegistrasiModel").Where("tgl between ? and ?",start,end).Find(&result)
	return result, query
}

func GetDataTransaksiYear(start string,end string) ([]TblTranasksi, *gorm.DB) {
	var result []TblTranasksi
	query := conntransaksi.Preload("RegistrasiModel").Where("YEAR(tgl) between ? and ?",start,end).Find(&result)
	return result, query
}
func GetDataTransaksi(row *TblTranasksi) []TblTranasksi {
	var result []TblTranasksi
	conntransaksi.Preload("RegistrasiModel").Where(&row).Find(&result)
	return result
}

func Getrumus(a int32, b int32) (persegi int32, segitiga int32) {
	persegi = a * b
	segitiga = a + b
	return persegi, segitiga
}

func ProdukTransaksi(row *TblKeranjang) []TblKeranjang {
	var result []TblKeranjang
	conntransaksi.Table("tbl_keranjang").
		Joins("JOIN tbl_barang ON tbl_keranjang.id_barang=tbl_barang.id").
		Select("tbl_keranjang.id as id_keranjang", "tbl_keranjang.nomor_keranjang", "tbl_keranjang.id_barang",
			"tbl_barang.nama_barang", "tbl_barang.id as id_produk",
			"tbl_keranjang.jumlah", "tbl_keranjang.harga", "tbl_keranjang.status", "tbl_barang.stok").
		Where(&row).
		Find(&result)
	return result

}

func PerbaruiTransaksi(where TblTranasksi, row map[string]interface{}) *gorm.DB {
	return conntransaksi.Model(TblTranasksi{}).Where(&where).Updates(&row)
}
