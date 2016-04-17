package client

import (
	"fmt"
	"github.com/sy-vendor/public"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

func init() {
	httpTransport := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(30 * time.Second)
			c, err := net.DialTimeout(netw, addr, 20*time.Second)
			if err != nil {
				return nil, err
			}

			c.SetDeadline(deadline)
			return c, nil
		},
		DisableKeepAlives: false,
	}
	httpClient = &http.Client{
		Transport: httpTransport,
	}
}

func HttpGet(host string) {
	resp, err := http.Get(host + "/get?userName=mingzhehao&password=123456")
	if err != nil {
		panic(err)
	}

	//important
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func HttpPost(host string) {
	resp, err := http.Post(host+"/post",
		"application/x-www-form-urlencoded",
		strings.NewReader("userName=mingzhehao&password=123456"))
	if err != nil {
		fmt.Println(err)
	}

	//important
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

/**
 * sign认证
 */
func HttpSign(host string) {
	var uid string = "111111"
	var rid string = "1"
	timestamp := time.Now().Unix()
	timestamp_string := strconv.FormatInt(timestamp, 10) //转换为string
	params_str, sign_str := public.MakeParams(uid, rid, timestamp_string)
	fmt.Println("params_str: ", params_str)
	fmt.Println("sign_str: ", sign_str)
	sign := public.MakeSign(sign_str)

	resp, err := httpClient.Post(host+"/sign",
		"application/x-www-form-urlencoded",
		strings.NewReader(params_str+"&sign="+sign))
	if err != nil {
		fmt.Println(err)
		return
	}

	//important
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func HttpPostForm(host string) {
	resp, err := http.PostForm(host+"/get.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		panic(err)
	}

	//important
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}

func HttpDo(host string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", host+"/post.php", strings.NewReader("name=go"))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=go")

	resp, err := client.Do(req)

	//important
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

//func main() {
//	c, _ := config.ReadDefault("config/config.ini")
//	ip, _ := c.String("SERVER", "base-url")
//	port, _ := c.String("SERVER", "port")
//	url := ip + ":" + port
//	HttpGet(url)
//	HttpPost(url)
//	//HttpSign(url)
//	//HttpDo()
//}
