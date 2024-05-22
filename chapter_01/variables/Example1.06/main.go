package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// type only
	var start, middle, end float32
	fmt.Println(start, middle, end)

	// initial value mixed type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	// works with function also
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
