package drivers

import (
	"fmt"
	"github.com/godoes/gorm-dameng"
	"gorm.io/gorm"
)

// NewDmDialector 达梦数据库驱动
func NewDmDialector(host, userName, password, database string, port int) gorm.Dialector {
	dsn := BuildDmDsn(host, userName, password, database, port)
	fmt.Println("达梦数据库连接信息dsn：", dsn)
	return dameng.Open(dsn)
}

func BuildDmDsn(host, userName, password, database string, port int) string {
	return fmt.Sprintf("dm://%s:%s@%s:%d?schema=%s", userName, password, host, port, database)
}

//// VARCHAR 类型大小为字符长度
////db, err := gorm.Open(dameng.New(dameng.Config{DSN: dsn, VarcharSizeIsCharLength: true}))
//// VARCHAR 类型大小为字节长度（默认）
//db, err := gorm.Open(dameng.Open(dsn), &gorm.Config{})
//if err != nil {
//	// panic error or log error info
//}
//
//// do somethings
//var versionInfo []map[string]interface{}
//db.Table("SYS.V$VERSION").Find(&versionInfo)
//if err := db.Error; err == nil {
//	versionBytes, _ := json.MarshalIndent(versionInfo, "", "  ")
//	fmt.Printf("达梦数据库版本信息：\n%s\n", versionBytes)
//}

/****************** 控制台输出内容 *****************

达梦数据库版本信息：
[
  {
    "BANNER": "DM Database Server 64 V8"
  },
  {
    "BANNER": "DB Version: 0x7000c"
  },
  {
    "BANNER": "03134284094-20230927-******-*****"
  }
]

*************************************************/
