package controllers

import (
	"net/http"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)
var LoginAkun=func(c *gin.Context){
	
	var username=c.Request.PostFormValue("username")
	var sandi=c.Request.PostFormValue("sandi")
	var data=models.TblLogin{Username: username,Sandi:sandi}
	result:=models.Login(&data)
	
	c.IndentedJSON(http.StatusOK,result)
}

