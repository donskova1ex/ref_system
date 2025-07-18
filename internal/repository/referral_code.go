package repository

type ReferralCodeRepository struct {
	repository *Repository
}

func NewReferralCodeRepository(repo *Repository) *ReferralCodeRepository {
	return &ReferralCodeRepository{
		repository: repo,
	}
}

//TODO: referral_code generator....
