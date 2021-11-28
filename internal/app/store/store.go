package store

// Store ...
type Store interface {
	GroupSchedule() GroupScheduleRepository
	TeacherSchedule() TeacherScheduleRepository
}
