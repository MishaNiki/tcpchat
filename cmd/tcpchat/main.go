package main

import (
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/tcpchat.json", "path to config file")
}

func main() {
	flag.Parse()

	

	
}