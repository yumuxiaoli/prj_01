package models

import (
	"fmt"
	"prj_01/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `vaild:"email"`
	Identity      string // 唯一标识
	ClientIP      string
	ClientPort    string // 客户端口
	LoginTime     uint64 // 登陆时间
	HeartbeatTime uint64 // 心跳
	LoginOutTime  uint64 `gorm:"column:login_out_time"` // 下线时间
	IsLogout      bool   //	是否下线
	DeviceInfo    string //设备信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model((&user)).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
