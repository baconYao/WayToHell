package lecture

type Student struct {
	name              string
	lectureAttendance *LectureAttendance // 代表此學生修了這們課

}

func NewStudent(name string) *Student {
	return &Student{
		name: name,
	}
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) getLectureAttendance() *LectureAttendance {
	return s.lectureAttendance
}

func (s *Student) setLectureAttendance(lectureAttendance *LectureAttendance) {
	s.lectureAttendance = lectureAttendance
}
