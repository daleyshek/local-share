package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	server()
}

var ports []string
var path string

func server() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	for _, p := range ports {
		go func(port string, dir string) {
			mux := http.NewServeMux()
			mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
			http.ListenAndServe(port, mux)
		}(p, dir)
	}
	fmt.Println("服务运行中")
	select {}
}

func init() {
	fmt.Println("请输入需要监听的端口（多个端口用;分开，默认监听80、88端口）:")
	var portScan string
	n, _ := fmt.Scanln(&portScan)
	if n != 0 {
		ports = strings.Split(portScan, ";")
		for i := range ports {
			if ports[i][0] != ':' {
				ports[i] = ":" + ports[i]
			}
		}
	} else {
		ports = []string{":80", ":88"}
	}

	fmt.Println("请输入监听URL的路径（默认监听默认路径 / ）")
	n, _ = fmt.Scanln(&path)
	if n != 0 {
		if path[0] != '/' {
			path = "/" + path
		}
		if path[len(path)-1] != '/' {
			path = path + "/"
		}
	} else {
		path = "/"
	}
}
