package dao

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //
	"github.com/jmoiron/sqlx"
	"github.com/xujintao/gorge/apps/video/conf"
)

// DB 抽象sqlx
var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Connect("mysql", conf.Config.DB.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxOpenConns(conf.Config.DB.MaxConn)

	go func() {
		if err := DB.Ping(); err != nil {
			log.Fatalln(err)
		}
		log.Println("db connected.")
		// 可以在这里建表
	}()
}
