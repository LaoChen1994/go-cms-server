package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"pd-go-server/pkg/setting"
	"time"
)

var DB *gorm.DB

var tx *gorm.DB = nil

type Model struct {
	ID        uint       `gorm:"primary_key" gorm:"column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deletedAt"`
}

func Init() {
	conf := setting.DatabaseConf
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseConf.Prefix,
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("db open is error", err)
	}
}

func GetDB() *gorm.DB {
	if tx != nil {
		return tx
	}

	return DB
}

func UseTransaction(txDb *gorm.DB) {
	if txDb != nil {
		tx = txDb
	}
}

func TransactionEnd() {
	tx = nil
}
