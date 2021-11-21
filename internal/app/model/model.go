package model

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

// GroupSchedule ...
type GroupSchedule struct {
	ID         int
	Name       string         `db:"group_name"`
	UpdateTime time.Time      `db:"update_time"`
	Info       types.JSONText `db:"info"`
}

// TeacherSchedule ...
type TeacherSchedule struct {
	// TODO: сделать по аналогии с GroupSchedule
}
