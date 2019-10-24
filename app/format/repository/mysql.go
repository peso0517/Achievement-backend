package repository

import (
	// gorm "Achievement/backend/app/db/database"
	"errors"
	model "Achievement/backend/app/format/domain"
	"log"

	"github.com/jinzhu/gorm" 
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
) 

type formatMysql struct{}

// NewFormatMysql is ..
func NewFormatMysql() FormatRepository {
	return &formatMysql{}
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "kokusara517"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "achievement"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Println(err, "gorm error")
		panic(err.Error())
	}
	return db
}

func (d *formatMysql) Create(data *model.ReceiveFormatData) error {
	db := gormConnect()
	defer db.Close()
	
	log.Println(data, "mysql--------data----------")

	err := db.Table("formats").
											Create(&data).
											Error
	if err != nil {
		log.Println(err, "mysql--------data----------")
				return err
	}
	return errors.New("作成できていません。")
}
