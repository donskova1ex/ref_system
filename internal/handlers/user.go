package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"ref_system/internal"
	"ref_system/internal/domain"
	"ref_system/internal/models"
	"ref_system/internal/repository"
)

type UserHandler struct {
	UserRepository *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: repo,
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	domainUser := &domain.User{}
	if err := c.ShouldBindJSON(domainUser); err != nil {
		apiErr := HandleError(http.StatusBadRequest, "invalid body", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}

	user := &models.User{
		UUID: domainUser.UUID,
	}
	createdUser, err := u.UserRepository.Create(user)
	if err != nil {
		apiErr := HandleError(http.StatusInternalServerError, "failed to create user", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, createdUser)

}
func (u *UserHandler) GetByUUID(c *gin.Context) {
	uuidStr := c.Param("uuid")
	if uuidStr == "" {
		apiErr := HandleError(http.StatusBadRequest, "uuid required", nil)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}

	newUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		apiErr := HandleError(http.StatusBadRequest, "invalid uuid format", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}

	user, err := u.UserRepository.GetByUUID(&newUUID)
	if err != nil {

		switch errors.Is(err, internal.ErrRecordNoFound) {
		case true:
			apiErr := HandleError(http.StatusNotFound, "user not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		default:
			apiErr := HandleError(http.StatusInternalServerError, "failed to get user", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}

	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		switch errors.Is(err, internal.ErrRecordNoFound) {
		case true:
			apiErr := HandleError(http.StatusNotFound, "users not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		default:
			apiErr := HandleError(http.StatusInternalServerError, "failed to get users", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
	}

	c.JSON(http.StatusOK, users)
}
