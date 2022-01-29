package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx/types"
	"github.com/ulstu-schedule/reserver/internal/model"
	"github.com/ulstu-schedule/reserver/internal/store"
	"time"
)

const (
	groupScheduleRepoName   = "groups_schedules"
	teacherScheduleRepoName = "teachers_schedules"
)

var (
	_ store.GroupScheduleRepository   = (*GroupScheduleRepository)(nil)
	_ store.TeacherScheduleRepository = (*TeacherScheduleRepository)(nil)
)

type GroupScheduleRepository struct {
	store *ScheduleStore
}

func (r *GroupScheduleRepository) GetAllSchedules() ([]model.GroupSchedule, error) {
	students := []model.GroupSchedule{}
	query := fmt.Sprintf("SELECT * FROM %s", groupScheduleRepoName)
	err := r.store.db.Select(&students, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *GroupScheduleRepository) GetSchedule(groupName string) (*model.GroupSchedule, error) {
	schedule := model.GroupSchedule{}
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

func (r *GroupScheduleRepository) AddSchedule(groupName string, fullSchedule types.JSONText) {
	updateTime := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (group_name, first_week_update_time, second_week_update_time, full_schedule) VALUES ($1, $2, $2, $3)", groupScheduleRepoName)
	r.store.db.MustExec(query, groupName, updateTime, fullSchedule)
}

func (r *GroupScheduleRepository) UpdateSchedule(groupName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText) {
	query := fmt.Sprintf("UPDATE %s SET first_week_update_time=$2, second_week_update_time=$3, full_schedule=$4 WHERE group_name=$1", groupScheduleRepoName)
	r.store.db.MustExec(query, groupName, firstWeekUpdateTime, secondWeekUpdateTime, fullSchedule)
}

func (r *GroupScheduleRepository) Information(groupName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText) error {
	schedule, err := r.GetSchedule(groupName)
	if err != nil {
		return err
	}

	if schedule != nil {
		r.UpdateSchedule(groupName, firstWeekUpdateTime, secondWeekUpdateTime, fullSchedule)
	} else {
		r.AddSchedule(groupName, fullSchedule)
	}
	return nil
}

type TeacherScheduleRepository struct {
	store *ScheduleStore
}

func (r *TeacherScheduleRepository) GetAllSchedules() ([]model.TeacherSchedule, error) {
	teachers := []model.TeacherSchedule{}
	query := fmt.Sprintf("SELECT * FROM %s", teacherScheduleRepoName)
	err := r.store.db.Select(&teachers, query)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeacherScheduleRepository) GetSchedule(teacherName string) (*model.TeacherSchedule, error) {
	schedule := model.TeacherSchedule{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE teacher_name=$1", teacherScheduleRepoName)
	err := r.store.db.Get(&schedule, query, teacherName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// if the teacher schedule is not in the database
	if schedule.Name == "" {
		return nil, nil
	}

	return &schedule, nil
}

func (r *TeacherScheduleRepository) AddSchedule(teacherName string, fullSchedule types.JSONText) {
	updateTime := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (teacher_name, first_week_update_time, second_week_update_time, full_schedule) VALUES ($1, $2, $2, $3)", teacherScheduleRepoName)
	r.store.db.MustExec(query, teacherName, updateTime, fullSchedule)
}

func (r *TeacherScheduleRepository) UpdateSchedule(teacherName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText) {
	query := fmt.Sprintf("UPDATE %s SET first_week_update_time=$2, second_week_update_time=$3, full_schedule=$4 WHERE teacher_name=$1", teacherScheduleRepoName)
	r.store.db.MustExec(query, teacherName, firstWeekUpdateTime, secondWeekUpdateTime, fullSchedule)
}

func (r *TeacherScheduleRepository) Information(teacherName string, firstWeekUpdateTime, secondWeekUpdateTime time.Time, fullSchedule types.JSONText) error {
	schedule, err := r.GetSchedule(teacherName)
	if err != nil {
		return err
	}

	if schedule != nil {
		r.UpdateSchedule(teacherName, firstWeekUpdateTime, secondWeekUpdateTime, fullSchedule)
	} else {
		r.AddSchedule(teacherName, fullSchedule)
	}
	return nil
}
