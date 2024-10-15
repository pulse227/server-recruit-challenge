package mysqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Initialize(user, pass, addr, name string) (*sql.DB, error) {
	c := mysql.Config{
		User:      user,
		Passwd:    pass,
		Net:       "tcp",
		Addr:      addr,
		DBName:    name,
		ParseTime: true,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
