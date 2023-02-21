package serviceuser

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"
	"time"

	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	daonotice "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_notice"
	daouser "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_user"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/util"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
	"github.com/sanyewudezhuzi/E-COMMERCE/service/upload"
	"gopkg.in/mail.v2"
)

type UserRegisterService struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	NickName string `json:"nick_name" form:"nick_name"`
	Key      string `json:"key" form:"key"` // 前端验证
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	// 1-绑定邮箱 2-解绑邮箱 3-修改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
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
	_, exist := userDao.ExistOrNotByAccount(s.Account)
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

func (s *UserRegisterService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := daouser.NewUserDao(ctx)
	user, exist := userDao.ExistOrNotByAccount(s.Account)
	if !exist {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	if !user.Check(s.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	tokenStr, err := util.GenerateToken(user.ID, s.Account, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: tokenStr,
		},
		Msg: e.GetMsg(code),
	}
}

func (s *UserRegisterService) Update(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	// 找到这个用户
	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	// 修改 nick_name
	if s.NickName != "" {
		user.NickName = s.NickName
	}
	err = userDao.UpdateUserByID(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildUser(user),
		Msg:        e.GetMsg(code),
	}
}

func (s *UserRegisterService) Upload(ctx context.Context, uid uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	filepath, err := upload.UploadAvatarToLocalStatic(file, uid, user.Account)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	user.Avatar = filepath
	if err := userDao.UpdateUserByID(uid, user); err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildUser(user),
		Msg:        e.GetMsg(code),
	}
}

func (s *SendEmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string
	emailTokenStr, err := util.GenerateEmailToken(uid, s.Email, s.Password, s.OperationType)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	noticeDao := daonotice.NewNoticeDao(ctx)
	notice, err := noticeDao.GetNoticeByID(s.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	address = conf.ValidEmail + emailTokenStr // 发送方
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", s.Email)
	m.SetHeader("Subject", "E-COMMERCE")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}

func (s *ValidEmailService) Valid(ctx context.Context, tokenStr string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	// 验证 token
	if tokenStr == "" {
		code = e.InvaildParams
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	} else {
		claims, err := util.ParseEmailToken(tokenStr)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	// 获取该用户的信息
	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	if operationType == 1 {
		// 绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		// 修改密码
		err = user.Bcrypt(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				StatusCode: code,
				Msg:        e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserByID(userID, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildUser(user),
		Msg:        e.GetMsg(code),
	}
}

func (s *ShowMoneyService) ShowMoney(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := daouser.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildMoney(user, s.Key),
		Msg:        e.GetMsg(code),
	}
}
