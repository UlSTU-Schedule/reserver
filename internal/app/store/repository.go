package store

import (
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/app/model"
	"time"
)

// GroupScheduleRepository ...
type GroupScheduleRepository interface {
	GetAllSchedules() ([]model.GroupSchedule, error)
	GetSchedule(groupName string) (*model.GroupSchedule, error)
	AddSchedule(groupName string, updateTime time.Time, info types.JSONText)
	UpdateSchedule(groupName string, updateTime time.Time, info types.JSONText)
	Information(groupName string, updateTime time.Time, info types.JSONText) error
}

// TeacherScheduleRepository ...
type TeacherScheduleRepository interface {
	// TODO: сделать по примеру с группами
}
