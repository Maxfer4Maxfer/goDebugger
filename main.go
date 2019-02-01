package goDebugger

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// DebugMode determines the mode of a current debug
var DebugMode = true

const debugSymbol = "✎ "

// GetCurrentFunctionName returns the name of a called function
func GetCurrentFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
	return f.Name()
}

// DebugTimeStamp prints a function timing.
// Time since a program has been started.
// Duration of a function execution.
func DebugTimeStamp(start time.Time, tabCount int, fName string, fInputs ...interface{}) {
	defer func(start time.Time, fInputs ...interface{}) {
		duration := time.Since(start)
		end := time.Now()
		re := regexp.MustCompile(`m+.*`)
		result := debugSymbol + fmt.Sprintf("%v%v() start %v end %v duration(ns) %v \t input = ",
			strings.Repeat("  ", tabCount),
			fName,
			strings.Trim(re.FindString(start.String()), "m=+"),
			strings.Trim(re.FindString(end.String()), "m=+"),
			duration.Nanoseconds(),
			// duration,
		)
		for _, input := range fInputs {
			result = result + fmt.Sprintf("%v ", input)
		}
		result = result + fmt.Sprintf("\n")
		fmt.Print(result)
	}(start, fInputs...)
}
