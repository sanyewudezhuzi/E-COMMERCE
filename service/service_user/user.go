package serviceuser

import (
	"context"
	"errors"

	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

type UserRegisterService struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	NickName string `json:"nick_name" form:"nick_name"`
	Key      string `json:"key" form:"key"` // 前端验证
}

func (s *UserRegisterService) Register(ctx context.Context) serializer.Response {
	// 创建 user
	var user model.User
	code := e.Success
	// 验证参数合法性
	if len(s.Key) != 16 {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      errors.New("Key format error."),
		}
	}
	// 新用户赠送 10000 金币，将金额对称加密
	util.Encrypt.SetKey(s.Key)
	// 数据持久化
	userDao := daouser.NewUserDao(ctx)
	exist := userDao.ExistOrNotByAccount(s.Account)
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	user = model.User{
		Account:  s.Account,
		NickName: s.NickName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("10000"),
	}
	// 密码加密
	if err := user.Bcrypt(s.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	// 创建用户
	if err := userDao.CreateUser(&user); err != nil {
		code = e.Error
	}
	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}
