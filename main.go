package main

import (
	"fmt"
	"vault-go/service"
)

func main() {
	service.InitConfiguration()

	testKey := service.GetConfiguration("test")

	hello := service.GetConfiguration("hello")

	fmt.Println(testKey)
	fmt.Println(hello)
}
