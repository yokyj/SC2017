package main

import (
	//"os"

	"github.com/yokyj/SC2017/cloudGO-IO/service"
	// flag "github.com/spf13/pflag"
)



func main() {
	

	server := service.NewServer()
	server.Run(":8080" )
}