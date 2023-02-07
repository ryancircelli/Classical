package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Class struct {
	ID        int64
	className string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "myClasses",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	classes, err := classesByName("CEN3031")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", classes)

}

// albumsByArtist queries for albums that have the specified artist name.
func classesByName(name string) ([]Class, error) {
	// An albums slice to hold data from returned rows.
	var classes []Class

	rows, err := db.Query("SELECT * FROM class WHERE className = ?", name)
	if err != nil {
		return nil, fmt.Errorf("classesByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cla Class
		if err := rows.Scan(&cla.ID, &cla.className); err != nil {
			return nil, fmt.Errorf("classesByName %q: %v", name, err)
		}
		classes = append(classes, cla)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("classesByName %q: %v", name, err)
	}
	return classes, nil
}
