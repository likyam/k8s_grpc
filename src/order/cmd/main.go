package main

import (
	"flag"
	"likyam.cn/src/order/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file location")

// main
func main() {
	flag.Parse()
	server.Run(*cfg)
}
