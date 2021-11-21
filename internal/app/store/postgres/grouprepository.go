package postgres

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/app/model"
	"time"
)

// GroupScheduleRepository ...
type GroupScheduleRepository struct {
	store *Store
}

// GetAllSchedules ...
func (r *GroupScheduleRepository) GetAllSchedules() ([]model.GroupSchedule, error) {
	var students []model.GroupSchedule
	err := r.store.db.Select(&students, "SELECT * FROM groups_schedule ORDER BY id")
	if err != nil {
		return nil, err
	}
	return students, nil
}

// GetSchedule ...
func (r *GroupScheduleRepository) GetSchedule(groupName string) (*model.GroupSchedule, error) {
	var schedule model.GroupSchedule
	err := r.store.db.Get(&schedule, "SELECT * FROM groups_schedule WHERE group_name=$1", groupName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// if the group schedule is not in the database
	if schedule.Name == "" {
		return nil, nil
	}

	return &schedule, nil
}

// AddSchedule ...
func (r *GroupScheduleRepository) AddSchedule(groupName string, updateTime time.Time, info types.JSONText) {
	r.store.db.MustExec("INSERT INTO groups_schedule (group_name, update_time, info) VALUES ($1, $2, $3)",
		groupName, updateTime, info)
}

// UpdateSchedule ...
func (r *GroupScheduleRepository) UpdateSchedule(groupName string, updateTime time.Time, info types.JSONText) {
	r.store.db.MustExec("UPDATE groups_schedule SET update_time=$2, info=$3 WHERE group_name=$1",
		groupName, updateTime, info)
}

// Information ...
func (r *GroupScheduleRepository) Information(groupName string, updateTime time.Time, info types.JSONText) error {
	schedule, err := r.GetSchedule(groupName)
	if err != nil {
		return err
	}

	if schedule != nil {
		r.UpdateSchedule(groupName, updateTime, info)
	} else {
		r.AddSchedule(groupName, updateTime, info)
	}
	return nil
}
