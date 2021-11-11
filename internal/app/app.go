package app

import (
	"fmt"
	"github.com/ulstu-schedule/parser/schedule"
	"log"
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
		groupSchedule, err := schedule.GetTextDailyGroupSchedule("ПИбд-21", 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(groupSchedule)
	}
}
