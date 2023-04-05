package service

import (
	"net/http"
	"prj_01/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser
// @Summary 显示用户
// @Tags 首页
// @Success 200 {string} json {"code","message"}
// @Router /user [get]
func GetUserList(ctx *gin.Context) {
	data := make([]*models.UserBasic, 1)
	data = models.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// CreatUser
// @Summary 新增用户
// @Tags 首页
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/creatuser [get]
func CreatUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Query("name")
	password := ctx.Query("password")
	repassword := ctx.Query("repassword")
	if password != repassword {
		ctx.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 首页
// @param id query string false "序号"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteuser [get]
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 首页
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateuser [post]
func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	user.ID = uint(id)
	user.Name = ctx.PostForm("name")
	user.Password = ctx.PostForm("password")
	models.UpdateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}
