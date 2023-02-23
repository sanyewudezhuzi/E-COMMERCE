package cache

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDB     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic("Failed to load conf file.")
	}
	load_redis(file)
	connect_redis()
}

func load_redis(f *ini.File) {
	RedisDB = f.Section("redis").Key("RedisDB").String()
	RedisAddr = f.Section("redis").Key("RedisAddr").String()
	RedisPw = f.Section("redis").Key("RedisPw").String()
	RedisDbName = f.Section("redis").Key("RedisDbName").String()

}

func connect_redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		// password
		DB: int(db),
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
