package main

import (
	"ghast/readflag"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	readflag.Get(args)
	if(len(args) == 0) {
	fmt.Println("ghast - A configurable command-line streaming tool\nArguments:\n	-v, --version - current version\n	-c, --config - provide config to run") 
	}
}
