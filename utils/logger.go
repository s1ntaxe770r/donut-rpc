package utils

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// NewDonutLogger returns a new instance of a donut logger
func NewDonutLogger() *log.Logger {
	prefix := color.YellowString("[Donut-Server]")
	loggr := log.New(os.Stdout, prefix, log.LstdFlags)
	return loggr
}

func NewDBLogger() *log.Logger {
	prefix := color.YellowString("[DATABASE]")
	LG := log.New(os.Stdout, prefix, log.Lshortfile)
	return LG
}
