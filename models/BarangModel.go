package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

var db *gorm.DB

type BarangModels struct {
	Id           int32   `json:"id"`
	NomorBarcode string  `json:"nomor_barcode"`
	NamaBarang   string  `json:"nama_barang"`
	Harga        float32 `json:"harga"`
	Stok         float32 `json:"stok"`
	HargaJual    float32 `json:"harga_jual"`
	Tgl string `json:"tgl"`
}


func (BarangModels) TableName() string {
	return "tbl_barang"
}
func init() {
	db = setup.ConnectDB()
}


func JumlahProduk()(int32){
	var stok int32
	row:=db.Table("tbl_barang").Select("sum(stok) as stok").Row()
	row.Scan(&stok)
	return stok

}
func CreatedProduk(row *BarangModels) *gorm.DB {
	return db.Create(&row)
}

func UpdateProduk(row *BarangModels, Id int32) *gorm.DB {
	return db.Where("id= ?", Id).Updates(row)
}
func DeleteProduk(row *BarangModels) *gorm.DB {
	return db.Delete(&row)
}

func GetAllData() []BarangModels {
	var Barang []BarangModels
	db.Find(&Barang)
	return Barang
}

func GetDataById(Id int64) (*BarangModels, *gorm.DB) {
	var getBarang BarangModels
	db := db.Where("id = ?", Id).Find(&getBarang)
	return &getBarang, db
}

func UpdateStokBarang(where BarangModels,row BarangModels) *gorm.DB{

	query:=db.Where(where).Updates(row)
	return query

}

func UpdateStokBaragArray(Id int32,row BarangModels) *gorm.DB{
	result:=db.Model(&row).Where("id=?",Id).Update("stok",row.Stok)
	return result
}
