package main

import (
	"usautoparts/stockpicker/config"
	"fmt"
)

func main() {


	config, err := config.Load("config.json");
	fmt.Println("%v %v", config, err)

}
