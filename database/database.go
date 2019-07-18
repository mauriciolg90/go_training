package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectSql(driver, user, pass, host, schema string) (*sql.DB, error) {
    return sql.Open(driver, user + ":" + pass + "@" + host + "/" + schema)
}
