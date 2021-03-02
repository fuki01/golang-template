package main

import (
	"golang-template/db"
	"golang-template/server"
)

func main() {
	db.Init()
	server.Init()
}
