package main

import (
	"github.com/GitHubLeoYeh/autoTableGomod/tablecreator"
	"github.com/GitHubLeoYeh/autoTableGomod/user"
	"log"
)

func main() {
	// 資料庫連接字串 (請根據你的環境修改)
	dsn := "root:root@tcp(127.0.0.1:53306)/try_it?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化 Creator
	creator, err := tablecreator.NewCreator(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// 自動建表
	if err := creator.AutoMigrate(&user.UserInfo{}); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}
