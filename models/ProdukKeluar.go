package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

var connprodukkeluar *gorm.DB
type ProdukKeluar struct{
	IdProduk int32 `json:"id_produk"`
	NamaBarang 		string			`json:"nama_barang"`
	Jumlah  		float32         `json:"jumlah"`
	Harga 		float32         `json:"harga"`
	
	Tgl string `json:"tgl"`
}

func init(){
	connprodukkeluar=setup.ConnectDB()
}
func JumlahProdukKeluar()int32{
	var jumlah int32
	row:=db.Table("tbl_keranjang").Select("sum(jumlah) as stok").Row()
	row.Scan(&jumlah)
	return jumlah

}
func DataProdukKeluaur() ([] ProdukKeluar,*gorm.DB){
	var result [] ProdukKeluar
	query:=connprodukkeluar.Table("tbl_keranjang").
	Joins("JOIN tbl_barang ON tbl_keranjang.id_barang=tbl_barang.id").
	Select("tbl_keranjang.id_barang as id_produk",
	"tbl_barang.nama_barang","tbl_barang.id as id_produk",
	"tbl_keranjang.jumlah","tbl_keranjang.harga","tbl_barang.stok","tbl_keranjang.tgl").
	Find(&result)
	return result,query
}