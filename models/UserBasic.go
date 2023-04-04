package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
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
