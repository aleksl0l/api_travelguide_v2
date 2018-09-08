package sight

import (
	"api_travelguide_v2/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SightRegister(router *gin.RouterGroup) {
	router.POST("/", CreateSight)
	router.GET("/", GetSight)
	router.PUT("/", ModifySight)
}

func CreateSight(c *gin.Context) {
	sightModelValidator := NewSightModelValidator()
	if err := sightModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	sight := sightModelValidator.sightModel

	err := SaveTown(sight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't save that object"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetSight(c *gin.Context) {
	getSightValidator := NewRequestGetSight()
	if err := getSightValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	sights, err := GetSightsByTownID(&getSightValidator.TownID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't retrieve that object"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sights})
}

func ModifySight(c *gin.Context) {

}
