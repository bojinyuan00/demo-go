package utils

import (
	"fmt"
	"runtime"
)

func CaptureError(err error, msg string) error {
	//判断错误是否为空
	if err == nil {
		return nil
	}
	//获取调用栈信息
	pc, file, line, ok := runtime.Caller(1) // 获取调用栈信息
	if !ok {
		return err
	}
	//获取函数名
	fn := runtime.FuncForPC(pc)
	function := "unknown"
	if fn != nil {
		function = fn.Name()
	}

	//返回带错误信息的错误
	return fmt.Errorf("detail: %v, custom_message: %s, error_file: %s, error_line:%d, error_function: %s", err, msg, file, line, function)
}
