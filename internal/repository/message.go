package repository

import (
	"context"
	"go-mygram/internal/infrastructure"
	"go-mygram/internal/model"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message model.Message) (model.Message, error)
	GetMessageByID(ctx context.Context, id uint64) (model.Message, error)
	UpdateMessage(ctx context.Context, id uint64, updatedMessage model.Message) (model.Message, error)
	DeleteMessage(ctx context.Context, id uint64) error
	GetMessagesByUserID(ctx context.Context, userID uint64) ([]model.Message, error)
	GetMessagesByPhotoID(ctx context.Context, photoID uint64) ([]model.Message, error)
}

type messageRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewMessageRepository(db infrastructure.GormPostgres) MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) CreateMessage(ctx context.Context, message model.Message) (model.Message, error) {
	err := r.db.GetConnection().WithContext(ctx).Create(&message).Error
	return message, err
}

func (r *messageRepositoryImpl) GetMessageByID(ctx context.Context, id uint64) (model.Message, error) {
	var message model.Message
	err := r.db.GetConnection().WithContext(ctx).Where("id = ?", id).First(&message).Error
	return message, err
}

func (r *messageRepositoryImpl) UpdateMessage(ctx context.Context, id uint64, updatedMessage model.Message) (model.Message, error) {
	err := r.db.GetConnection().WithContext(ctx).Model(&model.Message{}).Where("id = ?", id).Updates(updatedMessage).Error
	return updatedMessage, err
}

func (r *messageRepositoryImpl) DeleteMessage(ctx context.Context, id uint64) error {
	err := r.db.GetConnection().WithContext(ctx).Where("id = ?", id).Delete(&model.Message{}).Error
	return err
}

func (r *messageRepositoryImpl) GetMessagesByUserID(ctx context.Context, userID uint64) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.GetConnection().WithContext(ctx).Where("user_id = ?", userID).Find(&messages).Error
	return messages, err
}

func (r *messageRepositoryImpl) GetMessagesByPhotoID(ctx context.Context, photoID uint64) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.GetConnection().WithContext(ctx).Where("photo_id = ?", photoID).Find(&messages).Error
	return messages, err
}
