package main

import (
	//"os"

	"github.com/yokyj/cloudgo-data/service"
	// flag "github.com/spf13/pflag"
)



func main() {
	

	server := service.NewServer()
	server.Run(":8080")
}