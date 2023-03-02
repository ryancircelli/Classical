package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}
	// {username} {password} {database name}
	a.Initialize("root", "password123", "classical")

	a.Run(":8080")
}
