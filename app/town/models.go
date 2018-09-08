package town

import (
	"api_travelguide_v2/app/common"
	"github.com/globalsign/mgo/bson"
)

const townDocument = "town"

type Town struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	UrlPhoto string        `json:"url_photo" bson:"url_photo"`

	CountryID bson.ObjectId `json:"country_id" bson:"_countryID"`
}

func GetAllTowns() ([]Town, error) {
	db := common.GetDb()
	var res []Town
	err := db.C(townDocument).Find(nil).All(&res)
	return res, err
}

func SaveTown(data interface{}) error {
	db := common.GetDb()
	err := db.C(townDocument).Insert(data)
	return err
}
