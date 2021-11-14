package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Raven/models/targetModels"
	"net/http"
)

func (TargetRouters) GetTargetREPost(c *gin.Context) {
	targetModels.TargetInitDB()
	targetModels.GetTargetData()

	c.JSON(http.StatusOK, nil)
}
