package location

import "fmt"

//to be imported
type Location struct {
	Pickloca int64 `json:"Pickloca"`
	Picklocb int64 `json:"Picklocb"`
}

func (l *Location) displaylocation() {
	fmt.Println("Loacation %d to Location %d", l.Pickloca, l.Picklocb)
}

//to be exported
func Getpickloca(l *Location) int64 {
	return l.Pickloca
}

//to be exported
func Getpicklocb(l *Location) int64 {
	return l.Picklocb
}
