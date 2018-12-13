package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect(host string,
	port string,
	user string,
	password string,
) (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "")
}
