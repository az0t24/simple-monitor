package main

import (
	"encoding/json"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"time"
)

func printStats() {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	for _, s := range stat.CPUStats {
		v, _ := json.Marshal(s)
		// s.User
		// s.Nice
		// s.System
		// s.Idle
		// s.IOWait
		log.Printf("%s", v)
	}
}

func main() {
	for {
		printStats()
		time.Sleep(10 * time.Second)
	}
}
