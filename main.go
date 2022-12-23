package main

import (
	"fmt"
	server "goserver/cmd/httpserver"
	"goserver/cmd/logs"
)

func main() {
	fmt.Println("Server Started")
	logs.SetLogs()
	server.HTTPserver()
}
