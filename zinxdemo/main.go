package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("os.GetWd()执行异常%v", err)
			return
		}
		fmt.Println(dir)
	}()

	select {
	case <-time.After(time.Second * 20):
		break
	}

}
