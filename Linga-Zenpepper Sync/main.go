package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

func MgoInitialize() (db *mgo.Database) {

	//TODO: change in production
	session, conErr := mgo.Dial("localhost:27017")

	//dialInfo := &mgo.DialInfo{
	//	Addrs:    []string{rs1, rs2, rs3},
	//	Username: dbUsername,
	//	Password: dbPassword,
	//	DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
	//		return tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
	//	},
	//}
	//session, conErr := mgo.DialWithInfo(dialInfo)

	if conErr != nil {
		log.Println("DB Connection Error: ", conErr.Error())
	}

	session.SetMode(mgo.Monotonic, true)

	db = session.DB("Linga")

	return db
}

type RooCategory struct {
	C_ExternalID      bson.ObjectId   `json:"cExternalID" bson:"_id"`
	Name              string          `json:"name" bson:"name"`
	Level             int8            `json:"level" bson:"level"`
	SubCategories     []bson.ObjectId `json:"subCategories" bson:"-"`
	ImageAvailable    bool            `json:"imageAvailable" bson:"imageAvailable"`
	ExternalStore     bson.ObjectId   `json:"externalStore" bson:"createdFor"`
	ActiveStatus      bool            `json:"activeStatus" bson:"activeStatus"`
	RooTimeApplicable `json:"timeApplicable" bson:",inline"`
}

//type RooMenuItem struct {
//	MI_ExternalID           bson.ObjectId            `json:"miExternalID" bson:"_id"`
//	Name                    string                   `json:"name" bson:"name"`
//	Category                bson.ObjectId            `json:"category" bson:"category"`
//	ImageAvailable          bool                     `json:"imageAvailable" bson:"imageAvailable"`
//	CanSlice                bool                     `json:"canSlice" bson:"cutAndModify"`
//	NoOfSlice               int8                     `json:"noOfSlice" bson:"noOfSlice"`
//	Taxes                   []bson.ObjectId          `json:"taxes" bson:"-"`
//	InclusiveTaxes          []bson.ObjectId          `json:"inclusiveTaxes" bson:"-"`
//	ExternalStore           bson.ObjectId            `json:"externalStore" bson:"createdFor"`
//	ActiveStatus            bool                     `json:"activeStatus" bson:"activeStatus"`
//	CustomConfigs           []ServingSizeConfig      `json:"customConfigs" bson:"servingSizePrices"`
//	OptionalModifierGroups  []MandatoryModifierGroup `bson:"optionalModifierGroups" json:"-"`
//	IncludedModifiers       []IncludedModifier       `bson:"includedModifiers" json:"-"`
//	MandatoryModifierGroups []MandatoryModifierGroup `bson:"mandatoryModifierGroups" json:"-"`
//	RooTimeApplicable		RooTimeApplicable		 `json:"timeApplicable"`
//
//}

type RooTimeApplicable struct {
	TimeApplicableTypeStr string    `bson:"timeApplicableType" json:"-"`
	TimeApplicableType    int8      `json:"timeApplicableType"`
	DaysOfWeek            []int     `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	DaysOfMonth           []int     `bson:"daysOfMonth,omitempty" json:"daysOfMonth,omitempty"`
	Months                []int     `json:"months,omitempty" bson:"months,omitempty"`
	StartDate             time.Time `bson:"startDate,omitempty" json:"-"`
	StartDateStr          string    `bson:"-" json:"startDateStr,omitempty"`
	EndDate               time.Time `bson:"endDate,omitempty" json:"-"`
	EndDateStr            string    `bson:"-" json:"endDateStr,omitempty"`
	SpecialDate           time.Time `bson:"date,omitempty" json:"-"`
	SpecialDateStr        string    `bson:"-" json:"SpecialDateStr,omitempty"`
	IsRestrictToTime      bool      `bson:"isRestrictToTime,omitempty" json:"isRestrictToTime,omitempty"`
	IsRestrictToDays      bool      `json:"isRestrictToDays,omitempty" bson:"isRestrictToDays,omitempty"`
	IsRestrictToMonths    bool      `json:"isRestrictToMonths,omitempty" bson:"isRestrictToMonths,omitempty"`
	StartTime             string    `json:"startTimeStr,omitempty" bson:"startTime,omitempty"`
	EndTime               string    `json:"endTimeStr,omitempty" bson:"endTime,omitempty"`
}

func DateToString(date time.Time) string {
	return date.Format("02-Jan-2006")
}

var TimeApplicableTypes = []string{"Always", "Days of Week", "Days of Month", "Date Range",
	"Specific date", "Start date time & end date time"}

func (timeApplicable *RooTimeApplicable) DateToString() {

	for index, val := range TimeApplicableTypes {
		if timeApplicable.TimeApplicableTypeStr == val {
			timeApplicable.TimeApplicableType = int8(index)
		}

	}

	if timeApplicable.TimeApplicableTypeStr != "Always" {
		if !timeApplicable.StartDate.IsZero() {
			timeApplicable.StartDateStr = DateToString(timeApplicable.StartDate)
		}

		if !timeApplicable.EndDate.IsZero() {
			timeApplicable.EndDateStr = DateToString(timeApplicable.EndDate)
		}

		if !timeApplicable.SpecialDate.IsZero() {
			timeApplicable.SpecialDateStr = DateToString(timeApplicable.SpecialDate)
		}
	}
}

func formRooCategories(storeId bson.ObjectId, w http.ResponseWriter) {
	//category := RooSendCategory{}
	log.Println("storeId", storeId)
	categories := RooCategory{}
	MgoInitialize().C("category").Find(bson.M{"createdFor": storeId}).One(&categories)
	//log.Println("category", categories)
	categories.DateToString()
	//res :=categories.makeJson()
	res := categories.makeJsonFromStruct()
	RenderJSON(w, http.StatusOK, res)
}

func (c RooCategory) makeJson() map[string]interface{} {
	category := make(map[string]interface{})
	timeApplicable := make(map[string]interface{})
	timeApplicable["timeApplicableType"] = c.TimeApplicableType
	timeApplicable["daysOfWeek"] = c.DaysOfWeek
	timeApplicable["daysOfMonth"] = c.DaysOfMonth
	timeApplicable["startDateStr"] = c.StartDateStr
	timeApplicable["endDateStr"] = c.EndDateStr
	timeApplicable["SpecialDateStr"] = c.SpecialDateStr
	timeApplicable["isRestrictToTime"] = c.IsRestrictToTime
	timeApplicable["isRestrictToDays"] = c.IsRestrictToDays
	timeApplicable["isRestrictToMonths"] = c.IsRestrictToMonths
	timeApplicable["startTimeStr"] = c.StartTime
	timeApplicable["endTimeStr"] = c.EndTime

	category["cExternalID"] = c.C_ExternalID
	category["name"] = c.Name
	category["level"] = c.Level
	category["subCategories"] = c.SubCategories
	category["imageAvailable"] = c.ImageAvailable
	category["externalStore"] = c.ExternalStore
	category["activeStatus"] = c.ActiveStatus
	category["timeApplicable"] = timeApplicable

	return category
}

func (c RooCategory) makeJsonFromStruct() RooCategory {
	//log.Println(c

	fmt.Printf("%+v", c)

	//log.Println(c.RooTimeApplicable)
	//s := fmt.Sprintf("%v", c.RooTimeApplicable)
	//k := byte(s)
	//
	//timeApplicable := make(map[string]interface{})
	//timeApplicable["timeApplicableType"] = c.TimeApplicableType
	//timeApplicable["daysOfWeek"] = c.DaysOfWeek
	//timeApplicable["daysOfMonth"] = c.DaysOfMonth
	//timeApplicable["startDateStr"] = c.StartDateStr
	//timeApplicable["endDateStr"] = c.EndDateStr
	//timeApplicable["SpecialDateStr"] = c.SpecialDateStr
	//timeApplicable["isRestrictToTime"] = c.IsRestrictToTime
	//timeApplicable["isRestrictToDays"] = c.IsRestrictToDays
	//timeApplicable["isRestrictToMonths"] = c.IsRestrictToMonths
	//timeApplicable["startTimeStr"] = c.StartTime
	//timeApplicable["endTimeStr"] = c.EndTime
	//
	//c.time

	//log.Println(h)
	return c
}

func main() {
	fmt.Println(time.Now())
	http.ListenAndServe(":9999", http.HandlerFunc(CategoryHandler))
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	formRooCategories(bson.ObjectIdHex("57357cf1b46e16e0b9000104"), w)

	//RenderJSON(w,http.StatusOK,category)

}

func RenderJSON(w http.ResponseWriter, status int, res interface{}) {
	resByte, _ := json.MarshalIndent(res, "", "	")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(resByte)
}
