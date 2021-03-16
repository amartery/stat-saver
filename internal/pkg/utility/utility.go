package utility

import (
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
)

func CreatePostgresConnection(dbSettings string) (*sqlx.DB, error) {
	con, err := Open(dbSettings)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return con, nil
}

// InitDatabase ...
func InitDatabase(con *sqlx.DB) error {
	sqlInitFile, err := ioutil.ReadFile("scripts/statistics.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = con.Exec(string(sqlInitFile))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Open ...
func Open(dbSettings string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbSettings)
	if err != nil {
		fmt.Println("open db")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("ping db")
		return nil, err
	}
	if err := InitDatabase(db); err != nil {
		fmt.Println("init db")
		return nil, err
	}
	return db, nil
}

// Close ...
func Close(con *sqlx.DB) {
	con.Close()
}
