package log

import (
	"fmt"
	"time"
)

func Info(s string) {
	print("\x1b[1;32mINFO\x1b[0m", s)
}

func Warn(s string) {
	print("\x1b[1;93mWARN\x1b[0m", s)
}

func Fatal(s string) {
	print("\x1b[1;31mFATAL\x1b[0m", s)
}

func print(msgType, s string) {
	fmt.Printf("[%v] %v - %v\n", msgType, time.Now().Format("2006-01-02 15:04:05"), s)
}