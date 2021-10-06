package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swirling-melodies/Raven/controllers"
	"github.com/swirling-melodies/Raven/middlewares"

	_ "github.com/swirling-melodies/Raven/docs" // docs are generated by Swag CLI, you have to import it.
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
	workGroupV1(v1)
	logGroupV1(v1)
	userGroupV1(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func testGroup(c *gin.RouterGroup) {
	c.GET("/testInsert", controllers.TestHome)
}

func billGroupV1(c *gin.RouterGroup) {
	c.POST("/Bill/GetBillsYearAllData", controllers.GetBillsYearAllData)
	c.POST("/Bill/GetBillsDataByMonth", controllers.GetBillsDataByMonth)
	c.POST("/Bill/GetBillsTable", controllers.GetBillsTable)
	c.POST("/Bill/GetBillsOption", controllers.GetBillsTableOption)
	c.POST("/Bill/GetBillsDiagram", controllers.GetBillsDiagram)
	c.POST("/Bill/GetBillsDataByPage", controllers.GetBillsDataByPage)
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

func workGroupV1(c *gin.RouterGroup) {
	c.POST("/Work/BillNameSetWork", controllers.BillNameSetWork)
	c.POST("/Work/GetBillNameList", controllers.GetBillNameList)
	c.POST("/Work/UpdateBillName", controllers.UpdateBillName)
	c.POST("/Work/UserSetWork", controllers.UserSetWork)
	c.POST("/Work/InvestmentItemSetWork", controllers.InvestmentItemSetWork)
}

func logGroupV1(c *gin.RouterGroup) {
	c.POST("/Log/GetLogTable", controllers.GetLogTable)
}

func userGroupV1(c *gin.RouterGroup) {
	c.POST("/User/Login", controllers.Login)
	c.POST("/User/ValidateToken", controllers.ValidateToken)
}
