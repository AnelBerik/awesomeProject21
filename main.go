package main

import (
	"awesomeProject21/config"
	"fmt"
)

func main()  {
	conf:=config.GetConfig()
	fmt.Println(conf)

}
