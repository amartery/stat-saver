package utility

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq" // ...
)

func CreatePostgresConnection(dbSettings string) (*sql.DB, error) {
	con, err := Open(dbSettings)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return con, nil
}

// Open ...
func Open(dbSettings string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbSettings)
	if err != nil {
		fmt.Println("open db")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("ping db")
		return nil, err
	}
	return db, nil
}

// Close ...
func Close(con *sql.DB) {
	con.Close()
}
