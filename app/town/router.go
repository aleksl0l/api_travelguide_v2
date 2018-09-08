package town

import (
	"api_travelguide_v2/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TownRegister(router *gin.RouterGroup) {
	router.POST("/", CreateTown)
	router.GET("/", GetTowns)
}

func CreateTown(c *gin.Context) {
	townModelValidator := NewTownModelValidator()
	if err := townModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	todo := townModelValidator.townModel

	err := SaveTown(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't save that object"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetTowns(c *gin.Context) {
	towns, err := GetAllTowns()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": towns})
}
