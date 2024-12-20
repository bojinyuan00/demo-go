package initall

import (
	"demo-go/database"
)

func dbInit() error {
	//initialize database connection
	err := database.InitDB()
	if err != nil {
		return err
	}
	return nil
}
