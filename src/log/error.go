package log

import (
	"fmt"
	"os"
	"runtime"
)

func NewError(message string) {
	_, file, line, _ := runtime.Caller(1)
	err := fmt.Errorf("[%s][%d] ATC-Error: %s", file, line, message)
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
