package models

import (
	"fmt"
	"time"

	"apiswayalan.com/rest/api/setup"
)

type UserModel struct{
	Id int16 `gorm:"primaryKey" json:"id_user"`
	Username string `gorm:"varchar(30)" json:"username"`
	NamaLengkap string `gorm:"varchar(30)" json:"nama_lengkap"`
	Tgl time.Time `json:"tgl"`
}

func (UserModel) TableName() string{
	return "tbl_user"
}


func GenerateTableUser(){
	dbuser:=setup.ConnectDB()
	if(!dbuser.Migrator().HasTable(&UserModel{})){
		dbuser.Migrator().CreateTable(&UserModel{})
		return 
	}
	fmt.Println("Table already exits")
}