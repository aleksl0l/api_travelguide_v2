package town

import (
	"api_travelguide_v2/app/common"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type TownModelValidator struct {
	Name      string        `form:"name" json:"name" binding:"exists,alphanum,max=255"`
	UrlPhoto  string        `form:"url_photo" json:"url_photo" binding:"exists,max=1023"`
	CountryID bson.ObjectId `form:"countryID" json:"countryID" binding:"exists"`
	townModel Town          `json:"-"`
}

func (self *TownModelValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.townModel.Name = self.Name
	self.townModel.UrlPhoto = self.UrlPhoto
	self.townModel.CountryID = self.CountryID

	return nil
}

func NewTownModelValidator() TownModelValidator {
	TownModelValidator := TownModelValidator{}
	return TownModelValidator
}
