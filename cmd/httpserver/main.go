package main

import (
	"fmt"
	"os"

	. "goserver"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("httpserver starts now...")

	err := ServerStart()
	if err != nil {
		fmt.Printf("Fail to start httpserver :%s\n", err.Error())
		os.Exit(1)
	}
}
