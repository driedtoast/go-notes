package datamapper

import (
	"bytes"
	"database/sql"
	_ "github.com/lib/pq"
	"helpers"
)

func DbConnection() (*sql.DB, error) {
	cfg := helpers.Config()

	var buffer bytes.Buffer
	buffer.WriteString(" user=")
	dbuser, _ := cfg.Get("db.user")
	buffer.WriteString(dbuser)
	buffer.WriteString(" password=")
	dbpw, _ := cfg.Get("db.password")
	buffer.WriteString(dbpw)
	buffer.WriteString(" dbname=")
	dbname, _ := cfg.Get("db.dbname")
	buffer.WriteString(dbname)
	buffer.WriteString(" sslmode=disable")

	db, err := sql.Open("postgres", buffer.String())
	return db, err
}
