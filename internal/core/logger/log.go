package logger

import (
	"fmt"
	"os"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var gray = "\033[37m"

func Info(text string) {
	fmt.Println(green + "ℹ INFO: " + reset + text)
}

func ProgressInfo(text string) {
	fmt.Println(gray + "∴ INIT: " + reset + text)
}

func ProgressOk() {
	fmt.Println(green + "  ↳ ✔ OK!" + reset)
}

func Error(text string) {
	fmt.Fprintln(os.Stderr, red+"✘ ERROR: "+reset+text)
}

func Fatal(text string) {
	fmt.Fprintln(os.Stderr, red+"💥 FATAL: "+text+reset)
}

func ErrorWithDetail(text string, err error) {
	fmt.Fprintln(os.Stderr, red+"✘ ERROR: "+reset+text)
	fmt.Fprintln(os.Stderr, gray+"  ↳ ⚙ DETAIL: "+err.Error())
}

func FatalWithDetail(text string, err error) {
	fmt.Fprintln(os.Stderr, red+"💥 FATAL: "+text+reset)
	fmt.Fprintln(os.Stderr, gray+"   ↳ ⚙ DETAIL: "+err.Error())
}

func Warn(text string) {
	fmt.Println(yellow + "⚠ WARNING: " + reset + text)
}

func Debug(text string) {
	fmt.Println(blue + "⚙ DEBUG: " + text + reset)
}
