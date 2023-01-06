package controllers

import (
	"time"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

var DataProdukKeluar=func(c *gin.Context){

	data,err:=models.DataProdukKeluaur()
	if err.Error !=nil{
		c.JSON(500,map[string]interface{}{"message":"Ada Kesalahan"+err.Error.Error(),"count":0})
		return
	}
	if err.RowsAffected==0{
		c.JSON(500,map[string]interface{}{"message":"Data Kosong","count":0,"data":data})
		return
	}

	var datanew []models.ProdukKeluar

	var totalproduk=0
	for _, f := range data {
		totalproduk=totalproduk+int(f.Jumlah)
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		datanew = append(datanew, models.ProdukKeluar{
			IdProduk:f.IdProduk , NamaBarang: f.NamaBarang,
			Tgl: t.Format("2006-01-02"), Jumlah: f.Jumlah,Harga: f.Harga})
	}

	c.JSON(200,map[string]interface{}{"message":"Data Found","count":1,"data":datanew,"totalproduk":totalproduk})

}