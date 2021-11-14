package reserver

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ulstu-schedule/parser/schedule"
	"github.com/ulstu-schedule/reserver/internal/app/repository"
	"log"
	"time"
)

// Reserver ...
type Reserver struct {
	config     *Config
	logger     *logrus.Logger
	repository *repository.Repository
}

// New ...
func New(config *Config) *Reserver {
	return &Reserver{
		config: config,
		logger: logrus.New(),
	}
}

// Run runs worker.
func (r *Reserver) Run() error {
	if err := r.configureLogger(); err != nil {
		return err
	}

	if err := r.configureRepository(); err != nil {
		return err
	}

	r.logger.Infof("The program is running! Reservation every %d hours.", r.config.ReservationInterval)

	go worker(time.NewTicker(time.Duration(r.config.ReservationInterval) * time.Second))
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

// configureLogger ...
func (r *Reserver) configureLogger() error {
	level, err := logrus.ParseLevel(r.config.LogLevel)
	if err != nil {
		return err
	}

	r.logger.SetLevel(level)
	return nil
}

// configureRepository ...
func (r *Reserver) configureRepository() error {
	rep := repository.New(r.config.Repository)
	if err := rep.Open(); err != nil {
		return err
	}

	r.repository = rep
	return nil
}
