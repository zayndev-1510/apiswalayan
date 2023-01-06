package controllers

import (
	
	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

type SumDashboard struct{
	StokProduk int32 `json:"stokproduk"`
	ProdukMasuk int32 `json:"produkmasuk"`
	ProdukKeluar int32 `json:"produkkeluar"`
	TotalTransaksiKeluar float32 `json:"totaltransaksikeluar"`
	Kasir int32 `json:"kasir"`
}

func Dashboard(c *gin.Context){
	var sum SumDashboard
	sum.StokProduk=models.JumlahProduk()
	sum.ProdukMasuk=models.JumlahProdukMasuk()
	sum.ProdukKeluar=models.JumlahProdukKeluar()
	sum.TotalTransaksiKeluar=models.JumlahTransaksiKeluar()
	sum.Kasir=models.JumlahKasir()
	data,db:=models.GetGrafikProdukTahuna()

	if db.Error !=nil{
		c.JSON(500,gin.H{"message":"Database Error"})
		return
	}

	produkmasuk,err:=models.GetGrafikProdukMasukTahunan()
	if err.Error !=nil{
		c.JSON(500,gin.H{"message":"database errr"+err.ToSQL(err.Callback().Query().Execute)})
		return
	}


	c.JSON(200,gin.H{"message":"Data Ready","data":sum,"grafikprodukkeluar":data,"grafikprodukmasuk":produkmasuk})
}