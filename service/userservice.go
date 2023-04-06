package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"prj_01/models"
	"prj_01/utils"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUser
// @Summary 显示用户
// @Tags 首页
// @Success 200 {string} json {"code","message"}
// @Router /user [get]
func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询成功",
		"data":    data,
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
			"code":    -1,
			"message": "用户已存在",
			"data":    data,
		})
		return
	}
	// 补充

	if password != repassword {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
			"data":    data,
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    user,
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
		"code":    0,
		"message": "删除成功",
		"data":    user,
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
			"code":    -1,
			"message": "修改参数不匹配",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "修改成功",
			"data":    user,
		})
	}
}

// SelectUser
// @Summary 用户登录
// @Tags 首页
// @param name query string false "username"
// @param password query string false "password"
// @Success 200 {string} json {"code","message"}
// @Router /user/findandwd [post]
func FindUserByNameAndPwd(ctx *gin.Context) {
	data := models.UserBasic{}
	name := ctx.Query("name")
	password := ctx.Query("password")
	user_name := models.FindUserByName(name)
	if user_name.Name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}
	flag := utils.ValidPassword(password, user_name.Salt, user_name.Password)
	if !flag {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "登录失败",
			"data":    data,
		})
		return
	}
	user_pwd := utils.MakePassword(password, user_name.Salt)
	data = models.FindUserByNameAndPwd(name, user_pwd)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    data,
	})
}

// 防止跨越站点的伪请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(ctx *gin.Context) {
	ws, err := upGrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, ctx)
}

func MsgHandler(ws *websocket.Conn, ctx *gin.Context) {
	msg, err := utils.Subscribe(ctx, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
