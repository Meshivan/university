package group

type Group struct {
	Number      int
	GrpStudents []string
}

func NewGroup(number int) *Group {
	studSlice := make([]string, 0, 10)

	return &Group{
		Number:      number,
		GrpStudents: studSlice,
	}
}

func (grp Group) AddToGroup(stud Student) {
	append(grp.GrpStudents, stud)
}
