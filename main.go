package main

import (
	"apiswayalan.com/rest/api/models"
	"apiswayalan.com/rest/api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Databarang models.BarangModels

func main() {
	c := gin.Default()

	c.Use(cors.Default())

	routers.RouterDashboard(c)
	routers.RouterBarang(c)

	routers.RouterRegistrasi(c)

	routers.RouterLogin(c)

	routers.RouterTransaksi(c)

	routers.RouterKasir(c)

	routers.RouterSuplier(c)

	routers.RouterProduKeluar(c)
	routers.RouterProdukMasuk(c)

	c.Run()
}
