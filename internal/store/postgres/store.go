package postgres

import (
	"github.com/jmoiron/sqlx"
)

func NewDB(databaseURL string) (*sqlx.DB, error) {
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

type ScheduleStore struct {
	db              *sqlx.DB
	groupSchedule   *GroupScheduleRepository
	teacherSchedule *TeacherScheduleRepository
}

func NewScheduleStore(db *sqlx.DB) *ScheduleStore {
	return &ScheduleStore{
		db: db,
	}
}

func (s *ScheduleStore) GroupSchedule() *GroupScheduleRepository {
	if s.groupSchedule != nil {
		return s.groupSchedule
	}

	s.groupSchedule = &GroupScheduleRepository{
		store: s,
	}

	return s.groupSchedule
}

func (s *ScheduleStore) TeacherSchedule() *TeacherScheduleRepository {
	if s.teacherSchedule != nil {
		return s.teacherSchedule
	}

	s.teacherSchedule = &TeacherScheduleRepository{
		store: s,
	}

	return s.teacherSchedule
}
