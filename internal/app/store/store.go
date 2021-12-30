package store

// ScheduleStore ...
type ScheduleStore interface {
	// GroupSchedule ...
	GroupSchedule() GroupScheduleRepository

	// TeacherSchedule ...
	TeacherSchedule() TeacherScheduleRepository
}
