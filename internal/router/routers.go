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

func (b *Builder) UserRouters() {
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
		api.GET("/referral-codes", nil)
		api.POST("/referral-codes", nil)
		api.GET("/referral-codes/:code", nil)
		api.GET("/referral-codes/user/:uuid", nil)

	}
}

func (b *Builder) GetEngine() *gin.Engine {
	return b.engine
}
