package student

import "fmt"

type Student struct {
	Name      string
	Age       int
	Nickname  string
	GrpNumber int
	marks     []int
}

func NewStudent(name, nickname string, age, grpNumber int) *Student {
	newSlice := make([]int, 0, 5)

	return &Student{
		Name:      name,
		Nickname:  nickname,
		Age:       age,
		GrpNumber: grpNumber,
		marks:     newSlice,
	}
}

func (stud Student) AddMark() {
	mark := 0
	fmt.Println("Напишите добавляемую оценку")
	fmt.Scan(&mark)
	stud.marks = append(stud.marks, mark)
}

func (stud Student) Calculate() {
	numb := 0
	for _, val := range stud.marks {
		numb += val
	}
	result := numb / len(stud.marks)

	fmt.Println(result)
	fmt.Print("hi")
}
