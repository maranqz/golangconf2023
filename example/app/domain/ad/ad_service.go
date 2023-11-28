package ad

import (
	"context"

	ltime "ddd/example/app/pkg/time"
)

type adService struct {
	userRepo     UserRepo
	categoryRepo CategoryRepo
	adRepo       AdRepo
	nower        ltime.Nower
}

func NewAdService(userRepo UserRepo, categoryRepo CategoryRepo, adRepo AdRepo) *adService {
	return &adService{userRepo: userRepo, categoryRepo: categoryRepo, adRepo: adRepo}
}

func (s *adService) CreateDraft(
	ctx context.Context,
	categoryID CategoryID,
	userID UserID,
	title Title,
	description Description,
) (*Ad, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !user.IsBanned() {
		return nil, nil
	}

	_, err = s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	adID, err := s.adRepo.NextID(ctx)
	if err != nil {
		return nil, err
	}

	return newAd(
		adID,
		userID,
		categoryID,
		title,
		description,
		s.nower.Now(),
	)
}
