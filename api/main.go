package api

import (
	"backend/api/controllers"
	_ "backend/docs" //for swagger
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Construct(cont controllers.Controller) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./public/images")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	api := r.Group("/v1")
	{
		// Company endpoints
		api.POST("/companies", cont.CreateCompany)
		api.PUT("/companies/:company_id", cont.UpdateCompany)
		api.DELETE("/companies/:company_id", cont.DeleteCompany)
		api.GET("/companies/:company_id", cont.GetCompany)
		api.GET("/companies", cont.GetAllCompanies)

		//Driver endpoints
		api.POST("/drivers", cont.CreateDriver)
		api.PUT("/drivers/:driver_id", cont.UpdateDriver)
		api.DELETE("/drivers/:driver_id", cont.DeleteDriver)
		api.GET("/drivers/:driver_id", cont.GetDriver)
		api.GET("/drivers", cont.GetAllDrivers)

		api.POST("/employees", cont.CreateEmployee)
		api.PUT("/employees/:employee_id", cont.UpdateEmployee)
		api.DELETE("/employees/:employee_id", cont.DeleteEmployee)
		api.GET("/employees/:employee_id", cont.GetEmployee)
		api.GET("/employees", cont.GetAllEmployees)
	}

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
