package student

import "fmt"

type Student struct {
	Name      string
	Age       int
	Nickname  string
	GrpNumber int
	Marks     []int
	AvrgMark  float64
}

func NewStudent(name, nickname string, age, grpNumber int) *Student {
	newSlice := make([]int, 0, 5)
	var avrgRslt float64

	return &Student{
		Name:      name,
		Nickname:  nickname,
		Age:       age,
		GrpNumber: grpNumber,
		Marks:     newSlice,
		AvrgMark:  avrgRslt,
	}
}

func (stud *Student) AddMark() {
	mark := 0
	fmt.Println("Напишите добавляемую оценку для студента", stud.Name)
	fmt.Scan(&mark)
	stud.marks = append(stud.marks, mark)
	stud.Calculate()
}

func (stud *Student) Calculate() {
	numb := 0
	for _, val := range stud.marks {
		numb += val
	}

	stud.AvrgMark = float64(numb) / float64(len(stud.marks))
}
