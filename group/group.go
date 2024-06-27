package group

import (
	"awesomeProject3/student"
	"reflect"
)

type Group struct {
	Number      int
	GrpStudents []student.Student
	AvrgMark    float64
}

func NewGroup(number int) *Group {
	studSlice := make([]student.Student, 0, 10)
	var avrgGrpMrk float64

	return &Group{
		Number:      number,
		GrpStudents: studSlice,
		AvrgMark:    avrgGrpMrk,
	}
}

func (grp *Group) AddToGroup(stud *student.Student) {
	grp.GrpStudents = append(grp.GrpStudents, *stud)
}

func (grp *Group) AvrgGroupMark() {
	var avrgMrk float64
	var mrk reflect.Value

	for _, stdnt := range grp.GrpStudents {
		mrk = reflect.ValueOf(stdnt.AvrgMark)
		avrgMrk += mrk.Float()
	}
	grp.AvrgMark = avrgMrk / float64(len(grp.GrpStudents))
}

func (grp *Group) SortStudByAvrg() {
	len := len(grp.GrpStudents)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if grp.GrpStudents[j].AvrgMark > grp.GrpStudents[j+1].AvrgMark {
				grp.GrpStudents[j], grp.GrpStudents[j+1] = grp.GrpStudents[j+1], grp.GrpStudents[j]
			}
		}
	}
}
