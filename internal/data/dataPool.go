package data

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	// errDataConfIncomplete 代表数据库配置不完整
	errDataConfIncomplete = errors.New("the database configuration is incomplete ")
	// errDataConnIsNil conn返回为空
	errDataConnIsNil = errors.New("conn returns is nil")
	// errDataHandleIsNil 数据库句柄返回为空
	errDataHandleIsNil = errors.New("database handle returns is nil")
)

var (
	// db连接
	db *gorm.DB
	//db配置
	dbConf *conf = nil
)

type conf struct {
	UserName     string `mapstructure:"username"`      // 数据库用户名
	Password     string `mapstructure:"password"`      // 连接密码
	Host         string `mapstructure:"host"`          // 地址
	Port         string `mapstructure:"port"`          // 端口
	DatabaseName string `mapstructure:"database_name"` // 数据库名称
}

//后期可能会添加配置实时更新功能
func getConf() (c *conf, err error) {
	if dbConf != nil {
		c = dbConf
		return
	}
	err = viper.UnmarshalKey("app.database", &c)
	if err != nil {
		return
	}
	dbConf = c
	if dbConf.UserName == "" || dbConf.Password == "" || dbConf.Host == "" || dbConf.Port == "" || dbConf.DatabaseName == "" {
		err = errDataConfIncomplete
		return
	}
	return
}

// setup 初始化连接
func setup() (err error) {
	c, err := getConf()
	if err != nil {
		return
	}

	var dbURI string
	var dialect gorm.Dialector
	dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName)
	dialect = mysql.New(mysql.Config{
		DSN:                       dbURI, // 地址
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})

	conn, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return
	}
	if conn == nil {
		err = errDataConnIsNil
		return
	}
	sqlDB, err := conn.DB()
	if err != nil {
		return
	}
	if sqlDB == nil {
		err = errDataHandleIsNil
		return
	}
	//后期改成配置化
	sqlDB.SetMaxIdleConns(10)             // SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxOpenConns(100)            // SetMaxOpenConns 设置到数据库的最大打开连接数。
	sqlDB.SetConnMaxLifetime(time.Minute) // SetConnMaxLifetime 设置可重用连接的最大时间。
	db = conn
	return
}

// GetDB 开放给外部获得db连接
func GetDB() (*gorm.DB, error) {
	if db == nil {
		err := setup()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
