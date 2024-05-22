package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlOpts struct {
	User, Password, Host, DBName string
}

func NewMysqlConnection(opts MysqlOpts) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s/%s", opts.User, opts.Password, opts.Host, opts.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open the database: %s", err)
		os.Exit(1)
	}

	return db
}
