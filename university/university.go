package university

import (
	"awesomeProject3/group"
	"reflect"
)

type University struct {
	Name     string
	Groups   []group.Group
	AvrgMark float64
}

func NewUniversity(name string) *University {
	grpSlice := make([]group.Group, 0, 10)
	var avrgMrk float64

	return &University{
		Name:     name,
		Groups:   grpSlice,
		AvrgMark: avrgMrk,
	}
}

func (uni *University) AddToUni(grp *group.Group) {
	uni.Groups = append(uni.Groups, *grp)
}

func (uni *University) UniAvrgMark() {
	var avrgMrk float64
	var mrk reflect.Value

	for _, grp := range uni.Groups {
		mrk = reflect.ValueOf(grp.AvrgMark)
		avrgMrk += mrk.Float()
	}
	uni.AvrgMark = avrgMrk / float64(len(uni.Groups))
}
