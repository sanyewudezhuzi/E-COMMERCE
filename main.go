package main

import (
	"fmt"

	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
)

func init() {
	conf.LoadConf()
	model.AutomigrateMySQL()
	fmt.Println("CONTINUE")
}

func main() {
	fmt.Println("helloworld")
}
