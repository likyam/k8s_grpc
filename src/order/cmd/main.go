package main

import (
	"flag"
	"likyam.cn/src/order/cmd/server"
)

var cfg = flag.String("config", "src/order/config/config.yaml", "config file location")

// main
func main() {
	flag.Parse()
	server.Run(*cfg)
}
