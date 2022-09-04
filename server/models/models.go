package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pd-go-server/pkg/setting"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" gorm:"column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func init() {
	conf := setting.DatabaseConf
	var err error

	DB, err = gorm.Open(conf.DatabaseType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Name))

	if err != nil {
		fmt.Println("db open is error", err)
	}

	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(50)
	DB.DB().SetMaxIdleConns(50)
}

func close() {
	DB.Close()
}
