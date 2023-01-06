package controllers

import (
	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

type SumDashboard struct{
	StokProduk int32 `json:"stokproduk"`
	ProdukMasuk int32 `json:"produkmasuk"`
	ProdukKeluar int32 `json:"produkekluar"`
	TotalTransaksiKeluar float32 `json:"totaltransaksikeluar"`
}

func Dashboard(c *gin.Context){
	var sum SumDashboard
	sum.StokProduk=models.JumlahProduk()
	sum.ProdukMasuk=models.JumlahProdukMasuk()
	sum.ProdukKeluar,sum.TotalTransaksiKeluar=models.JumlahProdukKeluar()
	c.JSON(200,gin.H{"message":"Data Ready","data":sum})
}