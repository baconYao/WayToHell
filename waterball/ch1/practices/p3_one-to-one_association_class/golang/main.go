package main

import (
	"fmt"
	"one-to-one-assoociation-class/lecture"
)

func main() {
	s1 := lecture.NewStudent("bacon")
	lecture := lecture.NewLecture("Leetcode 100 天大集合")

	err := lecture.SignUp(s1)
	if err != nil {
		fmt.Println(err)
	}
	err = lecture.SignUp(s1)
	if err != nil {
		fmt.Println("Error: 不能重複註冊")
	}

	attendance := lecture.GetLectureAttendance()
	if attendance == nil {
		fmt.Printf("Error: 課程 '%s' 沒有任何學生參與", lecture.GetName())
	}
	attendance.ReceiveGrade(60)

	lecture.SignOff(s1)
	err = lecture.SignOff(s1)
	if err != nil {
		fmt.Println("Error: 學生已經取消上課，無法再次取消")
	}
}
