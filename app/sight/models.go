package sight

import (
	"api_travelguide_v2/app/common"
	"github.com/globalsign/mgo/bson"
)

const sightDocument = "sight"

type Sight struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Tags        []string      `json:"tags" bson:"tags"`
	Cost        float32       `json:"cost" bson:"cost"`
	Coordinate  [2]int        `json:"coordinate" bson:"coordinate"`
	Rating      float32       `json:"rating" bson:"rating"`
	TypeSight   string        `json:"type_sight" bson:"type_sight"`
	Urls        []string      `json:"urls" bson:"urls"`
	WebSite     string        `json:"web_site" bson:"web_site"`
	Description string        `json:"description" bson:"description"`
	History     string        `json:"history" bson:"history"`
	PhoneNumber string        `json:"phone_number" bson:"phone_number"`
	TownID      bson.ObjectId `json:"town_id" bson:"_townID"`
}

func SaveTown(data interface{}) error {
	db := common.GetDb()
	err := db.C(sightDocument).Insert(data)
	return err
}

func GetSightsByTownID(townId *bson.ObjectId) ([]Sight, error) {
	db := common.GetDb()
	var res []Sight
	err := db.C(sightDocument).Find(bson.M{"_townID": townId}).All(&res)
	return res, err
}
