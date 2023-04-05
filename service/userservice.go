package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"prj_01/models"
	"prj_01/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUser
// @Summary 显示用户
// @Tags 首页
// @Success 200 {string} json {"code","message"}
// @Router /user [get]
func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()
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
	salt := fmt.Sprintf("%06d", rand.Int31())
	// 用户名，电话号码和邮箱唯一
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "用户名已被注册",
		})
		return
	}
	// 补充

	if password != repassword {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
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
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateuser [post]
func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	user.ID = uint(id)
	user.Name = ctx.PostForm("name")
	user.Password = ctx.PostForm("password")
	user.Phone = ctx.PostForm("phone")
	user.Email = ctx.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "修改参数不匹配",
		})
	} else {
		models.UpdateUser(user)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
	}
}

// SelectUser
// @Summary 查询用户
// @Tags 首页
// @param name query string false "username"
// @param password query string false "password"
// @Success 200 {string} json {"code","message"}
// @Router /user/findandwd [post]
func FindUserByNameAndPwd(ctx *gin.Context) {
	name := ctx.Query("name")
	password := ctx.Query("password")
	user_name := models.FindUserByName(name)
	if user_name.Name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "该用户不存在",
		})
		return
	}
	flag := utils.ValidPassword(password, user_name.Salt, user_name.Password)
	if !flag {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "密码不正确",
		})
		return
	}
	user_pwd := utils.MakePassword(password, user_name.Salt)
	data := models.FindUserByNameAndPwd(name, user_pwd)
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
