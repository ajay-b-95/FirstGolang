package handlers

import (
	"Interview/cab"
	"Interview/location"
	"Interview/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type GCab struct {
	ID          int                `json:"ID"`
	IsAvailable bool               `json:"isAvailable"`
	Location    *location.Location `json:"PicklocA"`
	Time        time.Time          `json:"time"`
}

type bookinglist struct {
	Booked []GCab
}

//Cabdetails to be Exported
func Cabdetails(c echo.Context) error {
	var Cabs model.Listresponsemodelstrct
	Cabs.Ocab = cab.GetAllCabs()
	var cablists []byte
	cablists, err := json.MarshalIndent(Cabs, "", " ")

	if err != nil {
		log.Printf("Not working %s", err)
		return c.String(http.StatusInternalServerError, "Something Went Wrong")
	}

	return c.JSONBlob(http.StatusOK, cablists)
}

//func Getacab(c echo.Context) error {

//	gcab := GCab{}

//	defer c.Request().Body.Close()
//	err := json.NewDecoder(c.Request().Body).Decode(&gcab)

//	if err != nil {
//		log.Printf("Failed Processing %s", err)
//		return echo.NewHTTPError(http.StatusInternalServerError)

//	}

//	log.Printf("Cab is added")
//	return c.String(http.StatusOK, "Cab is Booked!")

//}

//Getacab to be Exported
func Getacab(c echo.Context) error {
	j := bookinglist{}
	gca := GCab{}

	//pla := model.Listresponsemodelstrct{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to read the request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//	errr := c.Bind(&gca)
	//	if errr != nil {
	//		log.Printf("Fialed")
	//		return c.String(http.StatusInternalServerError, "Field msissing")
	//	}

	err = json.Unmarshal(b, &gca)
	if err != nil || gca.ID < 0 || gca.ID > 5 {
		log.Printf("Failed to unmarhal %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if gca.IsAvailable == false {
		return c.String(http.StatusInternalServerError, "Cab is not available")
	}

	j.AddtomyBooking(gca)
	SetAvailability(gca.ID, false)
	return c.String(http.StatusOK, "Your Cab is booked !!")
}

//SetAvailability to be Exported
func SetAvailability(cabID int, status bool) {
	for index, val := range cab.Cablist {
		if val.ID == cabID {
			cab.Cablist[index].IsAvailable = status
		}
	}
}

func (book *bookinglist) AddtomyBooking(booked GCab) {
	book.Booked = append(book.Booked, booked)
}

//Getallbooking returns all the booking
func Getallbooking() []GCab {
	var h bookinglist
	return h.Booked
}

//Booking is the handler func to url path
func Booking(c echo.Context) error {
	var u []GCab
	u = Getallbooking()
	var booklist []byte
	booklist, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Printf("Not working %s", err)
		return c.String(http.StatusInternalServerError, "Something Went Wrong")
	}
	log.Printf("Reached")
	return c.JSONBlob(http.StatusOK, booklist)
}
