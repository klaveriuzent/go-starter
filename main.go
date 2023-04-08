package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("\n%s ⚡️ You are already connected to Golang!\n\n", now.Format("2006-01-02 15:04:05"))
}
