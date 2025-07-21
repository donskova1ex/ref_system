package handlers

import (
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

}
