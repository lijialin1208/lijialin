package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("C:/Users/32259/GolandProjects/douyin/douyin-file/public/play_url")))
	http.ListenAndServe("192.168.1.4:8081", nil)
}
