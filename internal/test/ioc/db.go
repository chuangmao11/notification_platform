package ioc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ecodeclub/ekit/retry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func WaitForDBSetup(dsn string) {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	const maxInterval = 10 * time.Second
	const maxRetries = 10
	strategy, err := retry.NewExponentialBackoffRetryStrategy(time.Second, maxInterval, maxRetries)
	if err != nil {
		panic(err)
	}
	const timeout = 5 * time.Second
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		// 这里是进行真正的连接测试
		err = sqlDB.PingContext(ctx)
		cancel()
		if err == nil {
			break
		}
		next, ok := strategy.Next()
		if !ok {
			panic("WaitForDBSetup重试失败")
		}
		time.Sleep(next)
	}
}

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:13316)/notification?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s"
	WaitForDBSetup(dsn)
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic(fmt.Errorf("数据库连接失败:%w", err))
	}
	return db
}
