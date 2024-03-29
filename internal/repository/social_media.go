package repository

import (
	"context"
	"go-mygram/internal/infrastructure"
	"go-mygram/internal/model"
	"time"
)

type SocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error)
	GetSocialMediaByID(ctx context.Context, id uint64) (model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMediaByID(ctx context.Context, id uint64) error
}

type socialMediaRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewSocialMediaRepository(db infrastructure.GormPostgres) SocialMediaRepository {
	return &socialMediaRepositoryImpl{db: db}
}

func (r *socialMediaRepositoryImpl) CreateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	socialMedia.CreatedAt = time.Now()
	socialMedia.UpdatedAt = time.Now()

	err := r.db.GetConnection().WithContext(ctx).Create(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepositoryImpl) GetSocialMediaByID(ctx context.Context, id uint64) (model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	err := r.db.GetConnection().WithContext(ctx).First(&socialMedia, id).Error
	return socialMedia, err
}

func (r *socialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	socialMedia.UpdatedAt = time.Now()

	err := r.db.GetConnection().WithContext(ctx).Save(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepositoryImpl) DeleteSocialMediaByID(ctx context.Context, id uint64) error {
	err := r.db.GetConnection().WithContext(ctx).Delete(&model.SocialMedia{}, id).Error
	return err
}
