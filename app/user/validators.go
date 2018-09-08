package user

import (
	"api_travelguide_v2/app/common"
	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	Username  string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	Password  string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	userModel User   `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.userModel.Username = self.Username
	self.userModel.SetPassword(self.Password)
	return nil
}

func NewUserModelValidator() UserModelValidator {
	UserModelValidator := UserModelValidator{}
	return UserModelValidator
}
