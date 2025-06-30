package lecture

import "fmt"

type Lecture struct {
	name              string
	lectureAttendance *LectureAttendance // 代表哪一個學生參與了這們課
}

func NewLecture(name string) *Lecture {
	return &Lecture{name: name}
}

func (l Lecture) GetName() string {
	return l.name
}

func (l Lecture) GetLectureAttendance() *LectureAttendance {
	return l.lectureAttendance
}

func (l *Lecture) SignUp(student *Student) error {
	if student.getLectureAttendance() != nil {
		return fmt.Errorf("student has taken a lecture")
	}

	if l.lectureAttendance != nil {
		return fmt.Errorf("can only containe one student")
	}

	l.lectureAttendance = NewLectureAttendance(student, l)
	student.setLectureAttendance(l.lectureAttendance)

	return nil
}

func (l *Lecture) SignOff(student *Student) error {
	if l.lectureAttendance == nil || l.lectureAttendance.getStudent() != student {
		return fmt.Errorf("the student didn't sign up")
	}

	l.lectureAttendance = nil
	student.setLectureAttendance(nil)
	return nil
}
