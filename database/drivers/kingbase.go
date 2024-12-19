package drivers

import (
	"fmt"
	"gorm.io/driver/postgres" // 金仓与 PostgreSQL 协议兼容
	"gorm.io/gorm"
)

// NewKingbaseDialector 人大金仓数据库驱动
func NewKingbaseDialector(host, userName, password, database string, port int) gorm.Dialector {
	dsn := BuildKingbaseDsn(host, userName, password, database, port)
	fmt.Println("人大金仓数据库连接信息dsn：", dsn)
	return postgres.Open(dsn)
}

func BuildKingbaseDsn(host, userName, password, database string, port int) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", userName, password, database, host, port)
}
