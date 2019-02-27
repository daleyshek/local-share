package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	server()
}

func server() {
	ports := []string{":80", ":88"}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	for _, p := range ports {
		fmt.Println(fmt.Sprintf("监听%v端口", p))
		go func(port string, dir string) {
			mux := http.NewServeMux()
			mux.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir(dir))))
			http.ListenAndServe(port, mux)
		}(p, dir)
	}
	select {}
}
