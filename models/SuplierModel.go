package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)


var connsuplier *gorm.DB
type TblSuplier struct{
	Id int32 `json:"id"`
	NamaSuplier string `json:"nama_suplier"`
	Alamat string `json:"alamat"`
	NomorHp string `json:"nomor_hp"`
	Email string `json:"email"`
	Tgl string `json:"tgl"`
}

func init(){
	connsuplier=setup.ConnectDB()
}

func GetDataSuplier()([] TblSuplier,*gorm.DB){

	var result [] TblSuplier

	query:=connsuplier.Table("tbl_suplier").Find(&result)
	return result,query
}

func SaveDataSuplier(row *TblSuplier) *gorm.DB{

	query:=connsuplier.Table("tbl_suplier").Create(&row)
	return query

}

func UpdateDataSuplier(id int32,row *TblSuplier) *gorm.DB{

	query:=connsuplier.Table("tbl_suplier").Where("id=?",id).Updates(&row)
	return query
}

func DeleteDataSuplier(row *TblSuplier) *gorm.DB{
	query:=connsuplier.Table("tbl_suplier").Delete(&row)
	return query
}