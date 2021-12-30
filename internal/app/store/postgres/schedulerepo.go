package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/app/model"
	"github.com/ulstu-schedule/reserver/internal/app/store"
	"time"
)

const groupScheduleRepoName = "groups_schedule"

var _ store.GroupScheduleRepository = (*GroupScheduleRepository)(nil)

type GroupScheduleRepository struct {
	store *ScheduleStore
}

func (r *GroupScheduleRepository) GetAllSchedules() ([]model.GroupSchedule, error) {
	var students []model.GroupSchedule
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", groupScheduleRepoName)
	err := r.store.db.Select(&students, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *GroupScheduleRepository) GetSchedule(groupName string) (*model.GroupSchedule, error) {
	var schedule model.GroupSchedule
	query := fmt.Sprintf("SELECT * FROM %s WHERE group_name=$1", groupScheduleRepoName)
	err := r.store.db.Get(&schedule, query, groupName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// if the group schedule is not in the database
	if schedule.Name == "" {
		return nil, nil
	}

	return &schedule, nil
}

func (r *GroupScheduleRepository) AddSchedule(groupName string, updateTime time.Time, info types.JSONText) {
	query := fmt.Sprintf("INSERT INTO %s (group_name, update_time, info) VALUES ($1, $2, $3)", groupScheduleRepoName)
	r.store.db.MustExec(query, groupName, updateTime, info)
}

func (r *GroupScheduleRepository) UpdateSchedule(groupName string, updateTime time.Time, info types.JSONText) {
	query := fmt.Sprintf("UPDATE %s SET update_time=$2, info=$3 WHERE group_name=$1", groupScheduleRepoName)
	r.store.db.MustExec(query, groupName, updateTime, info)
}

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

var _ store.TeacherScheduleRepository = (*TeacherScheduleRepository)(nil)

type TeacherScheduleRepository struct {
	// TODO: сделать по примеру GroupScheduleRepository
}
