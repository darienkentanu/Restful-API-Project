package config

import (
	"fmt"
	"altastore/models"

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

func initMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.CartItem{})
	db.AutoMigrate(&models.CheckoutItem{})
	db.AutoMigrate(&models.Transaction{})
}
