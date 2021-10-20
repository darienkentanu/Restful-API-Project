package config

import (
	"altastore/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfig() (config map[string]string) {
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conf
}

func InitDB() *gorm.DB {
	conf := GetConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf["DB_Username"], conf["DB_Password"], conf["DB_Host"],
		conf["DB_Port"], conf["DB_Name"],
	)
	var err error
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initMigration(db)
	return db
}

func InitDBSQL() *sql.DB {
	conf := GetConfig()
	connStr := fmt.Sprintf("%s:%s@/%s", conf["DB_Username"], conf["DB_Password"], conf["DB_Name"])
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func initMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.CartItem{})
	db.AutoMigrate(&models.Checkout{})
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.PaymentDetail{})
}
