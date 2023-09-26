package configs

import "prakerja11/models/user"

func initMigrate(){
	DB.AutoMigrate(&user.User{})
}