package handlers

import "ref_system/internal/repository"

type ReferralCodeHandler struct {
	ReferralCodeRepository *repository.ReferralCodeRepository
}

func NewReferralCodeHandler(repo *repository.ReferralCodeRepository) *ReferralCodeHandler {
	return &ReferralCodeHandler{
		ReferralCodeRepository: repo,
	}
}
