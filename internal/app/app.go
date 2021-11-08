package app

import (
	"fmt"
	"time"
)

// Run runs worker.
func Run() {
	go worker(time.NewTicker(5 * time.Second))
	select {}
}

// worker ...
func worker(ticker *time.Ticker) {
	for range ticker.C {
		fmt.Println(time.Now())
	}
}
