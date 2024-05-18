package main

import (
	"fmt"
	"syscall"
	"time"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procBeep = kernel32.NewProc("Beep")
)

func main() {

	start := time.Now()
	// Set frequency 750Hz and duration 60ms
	procBeep.Call(750, 60)
	fmt.Println("[+] Set Duration: ", 60*time.Millisecond)
	duration := time.Since(start)
	fmt.Println("[+] Actual Duration: ", duration)
	fmt.Println("[+] Beep Beep Beep!")
}
