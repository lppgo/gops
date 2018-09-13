package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gops/agent"
)

// https://github.com/google/gops
// gops is a command to list and diagonse(诊断)Go process currently running to your system.
func main() {
	fmt.Println("-----start-----")
	if err := agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----sleep-----")
	time.Sleep(time.Hour)
}
