package main

import (
	"./app/common"
	"./app/town"
	"./app/user"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.Init()
	defer db.Session.Close()

	r := gin.Default()

	v1 := r.Group("/api")

	user.UserRegister(v1.Group("/user"))
	town.TownRegister(v1.Group("/town"))
	//sight.SightRegister(v1.Group("/sight"))

	r.Run()
}
