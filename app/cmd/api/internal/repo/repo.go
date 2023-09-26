package repo

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/config"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/query"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var GlobalRepository *Repository
var GlobalDB *gorm.DB

type Repository struct {
	Model      *query.Query // gen生成的model模型
	Redis      *redis.Redis // redis
	LocalCache *cache.Cache
}

func NewRepository(c config.Config) *Repository {
	r := &Repository{
		Model:      query.Use(initMysql(c)),
		Redis:      initRedis(c),
		LocalCache: cache.New(5*time.Minute, 10*time.Minute), //创建一个默认过期时间为5分钟的缓存 每10分钟清除过期项
	}
	GlobalRepository = r
	return r
}

func initRedis(c config.Config) *redis.Redis {
	// redis连接
	return redis.MustNewRedis(
		redis.RedisConf{
			Host: c.Redis.Host,
			Pass: c.Redis.Pass,
			Type: c.Redis.Type,
		},
	)
}

func initMysql(c config.Config) *gorm.DB {
	m := c.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		m.UserName,
		m.Password,
		m.Path,
		m.Port,
		m.DataBase,
		m.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	// gorm配置
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	GlobalDB = db
	return GlobalDB
}
