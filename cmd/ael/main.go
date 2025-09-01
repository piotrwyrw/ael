package main

import (
	"ael/internal/cli"
	"ael/internal/config"
	"fmt"
)

func main() {
	_, cfg := config.ReadConfiguration(true)
	err := cli.InitializeCLI(cfg)
	if err != nil {
		fmt.Println("Error Encountered: " + err.Error())
	}
}
