package sight

import (
	"api_travelguide_v2/app/common"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type SightModelValidator struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name" binding:"required"`
	Tags       []string      `json:"tags" bson:"tags" binding:"len=2,dive,max=10"`
	Cost       float32       `json:"cost" bson:"cost"`
	Coordinate [2]int        `json:"coordinate" bson:"coordinate" binding:"len=2"`
	//Rating      float32       `json:"rating" bson:"rating""`
	TypeSight   string        `json:"type_sight" bson:"type_sight"`
	Urls        []string      `json:"urls" bson:"urls" binding:"max=100,dive,max=511"`
	WebSite     string        `json:"web_site" bson:"web_site" binding:"max=511"`
	Description string        `json:"description" bson:"description" binding:"max=511"`
	History     string        `json:"history" bson:"history" binding:"max=32767"`
	PhoneNumber string        `json:"phone_number" bson:"phone_number" binding:"max=12"`
	TownID      bson.ObjectId `json:"town_id" bson:"_townID" binding:"required"`

	sightModel Sight `json:"-"`
}

type RequestSightGet struct {
	TownID bson.ObjectId `json:"town_id" binding:"requ"`
}

func (self *SightModelValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.sightModel.Name = self.Name
	self.sightModel.Tags = self.Tags
	self.sightModel.Cost = self.Cost
	self.sightModel.Coordinate = self.Coordinate
	//self.sightModel.Rating = self.Rating
	self.sightModel.TypeSight = self.TypeSight
	self.sightModel.Urls = self.Urls
	self.sightModel.WebSite = self.WebSite
	self.sightModel.Description = self.Description
	self.sightModel.History = self.History
	self.sightModel.PhoneNumber = self.PhoneNumber
	self.sightModel.TownID = self.TownID
	return nil
}

func NewSightModelValidator() SightModelValidator {
	sightModelValidator := SightModelValidator{}
	return sightModelValidator
}

type RequestGetSightValidator struct {
	TownID bson.ObjectId `json:"town_id" bson:"_id"`
}

func (self *RequestGetSightValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}
	return nil
}

func NewRequestGetSight() RequestGetSightValidator {
	requestGetSight := RequestGetSightValidator{}
	return requestGetSight
}
