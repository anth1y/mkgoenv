package main

import (
	"os"
	//	"syscall"
	"fmt"
)

func main() {
	home := os.Getenv("HOME")
	fmt.Println(home)
}
