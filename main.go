package main

import (
	"awesomeProject3/group"
	"awesomeProject3/student"
	"fmt"
)

func main() {
	stud := student.NewStudent("Pete", "Student", 18, 11)
	stud.Name = "Pete"
	grp := group.NewGroup(1203)
	fmt.Print(stud)
	stud.AddMark()
	stud.AddMark()
	stud.Calculate()
	grp.AddToGroup(stud)
}
