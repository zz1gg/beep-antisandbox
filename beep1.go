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

func isSandbox() bool {
	start := time.Now()
	// Set frequency 750Hz and duration 60ms
	procBeep.Call(750, 60)

	duration := time.Since(start)

	fmt.Println("[+] Actual Duration: ", duration)

	if duration < 60*time.Millisecond {
		fmt.Println("[+] Set Duration: ", 60*time.Millisecond)
		return true
	}
	return false
}

func main() {

	if isSandbox() {
		fmt.Println("[-] Running in a sandbox environment!")
	} else {
		fmt.Println("[!] Not running in a sandbox environment!")
		fmt.Println("[+] Beep Beep Beep!")
	}
}
