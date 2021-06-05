package utils

import (
	"bytes"
	"log"
)

func NewDonutLogger() *log.Logger {
	var buf bytes.Buffer
	LG := log.New(&buf, "[Donut-Server]", log.Lshortfile)
	return LG
}
