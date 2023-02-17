package conf

import (
	"strings"

	"github.com/sanyewudezhuzi/E-COMMERCE/model"

	"gopkg.in/ini.v1"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDB     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

// 加载环境变量
func LoadConf() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic("Failed to load conf file.")
	}

	loadService(file)
	loadMysql(file)
	loadRedis(file)
	loadEmail(file)
	loadPath(file)

	// connect mysql
	mysql_dsn_read := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	mysql_dsn_write := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	model.ConnectMySQL(mysql_dsn_read, mysql_dsn_write)
}

func loadService(f *ini.File) {
	AppModel = f.Section("service").Key("AppMode").String()
	HttpPort = f.Section("service").Key("HttpPort").String()
}

func loadMysql(f *ini.File) {
	DB = f.Section("mysql").Key("DB").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}

func loadRedis(f *ini.File) {
	RedisDB = f.Section("redis").Key("RedisDB").String()
	RedisAddr = f.Section("redis").Key("RedisAddr").String()
	RedisPw = f.Section("redis").Key("RedisPw").String()
	RedisDbName = f.Section("redis").Key("RedisDbName").String()
}

func loadEmail(f *ini.File) {
	ValidEmail = f.Section("email").Key("ValidEmail").String()
	SmtpHost = f.Section("email").Key("SmtpHost").String()
	SmtpEmail = f.Section("email").Key("SmtpEmail").String()
	SmtpPass = f.Section("email").Key("SmtpPass").String()
}

func loadPath(f *ini.File) {
	Host = f.Section("path").Key("Host").String()
	ProductPath = f.Section("path").Key("ProductPath").String()
	AvatarPath = f.Section("path").Key("AvatarPath").String()
}
