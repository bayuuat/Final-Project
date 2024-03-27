package service

import (
	"context"
	"errors"

	"go-mygram/internal/model"
	"go-mygram/internal/repository"
)

type PhotoService interface {
	GetPhotos(ctx context.Context) ([]model.Photo, error)
	GetPhotoByID(ctx context.Context, id uint64) (model.Photo, error)
	UpdatePhoto(ctx context.Context, id uint64, updatedPhoto model.PhotoPost) (model.Photo, error)
	DeletePhotoByID(ctx context.Context, id uint64) error
	CreatePhoto(ctx context.Context, userID uint64, photo model.PhotoPost) (model.Photo, error)
}

type photoServiceImpl struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		photoRepository: photoRepository,
	}
}

func (s *photoServiceImpl) GetPhotos(ctx context.Context) ([]model.Photo, error) {
	return s.photoRepository.GetPhotos(ctx)
}

func (s *photoServiceImpl) GetPhotoByID(ctx context.Context, id uint64) (model.Photo, error) {
	return s.photoRepository.GetPhotoByID(ctx, id)
}

func (s *photoServiceImpl) UpdatePhoto(ctx context.Context, id uint64, updatedPhoto model.PhotoPost) (model.Photo, error) {
	// Get photo by ID
	photo, err := s.photoRepository.GetPhotoByID(ctx, id)
	if err != nil {
		return model.Photo{}, err
	}
	// If photo doesn't exist, return error
	if photo.ID == 0 {
		return model.Photo{}, errors.New("photo not found")
	}

	// Update photo fields
	photo.Title = updatedPhoto.Title
	photo.Caption = updatedPhoto.Caption
	photo.PhotoURL = updatedPhoto.PhotoURL

	// Save updated photo
	updatedPhotoResult, err := s.photoRepository.UpdatePhoto(ctx, photo)
	if err != nil {
		return model.Photo{}, err
	}

	return updatedPhotoResult, nil
}

func (s *photoServiceImpl) DeletePhotoByID(ctx context.Context, id uint64) error {
	return s.photoRepository.DeletePhotoByID(ctx, id)
}

func (s *photoServiceImpl) CreatePhoto(ctx context.Context, userID uint64, photo model.PhotoPost) (model.Photo, error) {
	newPhoto := model.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   userID,
	}

	return s.photoRepository.CreatePhoto(ctx, newPhoto)
}
