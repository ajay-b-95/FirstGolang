package model

import "Interview/cab"

//Bookingmodel to be exported
type Bookingmodel struct {
	Cab   cab.Cab `json:"Cab"`
	Error string  `json:"Errors"`
}

type Listresponsemodelstrct struct {
	Ocab []cab.Cab `json:"Ocab"`
	//Error string    `json:"Errors"`
}
