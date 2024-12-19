package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONMap 用于支持 MySQL 的 JSON 类型 ==jsonMap
type JSONMap map[string]interface{}

// JSONStringArray 用于支持 MySQL 的 JSON 类型 ==jsonMap
type JSONStringArray []string

// Scan 将 JSON 数据从数据库加载到结构中
func (m *JSONMap) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), m)
}

// Value 将结构转化为 JSON 数据保存到数据库
func (m JSONMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan 实现 Scanner 接口，用于从数据库读取
func (j *JSONStringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

// Value 实现 Valuer 接口，用于写入数据库
func (j JSONStringArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}
