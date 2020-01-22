package main

import (
	"flag"

	"github.com/MishaNiki/tcpchat/internal/app/tcpchat"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/tcpchat.json", "path to config file")
}

func main() {
	flag.Parse()

	config := tcpchat.NewConfig()

	//err := config.DecodeJFile(configPath)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(config.BindPort)
	serv := tcpchat.New(config)
	if err := serv.Start(); err != nil {
		panic(err)
	}

}
