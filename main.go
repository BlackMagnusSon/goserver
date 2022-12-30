package main

import (
	"fmt"
	server "goserver/cmd/httpserver"
	"goserver/cmd/logs"
	"goserver/cmd/telegram"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Server Started")
	logs.SetLogs()
	go server.HTTPserver()
	token := os.Getenv("TG_TOKEN")
	url := "https://api.telegram.org/"
	go telegram.TelegramParrot(&url, &token)
	fmt.Println("Running")
	for {
		fmt.Println("Gosched")
		time.Sleep(1 * time.Second)
		runtime.Gosched()
	}
}
