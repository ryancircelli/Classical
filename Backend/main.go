package main

import (
	"Classical/Backend/testMain"
)

func main() {
	var a = testMain.App{}
	// {username} {password} {database name}
	a.Initialize("root", "password123", "classical")

	a.Run(":8000")
}
