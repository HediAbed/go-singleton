package main

import (
	"fmt"
)

// Database is the struct we want to make a singleton
type Database struct {
	serverip string
	serverport int
}

// Connect using the db instance
func (db *Database) Connect() {
	fmt.Println(db.serverip, ":", db.serverport)
}

// SetServer sets the ip and port of the db instance
func (db *Database) SetServer(ip string, port int) {
	db.serverip = ip
	db.serverport = port
}

// The db instance
var dbinstance *Database

// Database class instance
func getDatabaseInstance() *Database {
	return dbinstance
}
