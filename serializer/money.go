package serializer

import (
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
)

type Money struct {
	UserID   uint   `json:"user_id"`
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	Money    string `json:"money"`
}

func BuildMoney(user *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserID:   user.ID,
		Account:  user.Account,
		NickName: user.NickName,
		Money:    util.Encrypt.AesDecoding(user.Money),
	}
}
