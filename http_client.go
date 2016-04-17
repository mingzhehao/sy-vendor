package main

import (
	"github.com/robfig/config"
	"github.com/sy-vendor/client"
)

func main() {
	c, _ := config.ReadDefault("config/config.ini")
	ip, _ := c.String("SERVER", "base-url")
	port, _ := c.String("SERVER", "port")
	url := ip + ":" + port
	client.HttpGet(url)
	client.HttpPost(url)
	client.HttpSign(url)
}
