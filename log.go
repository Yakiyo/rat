package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func initLogger() {
	env := os.Getenv("RAT_LOG")
	fmt.Println(env)
	if env == "" {
		env = "error"
	}
	level, err := log.ParseLevel(env)
	if err != nil {
		level = log.ErrorLevel
	}
	log.SetLevel(level)
	log.SetTimeFormat("")
}
