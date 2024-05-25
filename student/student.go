package student

import "fmt"

type Student struct {
	Name      string
	Age       int
	Nickname  string
	GrpNumber int
	marks     []int
	AvrgMark  float32
}

func NewStudent(name, nickname string, age, grpNumber int) *Student {
	newSlice := make([]int, 0, 5)
	avrgRslt := 0

	return &Student{
		Name:      name,
		Nickname:  nickname,
		Age:       age,
		GrpNumber: grpNumber,
		marks:     newSlice,
		AvrgMark:  avrgRslt,
	}
}

func (stud Student) AddMark() {
	mark := 0
	fmt.Println("Напишите добавляемую оценку")
	fmt.Scan(&mark)
	stud.marks = append(stud.marks, mark)
	stud.Calculate()
}

func (stud Student) Calculate() float32 {
	var avg float32
	numb := 0
	for _, val := range stud.marks {
		numb += val
	}

	stud.AvrgMark = (float32)numb / len(stud.marks)
}
