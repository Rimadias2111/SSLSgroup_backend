package api

import (
	"backend/api/controllers"
	_ "backend/docs" //for swagger
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Construct(cont controllers.Controller) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	r.Static("/images", "./public/images")

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	api := r.Group("/v1")
	{
		//Auth endpoints
		api.POST("/login", cont.Login)

		// Company endpoints
		api.POST("/companies", cont.CreateCompany)
		api.PUT("/companies/:company_id", cont.UpdateCompany)
		api.DELETE("/companies/:company_id", cont.DeleteCompany)
		api.GET("/companies/:company_id", cont.GetCompany)
		api.GET("/companies", cont.GetAllCompanies)

		// Driver endpoints
		api.POST("/drivers", cont.CreateDriver)
		api.PUT("/drivers/:driver_id", cont.UpdateDriver)
		api.DELETE("/drivers/:driver_id", cont.DeleteDriver)
		api.GET("/drivers/:driver_id", cont.GetDriver)
		api.GET("/drivers", cont.GetAllDrivers)

		// Employee endpoints
		api.POST("/employees", cont.CreateEmployee)
		api.PUT("/employees/:employee_id", cont.UpdateEmployee)
		api.DELETE("/employees/:employee_id", cont.DeleteEmployee)
		api.GET("/employees/:employee_id", cont.GetEmployee)
		api.GET("/employees", cont.GetAllEmployees)

		// Logistic endpoints
		api.POST("/logistics", cont.CreateLogistic)
		api.PUT("/logistics/:logistic_id", cont.UpdateLogistic)
		api.DELETE("/logistics/:logistic_id", cont.DeleteLogistic)
		api.GET("/logistics/:logistic_id", cont.GetLogistic)
		api.GET("/logistics", cont.GetAllLogistics)
		api.PUT("/logistics_with_cargo/:logistic_id", cont.UpdateLogisticCargo)
		api.POST("/terminate_logistics", cont.TerminateLogistic)
		api.POST("/cancel_late_logistics", cont.CancelLateLogistic)

		// Transaction endpoints
		api.POST("/transactions", cont.CreateTransaction)
		api.PUT("/transactions/:transaction_id", cont.UpdateTransaction)
		api.DELETE("/transactions/:transaction_id", cont.DeleteTransaction)
		api.GET("/transactions/:transaction_id", cont.GetTransaction)
		api.GET("/transactions", cont.GetAllTransactions)

		// Performance endpoints
		api.POST("/performances", cont.CreatePerformance)
		api.PUT("/performances/:performance_id", cont.UpdatePerformance)
		api.DELETE("/performances/:performance_id", cont.DeletePerformance)
		api.GET("/performances/:performance_id", cont.GetPerformance)
		api.GET("/performances", cont.GetAllPerformances)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
