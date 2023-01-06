package routers

import (
	"apiswayalan.com/rest/api/controllers"
	"github.com/gin-gonic/gin"
)

var RouterBarang = func(router *gin.Engine) {
	group := router.Group("produk")
	group.GET("loadDataProduk", controllers.GetDataBarang)
	group.GET("detailProduk", controllers.DetailBarang)
	group.POST("createProduk", controllers.CreateBarang)
	group.POST("updateProduk", controllers.UpdateBarang)
	group.GET("deleteProduk/:id", controllers.DeleteBarang)
}

var RouterRegistrasi = func(router *gin.Engine) {
	router.POST("daftarakun", controllers.RegistrasiAkun)
}

var RouterLogin = func(router *gin.Engine) {
	router.POST("loginakun", controllers.LoginAkun)
}

var RouterTransaksi = func(router *gin.Engine) {
	group := router.Group("transaksi")
	group.GET("loadDataTransaksi", controllers.GetDataTransaksi)
	group.GET("loadDataTransaksiAll", controllers.GetDataTransaksiAll)
	group.POST("loadDataTransaksiMonth", controllers.GetDataTransaksiMonth)
	group.POST("loadDataTransaksiYear", controllers.GetDataTransaksiYear)

	group.POST("DataProdukTransaksi", controllers.GetDataProdukTransaksi)
	group.POST("PerbaruiTransaksi", controllers.PerbaruiTransaksi)
}

var RouterKasir = func(router *gin.Engine) {
	group := router.Group("kasir")
	group.GET("DataKasir", controllers.Datakasir)
	group.POST("CreateKasir", controllers.CreateKasir)
	group.POST("UpdateKasir", controllers.UpdateKasir)
	group.POST("DeleteKasir", controllers.DeleteKasir)
	group.POST("coba", controllers.Coba)
}

var RouterSuplier = func(router *gin.Engine) {
	grup := router.Group("suplier")
	grup.GET("DataSuplier", controllers.GetDataSuplier)
	grup.POST("SaveSuplier", controllers.AddDataSuplier)
	grup.POST("UpdateSuplier", controllers.UpdateSuplier)
	grup.GET("DeleteSuplier/:id", controllers.DeleteSuplier)
}

var RouterProduKeluar = func(router *gin.Engine) {
	grup := router.Group("produk")
	grup.GET("keluar", controllers.DataProdukKeluar)
}

var RouterProdukMasuk = func(router *gin.Engine) {
	grup := router.Group("produk")
	grup.GET("masuk/:tgl", controllers.GetDataProdukMasuk)
	grup.POST("masuk/filter", controllers.GetDataProdukMasukFilter)
	grup.POST("masuk/save", controllers.SaveDataProdukMasuk)
}

func RouterDashboard(router *gin.Engine){
	grup:=router.Group("dashboard")
	grup.GET("main", controllers.Dashboard)
}
