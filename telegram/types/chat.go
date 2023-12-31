package types

import (
	"encoding/json"
)

const (
	// 私聊
	ChatPrivate = "private"
	// 普通群
	ChatGroup = "group"
	// 超级群
	SuperGroup = "supergroup"
	// 频道
	Channel = "channel"
)

// 聊天信息
type Chat struct {
	ID        int64   `json:"id"`                 // 聊天唯一ID
	Type      string  `json:"type"`               // 聊天类型
	FirstName string  `json:"first_name"`         // 用户昵称
	UserName  *string `json:"username,omitempty"` // 用户名
}

// 转换为JSON
func (chat *Chat) ToJSON() ([]byte, error) {
	return json.Marshal(chat)
}

// 从JSON反序列化
func (chat *Chat) FromJSON(jsb []byte) error {
	return json.Unmarshal(jsb, chat)
}
