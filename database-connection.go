package main

import (
	"fmt"
)

// DatabaseConnection is the struct we want to make a singleton
type DatabaseConnection struct {
	serverip string
	serverport int
}

// Connect using the db connection instance
func (dbc *DatabaseConnection) Connect() {
	fmt.Println(dbc.serverip, ":", dbc.serverport)
}

// SetServer sets the ip and port of the db connection instance
func (dbc *DatabaseConnection) SetServer(ip string, port int) {
	dbc.serverip = ip
	dbc.serverport = port
}

// The db connection instance
var dbConnection *DatabaseConnection


func getDatabaseConnectionInstance() *DatabaseConnection {
	return dbConnection
}
