package store

// ScheduleStore represents a database with backups of schedules.
type ScheduleStore interface {
	// GroupSchedule represents a database table with backups of group schedules.
	GroupSchedule() *GroupScheduleRepository

	// TeacherSchedule represents a database table with backups of teacher schedules.
	TeacherSchedule() *TeacherScheduleRepository
}
