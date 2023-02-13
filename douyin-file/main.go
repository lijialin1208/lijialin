package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("C:/Users/32259/GolandProjects/douyin/douyin-file/public/play_url")))
	http.ListenAndServe("192.168.0.103:8081", nil)
}
