package main

import (
	"fmt"
	"time"
)

func main() {
	go worker(time.NewTicker(5 * time.Second))
	select {}
}

func worker(ticker *time.Ticker) {
	for range ticker.C {
		fmt.Println(time.Now())
	}
}
