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

	logger.Infof("The reserver is running! Reservation every %d hours.", cfg.ReservationIntervalGroups)

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(cfg.ReservationIntervalGroups).Hours().Do(reserveGroupsSchedules, store, logger)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Every(cfg.ReservationIntervalTeachers).Hours().Do(reserveTeachersSchedules, store, logger)
	if err != nil {
		log.Fatal(err)
	}

	s.StartAsync()
}

// reserveGroupsSchedules loads schedules of all UlSTU groups into the database.
func reserveGroupsSchedules(store *postgres.ScheduleStore, logger *logrus.Logger) {
	logger.Info("Reservation of group schedules is started.")

	groups := schedule.GetGroups()
	for _, group := range groups {
		fullScheduleNew, err := schedule.GetFullGroupSchedule(group)
		if err != nil {
			logger.Errorf("[GROUP] %s: %v", group, err)
			continue
		}

		fullScheduleOldDB, err := store.GroupSchedule().GetSchedule(group)
		if err != nil {
			logger.Errorf("[GROUP] %s: %v", group, err)
			continue
		}

		fullScheduleOld := &types.Schedule{}
		if fullScheduleOldDB != nil {
			err = easyjson.Unmarshal(fullScheduleOldDB.Info, fullScheduleOld)
			if err != nil {
				logger.Errorf("[GROUP] %s: %v", group, err)
				continue
			}
		}

		hasOldScheduleChanged := false

		for idx, weekSchedule := range fullScheduleNew.Weeks {
			if !schedule.IsWeekScheduleEmpty(weekSchedule) {
				fullScheduleOld.Weeks[idx] = weekSchedule
				logger.Infof("[GROUP] %s: updated %d week schedule", group, idx+1)
				hasOldScheduleChanged = true
			}
		}

		if !hasOldScheduleChanged {
			logger.Infof("[GROUP] %s: schedule has not been updated", group)
			continue
		}

		bytes, err := easyjson.Marshal(fullScheduleOld)
		if err != nil {
			logger.Errorf("[GROUP] %s: %v", group, err)
			continue
		}

		err = store.GroupSchedule().Information(group, time.Now(), bytes)
		if err != nil {
			logger.Errorf("[GROUP] %s: %v", group, err)
			continue
		}

		time.Sleep(time.Second * 4) // DDOS-attack: off :D
	}

	logger.Info("Reservation of group schedules is completed.")
}

// reserveTeachersSchedules loads schedules of all UlSTU teachers into the database.
func reserveTeachersSchedules(store *postgres.ScheduleStore, logger *logrus.Logger) {
	logger.Info("Reservation of teacher schedules is started.")

	teachers, err := schedule.GetTeachers()

	if err != nil {
		logger.Errorf("Error getting all teachers: %v", err)
		return
	}

	for _, teacher := range teachers {
		fullScheduleNew, err := schedule.GetFullTeacherSchedule(teacher)

		if err != nil {
			logger.Errorf("[TEACHER] %s: %v", teacher, err)
			continue
		}

		fullScheduleOldDB, err := store.TeacherSchedule().GetSchedule(teacher)
		if err != nil {
			logger.Errorf("[TEACHER] %s: %v", teacher, err)
			continue
		}

		fullScheduleOld := &types.Schedule{}
		if fullScheduleOldDB != nil {
			err = easyjson.Unmarshal(fullScheduleOldDB.Info, fullScheduleOld)
			if err != nil {
				logger.Errorf("[TEACHER] %s: %v", teacher, err)
				continue
			}
		}

		hasOldScheduleChanged := false

		for idx, weekSchedule := range fullScheduleNew.Weeks {
			if !schedule.IsWeekScheduleEmpty(weekSchedule) {
				fullScheduleOld.Weeks[idx] = weekSchedule
				logger.Infof("[TEACHER] %s: updated %d week schedule", teacher, idx+1)
				hasOldScheduleChanged = true
			}
		}

		if !hasOldScheduleChanged {
			logger.Infof("[TEACHER] %s: schedule has not been updated", teacher)
			continue
		}

		bytes, err := easyjson.Marshal(fullScheduleOld)
		if err != nil {
			logger.Errorf("[TEACHER] %s: %v", teacher, err)
			continue
		}

		err = store.TeacherSchedule().Information(teacher, time.Now(), bytes)
		if err != nil {
			logger.Errorf("[TEACHER] %s: %v", teacher, err)
			continue
		}

		time.Sleep(time.Second * 4)
	}

	logger.Info("Reservation of teacher schedules is completed.")
}