package postgres

import (
	"github.com/jmoiron/sqlx"
)

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
	// TODO: сделать по примеру GroupSchedule()
	panic("implement me!")
}
