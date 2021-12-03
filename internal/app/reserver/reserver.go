package reserver

import (
	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/ulstu-schedule/parser/schedule"
	"github.com/ulstu-schedule/reserver/internal/app/store/postgres"
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

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(config.ReservationInterval).Hours().Do(reserveGroupsSchedules, store, logger)
	if err != nil {
		return err
	}
	s.StartBlocking()

	return nil
}

// reserveGroupsSchedules loads schedules of all UlSTU groups into the database.
func reserveGroupsSchedules(store *postgres.Store, logger *logrus.Logger) {
	logger.Info("Reservation of group schedules is started.")

	groups := schedule.GetGroups()
	for _, group := range groups {
		fullGroupSchedule, err := schedule.GetFullGroupSchedule(group)
		if err != nil {
			logger.Error(err)
		}

		bytes, err := easyjson.Marshal(fullGroupSchedule)
		if err != nil {
			logger.Error(err)
		}

		err = store.GroupSchedule().Information(group, time.Now(), bytes)
		if err != nil {
			logger.Error(err)
		}

		logger.Infof("%s OK", group)

		time.Sleep(time.Second) // DDOS-attack: off :D
	}

	logger.Info("Reservation of group schedules is completed.")
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
