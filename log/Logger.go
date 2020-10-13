package log

import (
	"fmt"
	"time"
)

var logLevel = 2 // @TODO: Configurable in the settings

func Info(s string) {
	print(2, "\x1b[1;32mINFO\x1b[0m", s)
}

func Warn(s string) {
	print(1, "\x1b[1;93mWARN\x1b[0m", s)
}

func Fatal(s string) {
	print(0, "\x1b[1;31mFATAL\x1b[0m", s)
}

func InfoAlwaysShown(s string) {
	print(-1, "\x1b[1;32mINFO\x1b[0m", s)
}

func print(level int, msgType, s string) {
	if logLevel >= level {
		fmt.Printf("[%v] %v - %v\n", msgType, time.Now().Format("2006-01-02 15:04:05"), s)
	}
}