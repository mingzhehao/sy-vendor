package server

import (
	"fmt"
	"net/http"
)

func WebServerInit(url string) {
	fmt.Println("This is webserver base!")

	/**
	 * 第一个参数为客户端发起http请求时的接口名
	 * 第二个参数是一个func，负责处理这个请求。
	 */
	http.HandleFunc("/get", getTask)
	http.HandleFunc("/post", postTask)
	http.HandleFunc("/sign", signTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe(url, nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}
