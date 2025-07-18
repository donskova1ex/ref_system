package handlers

import (
	"github.com/gin-gonic/gin"
	"ref_system/internal/repository"
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

}
