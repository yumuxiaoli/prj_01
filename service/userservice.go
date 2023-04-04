package service

import (
	"net/http"
	"prj_01/models"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})

}
