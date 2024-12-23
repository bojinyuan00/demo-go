package product_dao

import (
	"demo-go/app/model"
	"demo-go/common/global"
	"demo-go/common/log"
	"demo-go/common/utils"
	"time"
)

// Dao UserDao 数据访问对象
type Dao struct {
}

// NewDao 创建 ProductDao 实例
func NewDao() *Dao {
	return &Dao{}
}

// GetUserById 根据ID获取用户信息
/**
 * @description 根据ID获取用户信息
 * @param id 用户ID int
 * @return user用户信息,
 *       err错误信息,
 *       customErr自定义错误信息（用户调用栈信息的收集）,
 * 		 message提示语（错误信息逐层抛到最外层，用户接口返回错误提示语）
 * @throws 异常抛出 message
 * @author 卜锦元
 * @date 2024/12/21
 */
func (u *Dao) GetUserById(id int) (user *model.User, err error, customErr error, message string) {
	start := time.Now() //访问时间
	logParams := make(map[string]interface{})
	logParams["id"] = id
	defer func() { log.TimeTracker(start, logParams, user, customErr) }() // defer 捕获日志并执行

	//查询数据库
	result := global.GormDB.First(&user, id)
	err = result.Error

	//对于错误的处理（1、记录异常日志；2、往上一层抛出自定义错误；3、返回错误信息（用户提示语等））
	if err != nil {
		customErr = utils.CaptureError(err, "获取用户信息失败-UserDao.GetUserById")
		return nil, err, customErr, "用户信息不存在"
	}
	return user, nil, nil, ""
}
