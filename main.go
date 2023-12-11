package main

import (
	"log"
	"wms-by-gin/initdb"
	"wms-by-gin/router"
)

func main() {
	app := router.App()
	// Connect to the database.
	db, err := initdb.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Close the database connection pool after the main function returns.
	defer db.Close()
	// 启动服务
	app.Run(":8080")
}
