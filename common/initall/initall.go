package initall

// All initall all
func All() {
	//load config 文件
	err := loadConfig("./common/config")
	if err != nil {
		return
	}

	//init logger
	loggerInit()

	//init db
	err = dbInit()
	if err != nil {
		return
	}

	//init router
	routerInit()
}
