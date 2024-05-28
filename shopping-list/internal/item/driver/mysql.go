package driver

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"shopping-list/internal/item/driver/logger"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlOpts struct {
	User, Password, Host, DBName string
	Port                         int
}

func NewMysqlConnection(ctx context.Context, opts MysqlOpts) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.DBName,
	)

	logger.FromContext(ctx).Debug("[database] DSN built", "DSN", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open the database: %s\n", err)
		os.Exit(1)
	}

	return db
}
