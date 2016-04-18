package main

import (
	"fmt"
	"github.com/robfig/config"
	"github.com/sy-vendor/server"
	log "github.com/thinkboy/log4go"
)

func main() {
	c, _ := config.ReadDefault("config/config.ini")
	host, _ := c.String("SERVER", "host")
	port, _ := c.String("SERVER", "port")
	url := host + ":" + port
	fmt.Println("\n当前服务器IP及端口为：", url)
	log.LoadConfiguration("config/log-conf.xml")
	server.WebServerInit(url)
}
