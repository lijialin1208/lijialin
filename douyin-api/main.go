package main

import (
	"douyin-api/dal"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	h := server.Default(
		server.WithHostPorts(fmt.Sprint(viper.Get("server.ip"), ":", viper.Get("server.port"))),
		server.WithMaxRequestBodySize(16*1024*1024),
	)
	InitRouter(h)
	dal.Init()
	h.Spin()
}
