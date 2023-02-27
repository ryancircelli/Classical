package service

import (
	"Classical/Backend/db"
	obj "Classical/Backend/model"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

// classesByName queries for albums that have the specified class name.
func ClassesByName(name string) ([]obj.Class, error) {
	// An albums slice to hold data from returned rows.
	var classes []obj.Class

	rows, err := db.DB.Query("SELECT * FROM class WHERE className = ?", name)
	if err != nil {
		return nil, f.Errorf("classesByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cla obj.Class
		if err := rows.Scan(&cla.ID, &cla.ClassName); err != nil {
			return nil, f.Errorf("classesByName %q: %v", name, err)
		}
		classes = append(classes, cla)
	}
	if err := rows.Err(); err != nil {
		return nil, f.Errorf("classesByName %q: %v", name, err)
	}
	return classes, nil
}
