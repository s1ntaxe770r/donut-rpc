package utils

import (
	"bytes"
	"log"

	"github.com/fatih/color"
)

func NewDonutLogger() *log.Logger {
	var buf bytes.Buffer
	LG := log.New(&buf, "[Donut-Server]", log.Lshortfile)
	return LG
}

func NewDBLogger() *log.Logger {
	prefix := color.YellowString("[DATABASE]")
	var buf bytes.Buffer
	LG := log.New(&buf, prefix, log.Lshortfile)
	return LG
}
