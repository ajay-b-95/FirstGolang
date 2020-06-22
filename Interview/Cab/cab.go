package cab

import (
	"Interview/location"
	"time"
)

//Cab to be exported in Model to prevent cycle import
type Cab struct {
	ID          int                `json:"ID"`
	IsAvailable bool               `json:"isAvailable"`
	Location    *location.Location `json:"Pickloc"`
	Time        time.Time          `json:"time"`
}

var Cablist []Cab

func init() {

	Cablist = []Cab{
		{ID: 1,
			IsAvailable: true,
			Location:    &location.Location{100, 200},
			Time:        time.Now()},

		{ID: 2,
			IsAvailable: true,
			Location:    &location.Location{300, 400},
			Time:        time.Now()},

		{ID: 3,
			IsAvailable: true,
			Location:    &location.Location{400, 100},
			Time:        time.Now()},

		{ID: 4,
			IsAvailable: true,
			Location:    &location.Location{800, 700},
			Time:        time.Now()},

		{ID: 5,
			IsAvailable: true,
			Location:    &location.Location{900, 1300},
			Time:        time.Now()},
	}

}

//GetAllCabs to be exported
func GetAllCabs() []Cab {
	return Cablist
}
