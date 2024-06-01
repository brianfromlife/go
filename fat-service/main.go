package main

import "flag"

func main() {
	dsn := flag.String("dsn", "./db.sqlite", "sqlite3 DSN")
	println(*dsn)
}
