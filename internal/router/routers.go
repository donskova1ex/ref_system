package router

import (
	"ref_system/internal/handlers"
	"ref_system/internal/repository"

	"github.com/gin-gonic/gin"
)

const ()

type Builder struct {
	engine     *gin.Engine
	repository *repository.Repository
}

func InitBuilder(repo *repository.Repository) *Builder {
	return &Builder{
		engine:     gin.Default(),
		repository: repo,
	}
}

func (b *Builder) UserRoutersBuilder() {
	userRepo := repository.NewUserRepository(b.repository)
	userHandler := handlers.NewUserHandler(userRepo)

	api := b.engine.Group("/api/v1")
	{
		api.POST("/users", userHandler.Create)
		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:uuid", userHandler.GetByUUID)
	}
}

func (b *Builder) ReferralCodeBuilder() {
	referralCodeRepo := repository.NewReferralCodeRepository(b.repository)
	referralCodeHandler := handlers.NewReferralCodeHandler(referralCodeRepo)

	api := b.engine.Group("/api/v1")
	{
		api.GET("/referral-codes", referralCodeHandler.GetAllReferralCodes)
		api.POST("/referral-codes", referralCodeHandler.Create)
		api.GET("/referral-codes/:code", referralCodeHandler.GetByCode)
		api.GET("/referral-codes/users/:uuid", referralCodeHandler.GetByUserUUID)
		api.POST("/referral-codes/users", referralCodeHandler.CreateNewCodeAndNewUser)

	}
}

func (b *Builder) GetEngine() *gin.Engine {
	return b.engine
}
