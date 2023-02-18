package main

import (
	"fmt"

	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/router"
)

func init() {
	conf.LoadConf()
	model.AutomigrateMySQL()
	fmt.Println("CONTINUE")
}

func main() {
	fmt.Println("helloworld")
	r := router.Router()
	r.Run(conf.HttpPort)
}
