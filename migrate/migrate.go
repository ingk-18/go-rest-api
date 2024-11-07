package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"log"
)

func main() {
	// データベース接続を取得
	dbConn := db.NewDB() // 1つの戻り値を受け取る
	if dbConn == nil {
		log.Fatal("failed to connect to database")
	}
	defer db.CloseDB(dbConn)

	// マイグレーションの実行
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Successfully Migrated")
}