package main

import (
	"finego/handler"
	"fmt"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpPort string `yaml:"http_port"`
}

func main() {
	r := gin.Default()

	data, err := ioutil.ReadFile("conf/dev.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config is %+v\n", config)

	r.GET("/ping", handler.HelloWorld)
	r.Run(config.HttpPort)
}
