package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
	"xm-test-ilya-chicherin/internal/controllers/middleware"
	"xm-test-ilya-chicherin/internal/controllers/swagger"
	"xm-test-ilya-chicherin/internal/kafka"
	consumer "xm-test-ilya-chicherin/internal/kafka/jobs"
	"xm-test-ilya-chicherin/repository"
	"xm-test-ilya-chicherin/usecase"
)

const basePath = "/api/v1"

func New(router *gin.Engine) error {
	h := router.Group(basePath)
	if err := swagger.Register(h, GetSwagger, basePath); err != nil {
		return err
	}

	companyController, _ := NewCompanyController()
	kafkaWriter := repository.NewKafkaRepository(kafka.NewKafkaWriter())
	companyUC := usecase.NewCompany(kafkaWriter)
	companyController.CompanyUC = companyUC
	authUC := usecase.NewAuth()
	authController := NewAuthController(authUC)
	companyRepo := repository.NewRepository()
	kafkaConsumer := consumer.NewKafkaConsumer(os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_TOPIC"), companyUC, companyRepo)
	kafkaConsumer.StartDeleteConsumer(context.Background())
	kafkaConsumer.StartCreateConsumer(context.Background())
	kafkaConsumer.StartUpdateConsumer(context.Background())

	initializePublicRoutes(h, companyController, authController)
	initializeProtectedRoutes(h, companyController)
	return nil
}

func initializePublicRoutes(router *gin.RouterGroup, controller *CompanyController, authController *AuthController) {
	guest := router
	{
		guest.GET("/health", controller.Health)
		guest.POST("/auth/login", authController.Login)
		guest.GET("/companies/:id", controller.GetCompany)
	}
}

func initializeProtectedRoutes(router *gin.RouterGroup, controller *CompanyController) {
	protected := router.Group("/")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.DELETE("/companies/:id", controller.DeleteCompany)
		protected.POST("/companies", controller.CreateCompany)
		protected.PATCH("/companies/:id", controller.PatchCompany)
	}
}
