package model

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

// 配置 mysql
func ConnectMySQL(dsn_read, dsn_write string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn_read, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，mysql5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，mysql5.7 之前的数据库不支持
		DontSupportRenameColumn:   true,     // 用`change`重命名列，mysql8 之前的数据库不支持
		SkipInitializeWithVersion: false,    // 根据当前 mysql 版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数化
		},
	})
	if err != nil {
		panic("Failed to connect mysql.")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	DB = db

	// 主从配置
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn_write)},                      // 主数据库写
		Replicas: []gorm.Dialector{mysql.Open(dsn_read), mysql.Open(dsn_read)}, // 从数据库读写
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		panic("Failed to register dbresolver.")
	}
}

// 数据库表自动迁移
func AutomigrateMySQL() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&Address{},
		&Admin{},
		&Carousel{},
		&Cart{},
		&Category{},
		&Favorite{},
		&Notice{},
		&Order{},
		&ProductImg{},
		&Product{},
		&User{},
	)
	if err != nil {
		panic("Failed to automigrate.")
	}
}

// 将 DB 的上下文更改为 ctx
func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
