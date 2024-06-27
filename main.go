package main

import (
	"awesomeProject3/List"
	"awesomeProject3/group"
	"awesomeProject3/student"

	//"awesomeProject3/student"
	"awesomeProject3/university"
	//"errors"
	"fmt"
)

func main() {
	Spisok := List.NewList()
	universityMenu(Spisok)
}

func universityMenu(s *List.List) {
	flag := -1
	for flag != 0 {
		fmt.Println("0) Close\n1) Show all universities\n2) Create university\n3) Delete university\n4) Choose university")
		fmt.Println("Please choose action")

		fmt.Scan(&flag)

		switch flag {
		case 1:
			showAllUni(s)
		case 2:
			createUni(s)
		case 3:
			deleteUni(s)
		case 4:
			chooseUni(s)
		}
	}

}

func univerPersonalMenu(uni *List.Node) {
	flag := -1
	for flag != 0 {
		fmt.Println("0) Close\n1) Show all groups\n2) Create new group\n3) Delete group\n4) Choose group")
		fmt.Println("Please choose action")

		fmt.Scan(&flag)

		switch flag {
		case 1:
			showAllGroups(uni.Value)
		case 2:
			createGroup(uni.Value)
		case 3:
			fmt.Println("3) Delete group")
		case 4:
			chooseGroup(uni.Value)
		}
	}
}

func groupPersonalMenu(grp *group.Group) {
	flag := -1
	for flag != 0 {
		fmt.Println("0) Close\n1) Show all students\n2) Choose student\n3) Add student\n4) Delete student")
		fmt.Println("Please choose action")

		fmt.Scan(&flag)

		switch flag {
		case 1:
			showAllStuds(grp)
		case 2:
			chooseStud(grp)
		case 3:
			addStud(grp)
		case 4:
			fmt.Println("4) Delete student")
		}
	}
}

func studPersonalMenu(stud *student.Student) {
	flag := -1
	for flag != 0 {
		fmt.Println("0) Close\n1) Info\n2) Rename\n3) Add mark")
		fmt.Println("Please choose action")

		fmt.Scan(&flag)

		switch flag {
		case 1:
			studInfo(stud)
		case 2:
			studRename(stud)
		case 3:
			stud.AddMark()
		}
	}
}

// Функции для разных меню

func showAllUni(s *List.List) {
	temp := s.Head
	i := 1
	for temp != nil {
		fmt.Print(i, ".")
		uniInfo(temp.Value)
		temp = temp.Next
		i++
	}
	if s.Head == nil {
		fmt.Println("There is no university in it")
	}
}

func uniInfo(u *university.University) {
	fmt.Println(u.Name)
}

func createUni(spsk *List.List) {
	var uniName string
	fmt.Println("Add new university name")
	fmt.Scan(&uniName)
	newUni := university.NewUniversity(uniName)
	spsk.Push(newUni)
}

func deleteUni(spsk *List.List) {
	var numb int

	fmt.Println("Choose number of university to delete")
	showAllUni(spsk)
	fmt.Scan(&numb)

	temp := spsk.Head
	tempNumb := 1
	for tempNumb != numb {
		temp = temp.Next
		tempNumb++
	}

	temp.Value, _ = spsk.Delete(temp)
}

func chooseUni(spsk *List.List) {
	var numb int

	fmt.Println("Choose number of university")
	showAllUni(spsk)
	fmt.Scan(&numb)

	temp := spsk.Head
	tempNumb := 1
	for tempNumb != numb {
		temp = temp.Next
		tempNumb++
	}
	univerPersonalMenu(temp)
}

func showAllGroups(uni *university.University) {
	fmt.Println(uni.Name, "consist of next groups:")
	for num, grp := range uni.Groups {
		fmt.Println(num+1, grp.Number)
	}
}

func createGroup(uni *university.University) {
	var grpName int
	fmt.Println("Add new group number")
	fmt.Scan(&grpName)
	newGrp := group.NewGroup(grpName)
	uni.Groups = append(uni.Groups, *newGrp)
}

/*func deleteGroup(uni *university.University) {
	var numb int

	fmt.Println("Choose number of group to delete")
	showAllGroups(uni)
	fmt.Scan(&numb)

	copy(uni.Groups[numb-1:], uni.Groups[numb:])
	uni.Groups[len(uni.Groups)-1] = //нулевое значение
	uni.Groups = uni.Groups[len(uni.Groups)-1]
}*/

func chooseGroup(uni *university.University) {
	var numb int

	fmt.Println("Choose number of group")
	showAllGroups(uni)
	fmt.Scan(&numb)

	grp := uni.Groups[numb+1]
	groupPersonalMenu(&grp)
}

func showAllStuds(grp *group.Group) {
	fmt.Println(grp.Number, "consist of next students:")
	for num, stud := range grp.GrpStudents {
		fmt.Println(num+1, stud.Name)
	}
}

func addStud(grp *group.Group) {
	var name, nick string
	var age int
	fmt.Println("Please, enter some info/nName:")
	fmt.Scan(&name)
	fmt.Println("Nickname:")
	fmt.Scan(&nick)
	fmt.Println("Age:")
	fmt.Scan(&age)
	newStud := student.NewStudent(name, nick, age, grp.Number)
	grp.GrpStudents = append(grp.GrpStudents, *newStud)
}

func chooseStud(grp *group.Group) {
	var numb int

	fmt.Println("Choose number of student")
	showAllStuds(grp)
	fmt.Scan(&numb)

	stud := grp.GrpStudents[numb+1]
	studPersonalMenu(&stud)
}

/*func deleteStud(grp *group.Group {

}*/

func studInfo(std *student.Student) {
	fmt.Println("Student", std.Name, std.Nickname)
	fmt.Println("Age:", std.Age)
	fmt.Println("Group:", std.GrpNumber)
	fmt.Print("Marks: ")
	for i := 0; i < len(std.Marks); i++ {
		fmt.Print(std.Marks[i], "")
	}
	fmt.Print()
	fmt.Println("Average mark:", std.AvrgMark)
}

func studRename(stud *student.Student) {
	var name, nick string
	fmt.Println("Enter new name")
	fmt.Scan(&name)
	fmt.Println("Enter new nickname")
	fmt.Scan(&nick)
	stud.Name = name
	stud.Nickname = nick
}
