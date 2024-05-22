package main

import (
	"fmt"
	"time"
)

//var (
//	Debugs       bool
//	LogLevels    = "info"
//	startUpTimes = time.Now()
//)

func main() {
	Debugs, LogLevels, startUpTimes := false, "info", time.Now()

	fmt.Println(Debugs, LogLevels, startUpTimes)
}
