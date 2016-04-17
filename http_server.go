package main

import (
	"fmt"
	"github.com/robfig/config"
	"github.com/sy-vendor/server"
)

func main() {
	c, _ := config.ReadDefault("config/config.ini")
	host, _ := c.String("SERVER", "host")
	port, _ := c.String("SERVER", "port")
	url := host + ":" + port
	fmt.Println("\n当前服务器IP及端口为：", url)
	server.WebServerInit(url)
}
