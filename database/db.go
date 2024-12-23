package database

import (
	"context"
	"demo-go/app/model"
	"demo-go/common/global"
	"demo-go/config"
	"demo-go/database/drivers"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"

	"sync"
	"time"

	"gorm.io/gorm"
)

var mu sync.Mutex

// InitDB 初始化数据库连接
func InitDB() error {
	cfg := global.Config

	// 初始化数据库连接 gorm
	fmt.Println("gorm-database-config-initall", cfg.Databases)
	_, err := InitGormDB(cfg.Databases.Driver, cfg.Databases.Host, cfg.Databases.User, cfg.Databases.Password, cfg.Databases.DBName, cfg.Databases.Port, cfg.Databases.MaxIdleConns, cfg.Databases.MaxOpenConns, cfg.Databases.ConnMaxLife)
	if err != nil {
		log.Fatalf("Failed to connect to gorm: %v", err)
		return err
	}

	// 初始化 Redis 连接
	fmt.Println("redis-config-initall", cfg.Redis)
	if cfg.Redis.Host != "" {
		_, err := InitRedisDB(cfg.Redis)
		if err != nil { // 初始化 Redis 连接失败
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	}

	// 初始化其他连接（es、kafka、mongodb）等 todo

	return nil
}

// InitRedisDB 初始化 Redis 客户端
func InitRedisDB(config config.RedisConfig) (*redis.Client, error) {
	// 全局 Redis 连接存在，则直接返回
	if global.RedisDb != nil {
		return global.RedisDb, nil
	}

	// 创建 Redis 客户端
	Addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,            // Redis 地址
		Password: config.Password, // 密码
		DB:       config.DB,       // 数据库编号
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
		return nil, err
	}

	global.RedisDb = client // 设置全局 Redis 连接 todo 这两种方式都可以
	log.Printf("Connected redissuccessfully, addr: %s, ", Addr)

	return client, nil
}

// InitGormDB 封装具体的数据库连接逻辑
func InitGormDB(driver, host, userName, password, database string, port int, maxIdle, maxOpen, maxLifetime int) (*gorm.DB, error) {
	//加锁实现线程安全
	mu.Lock()
	defer mu.Unlock()

	// 全局数据库连接存在，则直接返回
	if global.GormDB != nil {
		return global.GormDB, nil
	}

	// 创建数据库连接
	var dialector gorm.Dialector
	switch driver {
	case "mysql": // mysql 数据库
		dialector = drivers.NewMysqlDialector(host, userName, password, database, port)
	case "postgres": // postgres 数据库
		dialector = drivers.NewPostgresDialector(host, userName, password, database, port)
	case "kingbase": // kingbase 数据库-人大金仓
		dialector = drivers.NewKingbaseDialector(host, userName, password, database, port)
	case "dm": //达梦数据库
		dialector = drivers.NewDmDialector(host, userName, password, database, port)
	default:
		return nil, errors.New("gorm Database connection not found")
	}

	// 连接数据库
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	go func() {
		err := autoMigrate(db)
		if err != nil {
			log.Printf("auto migrate failed: %v", err)
		}
	}()

	// 设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdle)                                     //设置最大空闲连接数
	sqlDB.SetMaxOpenConns(maxOpen)                                     //设置可打开的最大连接数为 100 个
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second) //设置一个连接空闲后在多长时间内可复用

	//数据库连接健康检查
	// 数据库健康检查
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// 添加到全局管理器
	global.GormDB = db // 设置全局数据库连接
	log.Printf("Connected to gorm successfully, driver: %s", driver)
	return db, nil
}

// autoMigrate 自动迁移表结构
func autoMigrate(db *gorm.DB) error {
	// 自动迁移表结构
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
