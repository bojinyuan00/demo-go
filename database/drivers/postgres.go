package drivers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDialector postgres数据库驱动
func NewPostgresDialector(host, userName, password, database string, port int) gorm.Dialector {
	dsn := BuildPostgresDsn(host, userName, password, database, port)
	fmt.Println("postgres数据库连接信息dsn：", dsn)
	return postgres.Open(dsn)
}

func BuildPostgresDsn(host, userName, password, database string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, database)
}
