package app

import (
	"github.com/go-co-op/gocron"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/ulstu-schedule/parser/schedule"
	"github.com/ulstu-schedule/parser/types"
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
		fullScheduleNew, err := schedule.GetFullGroupSchedule(group)
		if err != nil {
			logger.Errorf("%s: %v", group, err)
			continue
		}

		fullScheduleOldDB, err := store.GroupSchedule().GetSchedule(group)
		if err != nil {
			logger.Errorf("%s: %v", group, err)
			continue
		}

		fullScheduleOld := &types.Schedule{}
		err = easyjson.Unmarshal(fullScheduleOldDB.Info, fullScheduleOld)
		if err != nil {
			logger.Errorf("%s: %v", group, err)
			continue
		}

		hasOldScheduleChanged := false

		firstWeekScheduleNew := fullScheduleNew.Weeks[0]
		if !schedule.IsWeekScheduleEmpty(firstWeekScheduleNew) {
			fullScheduleOld.Weeks[0] = firstWeekScheduleNew
			logger.Infof("%s: updated the first week schedule", group)
			hasOldScheduleChanged = true
		}

		secondWeekScheduleNew := fullScheduleNew.Weeks[1]
		if !schedule.IsWeekScheduleEmpty(secondWeekScheduleNew) {
			fullScheduleOld.Weeks[1] = secondWeekScheduleNew
			logger.Infof("%s: updated the second week schedule", group)
			hasOldScheduleChanged = true
		}

		if !hasOldScheduleChanged {
			logger.Infof("%s: schedule has not been updated", group)
			continue
		}

		bytes, err := easyjson.Marshal(fullScheduleOld)
		if err != nil {
			logger.Errorf("%s: %v", group, err)
			continue
		}

		err = store.GroupSchedule().Information(group, time.Now(), bytes)
		if err != nil {
			logger.Errorf("%s: %v", group, err)
			continue
		}

		time.Sleep(time.Second * 3) // DDOS-attack: off :D
	}

	logger.Info("Reservation of group schedules is completed.")
}
