package group

import (
	"awesomeProject3/student"
)

type Group struct {
	Number      int
	GrpStudents []student.Student
}

func NewGroup(number int) *Group {
	studSlice := make([]student.Student, 0, 10)

	return &Group{
		Number:      number,
		GrpStudents: studSlice,
	}
}

func (grp Group) AddToGroup(stud student.Student) {
	append(grp.GrpStudents, stud)
}
