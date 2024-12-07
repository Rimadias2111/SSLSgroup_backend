package api

import (
	"backend/api/controllers"
	"backend/api/middleware"
	_ "backend/docs" //for swagger
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
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

	// To start again
	r.Static("/images", "./public/images")

	api := r.Group("/v1")
	{
		r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

		//Auth endpoints
		api.POST("/login", cont.Login)

		// Company endpoints
		api.POST("/companies", middleware.AuthMiddleware(2), cont.CreateCompany)
		api.PUT("/companies/:company_id", middleware.AuthMiddleware(2), cont.UpdateCompany)
		api.DELETE("/companies/:company_id", middleware.AuthMiddleware(2), cont.DeleteCompany)
		api.GET("/companies/:company_id", middleware.AuthMiddleware(3), cont.GetCompany)
		api.GET("/companies", middleware.AuthMiddleware(3), cont.GetAllCompanies)

		// Driver endpoints
		api.POST("/drivers", middleware.AuthMiddleware(2), cont.CreateDriver)
		api.PUT("/drivers/:driver_id", middleware.AuthMiddleware(2), cont.UpdateDriver)
		api.DELETE("/drivers/:driver_id", middleware.AuthMiddleware(2), cont.DeleteDriver)
		api.GET("/drivers/:driver_id", middleware.AuthMiddleware(3), cont.GetDriver)
		api.GET("/drivers", middleware.AuthMiddleware(3), cont.GetAllDrivers)

		// Employee endpoints
		api.POST("/employees", cont.CreateEmployee)
		api.PUT("/employees/:employee_id", middleware.AuthMiddleware(2), cont.UpdateEmployee)
		api.DELETE("/employees/:employee_id", middleware.AuthMiddleware(2), cont.DeleteEmployee)
		api.GET("/employees/:employee_id", middleware.AuthMiddleware(3), cont.GetEmployee)
		api.GET("/employees", middleware.AuthMiddleware(3), cont.GetAllEmployees)

		// Logistic endpoints
		api.POST("/logistics", middleware.AuthMiddleware(1), cont.CreateLogistic)
		api.PUT("/logistics/:logistic_id", middleware.AuthMiddleware(3), cont.UpdateLogistic)
		api.DELETE("/logistics/:logistic_id", middleware.AuthMiddleware(1), cont.DeleteLogistic)
		api.GET("/logistics/:logistic_id", middleware.AuthMiddleware(3), cont.GetLogistic)
		api.GET("/logistics", middleware.AuthMiddleware(3), cont.GetAllLogistics)
		api.PUT("/logistics_with_cargo/:logistic_id", middleware.AuthMiddleware(3), cont.UpdateLogisticCargo)
		api.POST("/terminate_logistics", middleware.AuthMiddleware(3), cont.TerminateLogistic)
		api.POST("/cancel_late_logistics", middleware.AuthMiddleware(3), cont.CancelLateLogistic)
		api.GET("/logistics/overview", middleware.AuthMiddleware(3), cont.Overview)

		// Transaction endpoints
		api.POST("/transactions", middleware.AuthMiddleware(1), cont.CreateTransaction)
		api.PUT("/transactions/:transaction_id", middleware.AuthMiddleware(1), cont.UpdateTransaction)
		api.DELETE("/transactions/:transaction_id", middleware.AuthMiddleware(1), cont.DeleteTransaction)
		api.GET("/transactions/:transaction_id", middleware.AuthMiddleware(2), cont.GetTransaction)
		api.GET("/transactions", middleware.AuthMiddleware(2), cont.GetAllTransactions)

		// Performance endpoints
		api.POST("/performances", middleware.AuthMiddleware(2), cont.CreatePerformance)
		api.PUT("/performances/:performance_id", middleware.AuthMiddleware(2), cont.UpdatePerformance)
		api.DELETE("/performances/:performance_id", middleware.AuthMiddleware(2), cont.DeletePerformance)
		api.GET("/performances/:performance_id", middleware.AuthMiddleware(3), cont.GetPerformance)
		api.GET("/performances", middleware.AuthMiddleware(3), cont.GetAllPerformances)

		// History endpoints
		api.GET("/histories", middleware.AuthMiddleware(3), cont.GetAllHistories)
		api.GET("/histories/:history_id", middleware.AuthMiddleware(3), cont.GetHistory)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
