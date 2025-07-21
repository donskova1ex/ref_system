package handlers

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"ref_system/internal"
	"ref_system/internal/domain"
	"ref_system/internal/models"
	"ref_system/internal/repository"

	"github.com/gin-gonic/gin"
)

type ReferralCodeHandler struct {
	ReferralCodeRepository *repository.ReferralCodeRepository
}

func NewReferralCodeHandler(repo *repository.ReferralCodeRepository) *ReferralCodeHandler {
	return &ReferralCodeHandler{
		ReferralCodeRepository: repo,
	}
}

func (r *ReferralCodeHandler) GetAllReferralCodes(c *gin.Context) {
	codes, err := r.ReferralCodeRepository.GetAll()
	if err != nil {
		switch errors.Is(err, internal.ErrRecordNoFound) {
		case true:
			apiErr := HandleError(http.StatusNotFound, "codes no found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		default:
			apiErr := HandleError(http.StatusInternalServerError, "error getting codes", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
	}
	c.JSON(http.StatusOK, codes)
}
func (r *ReferralCodeHandler) Create(c *gin.Context) {
	domainReferralCode := &domain.ReferralCode{}
	if err := c.BindJSON(domainReferralCode); err != nil {
		apiErr := HandleError(http.StatusBadRequest, "error parsing body", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	referralCode := &models.ReferralCode{
		OwnerUUID: domainReferralCode.OwnerUUID,
		Code:      domainReferralCode.Code,
	}
	code, err := r.ReferralCodeRepository.Create(referralCode)
	if err != nil {
		apiErr := HandleError(http.StatusInternalServerError, "error creating referral code", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, code)
}
func (r *ReferralCodeHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		apiErr := HandleError(http.StatusBadRequest, "code is required", nil)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
	}
	referralCode, err := r.ReferralCodeRepository.GetByCode(code)
	if err != nil {
		switch errors.Is(err, internal.ErrRecordNoFound) {
		case true:
			apiErr := HandleError(http.StatusNotFound, "codes not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		default:
			apiErr := HandleError(http.StatusInternalServerError, "error getting referral code", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
	}
	c.JSON(http.StatusOK, referralCode)
}
func (r *ReferralCodeHandler) GetByUserUUID(c *gin.Context) {
	stringOwnerUUID := c.Param("uuid")
	if stringOwnerUUID == "" {
		apiErr := HandleError(http.StatusBadRequest, "user uuid is required", nil)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	ownerUUID, err := uuid.Parse(stringOwnerUUID)
	if err != nil {
		apiErr := HandleError(http.StatusInternalServerError, "error parsing uuid", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	code, err := r.ReferralCodeRepository.GetByOwnerUUID(ownerUUID)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrRecordNoFound):
			apiErr := HandleError(http.StatusNotFound, "code not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		case errors.Is(err, internal.ErrOwnerNotFound):
			apiErr := HandleError(http.StatusNotFound, "owner not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		default:
			apiErr := HandleError(http.StatusInternalServerError, "error getting referral code", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
	}
	c.JSON(http.StatusOK, code)

}
func (r *ReferralCodeHandler) CreateNewCodeAndNewUser(c *gin.Context) {
	domainReferralCode := &domain.ReferralCode{}
	if err := c.BindJSON(domainReferralCode); err != nil {
		apiErr := HandleError(http.StatusBadRequest, "error parsing body", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	if domainReferralCode.OwnerUUID == nil {
		apiErr := HandleError(http.StatusBadRequest, "owner_uuid required", internal.ErrOwnerUUIDRequired)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}

	referralCode := &models.ReferralCode{
		OwnerUUID: domainReferralCode.OwnerUUID,
		Code:      domainReferralCode.Code,
	}

	newReferralCode, err := r.ReferralCodeRepository.CreateNewCodeAndNewUser(referralCode)
	if err != nil {
		if errors.Is(err, internal.ErrOwnerNotFound) {
			apiErr := HandleError(http.StatusNotFound, "owner not found", err)
			c.AbortWithStatusJSON(apiErr.Status, apiErr)
		}
		apiErr := HandleError(http.StatusBadRequest, "error creating referral code", err)
		c.AbortWithStatusJSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, newReferralCode)

}
