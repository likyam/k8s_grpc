package main

import (
	"flag"
	"likyam.cn/src/user/cmd/server"
)

var cfg = flag.String("config", "/src/user/config/config.yaml", "config file location")

// main
func main() {
	flag.Parse()
	server.Run(*cfg)
}
