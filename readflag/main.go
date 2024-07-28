package readflag

import (
	"fmt";
	"strings";
	"os";
)

func sort(argtosort string) {
	switch argtosort {
	case "-v", "--version": fmt.Println("ghast version 0.01")
	case "-c", "--config": fmt.Println("Waiting for config to arrive..."); os.Setenv("waitForConf", "true")
	default: 
		if(os.Getenv("waitForConf") == "true") {
			if (strings.Contains(argtosort, ".ini")) { fmt.Println("Config arrived!"); os.Setenv("waitForConf", "false"); os.Setenv("configPath", argtosort)} else if (!strings.Contains(argtosort, ".ini")) {
				fmt.Println("Provided config file is in the wrong format! Cancelling...") }} else if (os.Getenv("waitForConf") != "true") {
			if(strings.Contains(argtosort[:1], "-")) {
				fmt.Println("Invalid argument!", argtosort) }
			if(strings.Contains(argtosort, ".ini")) {
				fmt.Println("Not waiting for config, cancelling...", argtosort) } else if (!strings.Contains(argtosort, ".ini")) {
						fmt.Println("Invalid argument!", argtosort) } else if (!strings.Contains(argtosort[:1], "-")) {
							fmt.Println(argtosort, "DEBUG!")
						}}	
	
	//if i have to write this long of a switch case again, please kill me
}}

func Get(arguments []string) {
	for i := 0; i < len(arguments); i++ {
		sort(arguments[i])
	}
}