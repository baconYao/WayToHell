package lecture

import (
	"fmt"
	"one-to-one-assoociation-class/utils"
)

type LectureAttendance struct {
	grade   int
	student *Student
	lecture *Lecture
}

func NewLectureAttendance(student *Student, lecture *Lecture) *LectureAttendance {
	return &LectureAttendance{
		grade: 0, student: student, lecture: lecture,
	}
}

// ReceiveGrade assigns the grade value
func (la *LectureAttendance) ReceiveGrade(grade int) error {
	g, err := utils.ShouldBeWithinRange("Grade", grade, 0, 100)
	if err != nil {
		return err
	}
	la.grade = g
	fmt.Printf("在 '%s' 課中，學生 '%s' 拿了 '%d' 分\n", la.lecture.GetName(), la.student.getName(), grade)
	return nil
}

func (la LectureAttendance) getGrade() int {
	return la.grade
}

func (la LectureAttendance) getStudent() *Student {
	return la.student
}

func (la LectureAttendance) getLecture() *Lecture {
	return la.lecture
}
