package serializer

import (
	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
)

type User struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// 序列化 user
func BuildUser(user *model.User) User {
	return User{
		ID:       user.ID,
		Account:  user.Account,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
