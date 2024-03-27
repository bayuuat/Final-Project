package repository

import (
	"context"

	"go-mygram/internal/infrastructure"
	"go-mygram/internal/model"
)

type PhotoRepository interface {
	GetPhotos(ctx context.Context) ([]model.Photo, error)
	GetPhotoByID(ctx context.Context, id uint64) (model.Photo, error)
	UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	DeletePhotoByID(ctx context.Context, id uint64) error
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
}

type photoRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewPhotoRepository(db infrastructure.GormPostgres) PhotoRepository {
	return &photoRepositoryImpl{db: db}
}

func (p *photoRepositoryImpl) GetPhotos(ctx context.Context) ([]model.Photo, error) {
	db := p.db.GetConnection()
	photos := []model.Photo{}
	if err := db.
		WithContext(ctx).
		Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (p *photoRepositoryImpl) GetPhotoByID(ctx context.Context, id uint64) (model.Photo, error) {
	db := p.db.GetConnection()
	photo := model.Photo{}
	if err := db.
		WithContext(ctx).
		First(&photo, id).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (p *photoRepositoryImpl) UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := p.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Save(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (p *photoRepositoryImpl) DeletePhotoByID(ctx context.Context, id uint64) error {
	db := p.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Delete(&model.Photo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (p *photoRepositoryImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := p.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Create(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}
