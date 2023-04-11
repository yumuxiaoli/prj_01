package models

import "gorm.io/gorm"

// 消息
type Message struct {
	gorm.Model
	FormId   string // 发送者
	TargetId string // 接收者
	Type     string // 消息类型 群聊 私聊 广播
	Media    int    // 消息类型 文字 图片 音频
	Content  string // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}
