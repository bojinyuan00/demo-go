package utils

import (
	"errors"
	"fmt"
	"runtime"
)

// StringInSlice string in slices(判断字符串是否在切片内)
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ToInt64 interface To Int64 类型转int64
func ToInt64(val interface{}) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		var intValue int64
		_, err := fmt.Sscanf(v, "%d", &intValue)
		if err != nil {
			return 0, errors.New("invalid string for conversion")
		}
		return intValue, nil
	default:
		return 0, errors.New("unsupported type")
	}
}

// GetCallerInfo 获取调用日志函数的文件名、行号和方法名
func GetCallerInfo(skip int) (file string, line int, function string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0, "unknown"
	}
	fn := runtime.FuncForPC(pc)
	function = "unknown"
	if fn != nil {
		function = fn.Name()
	}
	return file, line, function
}
