package store

import (
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/model"
	"time"
)


type ScheduleRepository interface {

	// Information executes AddSchedule if there is no schedule in the table and executes UpdateSchedule if there is in the table.
	Information(name string, updateTime time.Time, info types.JSONText) error

	// AddSchedule adds the group schedule to the database table.
	AddSchedule(name string, updateTime time.Time, info types.JSONText)

	// UpdateSchedule updates information about the group schedule.
	UpdateSchedule(name string, updateTime time.Time, info types.JSONText)
}

// GroupScheduleRepository represents a database table with backups of group schedules.
type GroupScheduleRepository interface {

	ScheduleRepository

	// GetAllSchedules returns a backup copies of all groups schedules with additional information.
	GetAllSchedules() ([]model.GroupSchedule, error)

	// GetSchedule returns a backup copy of the group schedule with additional information.
	GetSchedule(groupName string) (*model.GroupSchedule, error)

}


// TeacherScheduleRepository represents a database table with backups of teacher schedules.
type TeacherScheduleRepository interface {

	ScheduleRepository

	GetAllSchedules() ([]model.TeacherSchedule, error)

	GetSchedule(teacherName string) (*model.TeacherSchedule, error)

}
