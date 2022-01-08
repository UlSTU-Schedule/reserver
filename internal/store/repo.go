package store

import (
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/model"
	"time"
)

// GroupScheduleRepository represents a database table with backups of group schedules.
type GroupScheduleRepository interface {
	// GetAllSchedules returns a backup copies of all groups schedules with additional information.
	GetAllSchedules() ([]model.GroupSchedule, error)

	// GetSchedule returns a backup copy of the group schedule with additional information.
	GetSchedule(groupName string) (*model.GroupSchedule, error)

	// Information executes AddSchedule if there is no schedule in the table and executes UpdateSchedule if there is in the table.
	Information(groupName string, updateTime time.Time, info types.JSONText) error

	// AddSchedule adds the group schedule to the database table.
	AddSchedule(groupName string, updateTime time.Time, info types.JSONText)

	// UpdateSchedule updates information about the group schedule.
	UpdateSchedule(groupName string, updateTime time.Time, info types.JSONText)
}

// TeacherScheduleRepository represents a database table with backups of teacher schedules.
type TeacherScheduleRepository interface {
	// TODO: сделать по примеру с группами
}
