package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.GetWd()执行异常%v", err)
		return
	}
	fmt.Println(dir)
	select {}

}
