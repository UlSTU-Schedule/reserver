package reserver

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/ulstu-schedule/parser/schedule"
	"github.com/ulstu-schedule/parser/types"
	"github.com/ulstu-schedule/reserver/internal/app/store/postgres"
	"log"
	"time"
)

// Run runs worker.
func Run(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := postgres.New(db)

	logger := logrus.New()
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	logger.SetLevel(level)

	logger.Infof("The program is running! Reservation every %d hours.", config.ReservationInterval)

	// Пример обращения к БД (предварительно там лежит 1 запись)
	groupSchedules, err := store.GroupSchedule().GetAllSchedules()
	if err != nil {
		return err
	}

	groupSchedule := &types.Schedule{}

	err = easyjson.Unmarshal(groupSchedules[0].Info, groupSchedule)
	if err != nil {
		return err
	}

	fmt.Println(groupSchedule)

	go worker(time.NewTicker(time.Duration(config.ReservationInterval) * time.Second))
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

// newDB ...
func newDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
