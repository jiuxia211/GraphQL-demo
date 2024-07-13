package db

import (
	"log"

	"github.com/jiuxia211/GraphQL-demo/pkg/constants"
	"github.com/jiuxia211/GraphQL-demo/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(utils.GetMysqlDsn()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatalf("mysql open error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("mysql setting error: %v", err)
	}

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间

	DB = db

	if err != nil {
		log.Fatalf("mysql: 创建管理员用户失败%v\n", err)
	}
}
