package main

import (
	"fmt"
	"github.com/robfig/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet(host string) {
	resp, err := http.Get(host + "/get?userName=mingzhehao&password=123456")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func httpPost(host string) {
	resp, err := http.Post(host+"/post",
		"application/x-www-form-urlencoded",
		strings.NewReader("userName=mingzhehao&password=123456"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func httpPostForm(host string) {
	resp, err := http.PostForm(host+"/get.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}

func httpDo(host string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", host+"/post.php", strings.NewReader("name=go"))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=go")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func main() {
	c, _ := config.ReadDefault("config/config.ini")
	ip, _ := c.String("SERVER", "base-url")
	port, _ := c.String("SERVER", "port")
	url := ip + ":" + port
	httpGet(url)
	httpPost(url)
	//httpPostForm()
	//httpDo()
}
