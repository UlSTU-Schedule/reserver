package app

import (
	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
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
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	store := postgres.NewScheduleStore(db)

	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLevel(level)

	logger.Infof("The reserver is running! Groups reservation every %d hours and teachers reservation every %d hours.",
		cfg.ReservationIntervalGroups, cfg.ReservationIntervalTeachers)

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(cfg.ReservationIntervalGroups).Hours().Do(reserveGroupsSchedules, store, logger)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Every(cfg.ReservationIntervalTeachers).Hours().Do(reserveTeachersSchedules, store, logger)
	if err != nil {
		log.Fatal(err)
	}

	s.StartBlocking()
}

// reserveGroupsSchedules loads schedules of all UlSTU groups into the database.
func reserveGroupsSchedules(store *postgres.ScheduleStore, logger *logrus.Logger) {
	logger.Info("[GROUPS] Reservation is started.")

	groups := schedule.GetGroups()
	for _, group := range groups {
		fullScheduleNew, err := schedule.GetFullGroupSchedule(group)
		if err != nil {
			logger.Errorf("[GROUPS] %s: %v", group, err)
			continue
		}

		fullScheduleOldDB, err := store.GroupSchedule().GetSchedule(group)
		if err != nil {
			logger.Errorf("[GROUPS] %s: %v", group, err)
			continue
		}

		var firstWeekUpdateTimeNew, secondWeekUpdateTimeNew time.Time
		fullScheduleOld := &types.Schedule{}

		if fullScheduleOldDB != nil {
			err = easyjson.Unmarshal(fullScheduleOldDB.FullSchedule, fullScheduleOld)
			if err != nil {
				logger.Errorf("[GROUPS] %s: %v", group, err)
				continue
			}

			firstWeekUpdateTimeNew = fullScheduleOldDB.FirstWeekUpdateTime
			secondWeekUpdateTimeNew = fullScheduleOldDB.SecondWeekUpdateTime
		}

		hasOldFirstWeekScheduleChanged := false
		if !schedule.IsWeekScheduleEmpty(fullScheduleNew.Weeks[0]) {
			fullScheduleOld.Weeks[0] = fullScheduleNew.Weeks[0]
			firstWeekUpdateTimeNew = time.Now()

			logger.Infof("[GROUPS] %s: updated 1 week schedule", group)
			hasOldFirstWeekScheduleChanged = true
		}

		hasOldSecondWeekScheduleChanged := false
		if !schedule.IsWeekScheduleEmpty(fullScheduleNew.Weeks[1]) {
			fullScheduleOld.Weeks[1] = fullScheduleNew.Weeks[1]
			secondWeekUpdateTimeNew = time.Now()

			logger.Infof("[GROUPS] %s: updated 2 week schedule", group)
			hasOldSecondWeekScheduleChanged = true
		}

		if !hasOldFirstWeekScheduleChanged && !hasOldSecondWeekScheduleChanged {
			logger.Infof("[GROUPS] %s: schedule has not been updated", group)
			continue
		}

		bytes, err := easyjson.Marshal(fullScheduleOld)
		if err != nil {
			logger.Errorf("[GROUPS] %s: %v", group, err)
			continue
		}

		if err = store.GroupSchedule().Information(group, firstWeekUpdateTimeNew, secondWeekUpdateTimeNew, bytes); err != nil {
			logger.Errorf("[GROUPS] %s: %v", group, err)
			continue
		}

		time.Sleep(time.Second * 4) // DDOS-attack: off :D
	}

	logger.Info("[GROUPS] Reservation is completed.")
}

// reserveTeachersSchedules loads schedules of all UlSTU teachers into the database.
func reserveTeachersSchedules(store *postgres.ScheduleStore, logger *logrus.Logger) {
	logger.Info("[TEACHERS] Reservation is started.")

	teachers, err := schedule.GetTeachers()
	if err != nil {
		logger.Errorf("[TEACHERS] error occured while parsing all teachers: %v", err)
		return
	}

	for _, teacher := range teachers {
		fullScheduleNew, err := schedule.GetFullTeacherSchedule(teacher)
		if err != nil {
			logger.Errorf("[TEACHERS] %s: %v", teacher, err)
			continue
		}

		fullScheduleOldDB, err := store.TeacherSchedule().GetSchedule(teacher)
		if err != nil {
			logger.Errorf("[TEACHERS] %s: %v", teacher, err)
			continue
		}

		var firstWeekUpdateTimeNew, secondWeekUpdateTimeNew time.Time
		fullScheduleOld := &types.Schedule{}

		if fullScheduleOldDB != nil {
			err = easyjson.Unmarshal(fullScheduleOldDB.FullSchedule, fullScheduleOld)
			if err != nil {
				logger.Errorf("[TEACHERS] %s: %v", teacher, err)
				continue
			}

			firstWeekUpdateTimeNew = fullScheduleOldDB.FirstWeekUpdateTime
			secondWeekUpdateTimeNew = fullScheduleOldDB.SecondWeekUpdateTime
		}

		hasOldFirstWeekScheduleChanged := false
		if !schedule.IsWeekScheduleEmpty(fullScheduleNew.Weeks[0]) {
			fullScheduleOld.Weeks[0] = fullScheduleNew.Weeks[0]
			firstWeekUpdateTimeNew = time.Now()

			logger.Infof("[TEACHERS] %s: updated 1 week schedule", teacher)
			hasOldFirstWeekScheduleChanged = true
		}

		hasOldSecondWeekScheduleChanged := false
		if !schedule.IsWeekScheduleEmpty(fullScheduleNew.Weeks[1]) {
			fullScheduleOld.Weeks[1] = fullScheduleNew.Weeks[1]
			secondWeekUpdateTimeNew = time.Now()

			logger.Infof("[TEACHERS] %s: updated 2 week schedule", teacher)
			hasOldSecondWeekScheduleChanged = true
		}

		if !hasOldFirstWeekScheduleChanged && !hasOldSecondWeekScheduleChanged {
			logger.Infof("[TEACHERS] %s: schedule has not been updated", teacher)
			continue
		}

		bytes, err := easyjson.Marshal(fullScheduleOld)
		if err != nil {
			logger.Errorf("[TEACHERS] %s: %v", teacher, err)
			continue
		}

		if err = store.TeacherSchedule().Information(teacher, firstWeekUpdateTimeNew, secondWeekUpdateTimeNew, bytes); err != nil {
			logger.Errorf("[TEACHERS] %s: %v", teacher, err)
			continue
		}

		time.Sleep(time.Second * 4)
	}

	logger.Info("[TEACHERS] Reservation is completed.")
}
