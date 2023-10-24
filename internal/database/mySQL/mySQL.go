package mysql

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(user, pass, dbName, host string, port int) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mySQL := &MySQL{
		DB: dbConn,
	}

	return mySQL, nil
}
