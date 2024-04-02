package db

import (
	logs "SearchEngine/pkg/logger"
	"SearchEngine/repository/mysql/model"
	"os"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.InputData{},
		&model.Favorite{},
		&model.FavoriteDetail{},
	)
	if err != nil {
		logs.LogrusObj.Infoln("register tabel failed")
		os.Exit(0)
	}
	logs.LogrusObj.Infoln("register table success")

}
