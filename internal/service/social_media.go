package service

import (
	"context"
	"go-mygram/internal/model"
	"go-mygram/internal/repository"
	"time"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, userID uint64, socialMediaPost model.SocialMediaPost) (model.SocialMedia, error)
	GetSocialMediaByID(ctx context.Context, id uint64) (model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, id uint64, updatedSocialMedia model.SocialMediaPost) (model.SocialMedia, error)
	DeleteSocialMediaByID(ctx context.Context, id uint64) error
}

type socialMediaServiceImpl struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{
		socialMediaRepository: socialMediaRepository,
	}
}

func (s *socialMediaServiceImpl) CreateSocialMedia(ctx context.Context, userID uint64, socialMediaPost model.SocialMediaPost) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{
		Name:           socialMediaPost.Name,
		SocialMediaURL: socialMediaPost.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return s.socialMediaRepository.CreateSocialMedia(ctx, socialMedia)
}

func (s *socialMediaServiceImpl) GetSocialMediaByID(ctx context.Context, id uint64) (model.SocialMedia, error) {
	return s.socialMediaRepository.GetSocialMediaByID(ctx, id)
}

func (s *socialMediaServiceImpl) UpdateSocialMedia(ctx context.Context, id uint64, updatedSocialMedia model.SocialMediaPost) (model.SocialMedia, error) {
	existingSocialMedia, err := s.socialMediaRepository.GetSocialMediaByID(ctx, id)
	if err != nil {
		return model.SocialMedia{}, err
	}

	existingSocialMedia.Name = updatedSocialMedia.Name
	existingSocialMedia.SocialMediaURL = updatedSocialMedia.SocialMediaURL
	existingSocialMedia.UpdatedAt = time.Now()

	return s.socialMediaRepository.UpdateSocialMedia(ctx, existingSocialMedia)
}

func (s *socialMediaServiceImpl) DeleteSocialMediaByID(ctx context.Context, id uint64) error {
	return s.socialMediaRepository.DeleteSocialMediaByID(ctx, id)
}
