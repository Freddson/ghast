package main

import (
	"ghast/readflag"
	"fmt"
	"os"
	"gopkg.in/ini.v1"
)

func main() {
	args := os.Args[1:]
	readflag.Get(args)
	//fmt.Println(os.Getenv("configPath"))
	if(len(args) == 0) {
	fmt.Println("ghast - A configurable command-line streaming tool\nArguments:\n	-v, --version - current version\n	-c, --config - provide config to run") 
	}

	if(len(os.Getenv("configPath")) == 0) {
		fmt.Println("Config path is empty or a directory!")
	}

	readConf(os.Getenv("configPath"))
}

func readConf(filepath string) {
	conf, err := ini.Load(filepath)
	if (err != nil) {
		fmt.Printf("Failed to read config file.\n%v", err)
		fmt.Printf("\n")
		os.Exit(1)
	}

	fmt.Println(conf.Section("Connection").Key("AccessToken").String())
}