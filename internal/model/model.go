package model

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

// GroupSchedule represents an entry in the database table with backups of the group schedule.
type GroupSchedule struct {
	ID                   int
	Name                 string         `db:"group_name"`
	FirstWeekUpdateTime  time.Time      `db:"first_week_update_time"`
	SecondWeekUpdateTime time.Time      `db:"second_week_update_time"`
	FullSchedule         types.JSONText `db:"full_schedule"`
}

// TeacherSchedule represents an entry in the database table with backups of the teacher schedule.
type TeacherSchedule struct {
	ID                   int
	Name                 string         `db:"teacher_name"`
	FirstWeekUpdateTime  time.Time      `db:"first_week_update_time"`
	SecondWeekUpdateTime time.Time      `db:"second_week_update_time"`
	FullSchedule         types.JSONText `db:"full_schedule"`
}
