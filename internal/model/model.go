package model

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

// GroupSchedule represents an entry in the database table with backups of the group schedule.
type GroupSchedule struct {
	ID         int
	Name       string         `db:"group_name"`
	UpdateTime time.Time      `db:"update_time"`
	Info       types.JSONText `db:"info"`
}

// TeacherSchedule represents an entry in the database table with backups of the teacher schedule.
type TeacherSchedule struct {
	ID         int
	Name       string         `db:"teacher_name"`
	UpdateTime time.Time      `db:"update_time"`
	Info       types.JSONText `db:"info"`
}
