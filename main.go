package main

import (
	//"fmt"
	"./controller"

)

func main() {
	controller := controller.NewController("http://www.ifeng.com",1)
	//fmt.Println(controller.StartUrl)
	controller.Do()

}
