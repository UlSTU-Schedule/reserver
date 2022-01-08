package app

import (
	"github.com/go-co-op/gocron"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/ulstu-schedule/parser/schedule"
	"github.com/ulstu-schedule/reserver/internal/config"
	"github.com/ulstu-schedule/reserver/internal/store/postgres"
	"log"
	"time"
)

// Run runs the reserver.
func Run(configsPath string) {
	cfg, err := config.New(configsPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := postgres.NewScheduleStore(db)

	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLevel(level)

	logger.Infof("The reserver is running! Reservation every %d hours.", cfg.ReservationInterval)

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(cfg.ReservationInterval).Hours().Do(reserveGroupsSchedules, store, logger)
	if err != nil {
		log.Fatal(err)
	}
	s.StartBlocking()
}

// reserveGroupsSchedules loads schedules of all UlSTU groups into the database.
func reserveGroupsSchedules(store *postgres.ScheduleStore, logger *logrus.Logger) {
	logger.Info("Reservation of group schedules is started.")

	groups := schedule.GetGroups()
	for _, group := range groups {
		fullGroupSchedule, err := schedule.GetFullGroupSchedule(group)
		if err != nil {
			logger.Errorf("%s %v", group, err)
		}

		bytes, err := easyjson.Marshal(fullGroupSchedule)
		if err != nil {
			logger.Errorf("%s %v", group, err)
		}

		err = store.GroupSchedule().Information(group, time.Now(), bytes)
		if err != nil {
			logger.Errorf("%s %v", group, err)
		}

		logger.Infof("%s OK", group)

		time.Sleep(time.Second) // DDOS-attack: off :D
	}

	logger.Info("Reservation of group schedules is completed.")
}