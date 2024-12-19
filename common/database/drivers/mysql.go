package drivers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMysqlDialector 创建mysql数据库驱动
func NewMysqlDialector(host, userName, password, database string, port int) gorm.Dialector {
	dsn := BuildMysqlDsn(host, userName, password, database, port)
	fmt.Println("mysql数据库连接信息dsn：", dsn)
	return mysql.Open(dsn)
}

func BuildMysqlDsn(host, userName, password, database string, port int) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, database)
}
