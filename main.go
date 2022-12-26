package main

import (
	"fmt"
	server "goserver/cmd/httpserver"
	"goserver/cmd/logs"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Server Started")
	logs.SetLogs()
	cobra.AddTemplateFunc
	server.HTTPserver()
}
