package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ulstu-schedule/reserver/internal/app/store"
)

// Store ...
type Store struct {
	db              *sqlx.DB
	groupSchedule   *GroupScheduleRepository
	teacherSchedule *TeacherScheduleRepository
}

// New ...
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// GroupSchedule ...
func (s *Store) GroupSchedule() store.GroupScheduleRepository {
	if s.groupSchedule != nil {
		return s.groupSchedule
	}

	s.groupSchedule = &GroupScheduleRepository{
		store: s,
	}

	return s.groupSchedule
}

// TeacherSchedule ...
func (s Store) TeacherSchedule() store.TeacherScheduleRepository {
	// TODO: сделать по примеру GroupSchedule()
	return nil
}
