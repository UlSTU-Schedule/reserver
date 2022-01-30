package store

import (
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/model"
	"time"
)

// ScheduleRepository represents a database tables with backups of group and teacher schedules.
type ScheduleRepository interface {
	// Information executes AddSchedule if there is no schedule in the table and executes UpdateSchedule if there is in the table.
	Information(entityName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText) error

	// AddSchedule adds the schedule to the database table.
	AddSchedule(entityName string, fullSchedule types.JSONText)

	// UpdateSchedule updates information about the schedule.
	UpdateSchedule(entityName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText)
}

// GroupScheduleRepository represents a database table with backups of group schedules.
type GroupScheduleRepository interface {
	ScheduleRepository

	// GetAllSchedules returns the backup copies of all group schedules with additional information.
	GetAllSchedules() ([]model.GroupSchedule, error)

	// GetSchedule returns the backup copy of the group schedule with additional information.
	GetSchedule(groupName string) (*model.GroupSchedule, error)
}

// TeacherScheduleRepository represents a database table with backups of teacher schedules.
type TeacherScheduleRepository interface {
	ScheduleRepository

	// GetAllSchedules returns the backup copies of all teacher schedules with additional information.
	GetAllSchedules() ([]model.TeacherSchedule, error)

	// GetSchedule returns the backup copy of the teacher schedule with additional information.
	GetSchedule(teacherName string) (*model.TeacherSchedule, error)
}
