package router

import (
	"Raven/src/web/controllers"
	"Raven/src/web/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.CrossDomain())

	test := router.Group("test")
	testGroup(test)

	v1 := router.Group("v1")
	billGroupV1(v1)
	desireGroupV1(v1)
	investmentGroupV1(v1)
	targetGroupV1(v1)

	router.Run()
}

func testGroup(c *gin.RouterGroup) {
	c.GET("/testInsert", controllers.TestHome)
}

func billGroupV1(c *gin.RouterGroup) {
	c.POST("/Bill/GetYearAllData", controllers.GetYearAllData)
	c.POST("/Bill/GetFourMonthAllData", controllers.GetFourMonthData)
}

func investmentGroupV1(c *gin.RouterGroup) {
	c.GET("/Investment/test", controllers.GetInvestments)
	c.POST("/Investment/GetInvestments", controllers.GetInvestments)
	c.POST("/Investment/GetInvestmentsTable", controllers.GetInvestmentsTable)
	c.POST("/Investment/AddInvestmentsTable", controllers.AddInvestmentsTable)
	c.POST("/Investment/UpdateInvestmentsTable", controllers.UpdateInvestmentsTable)
	c.POST("/Investment/GetInvestmentDiagram", controllers.GetInvestmentDiagram)
	c.POST("/Investment/GetInvestmentOption", controllers.GetInvestmentOption)
}

func desireGroupV1(c *gin.RouterGroup) {
	c.POST("/Desire/GetDesire", controllers.GetDesire)
}

func targetGroupV1(c *gin.RouterGroup) {
	c.POST("/Target/GetTarget", controllers.GetTarget)
}