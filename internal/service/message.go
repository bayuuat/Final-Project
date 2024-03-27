package service

import (
	"context"
	"go-mygram/internal/model"
	"go-mygram/internal/repository"
	"time"
)

type MessageService interface {
	CreateMessage(ctx context.Context, userID uint64, messagePost model.MessagePost) (model.Message, error)
	GetMessagesByUserID(ctx context.Context, userID uint64) ([]model.Message, error)
	GetMessagesByPhotoID(ctx context.Context, photoID uint64) ([]model.Message, error)
	UpdateMessage(ctx context.Context, id uint64, updatedMessage model.MessagePost) (model.Message, error)
	DeleteMessage(ctx context.Context, id uint64) error
}

type messageServiceImpl struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(messageRepository repository.MessageRepository) MessageService {
	return &messageServiceImpl{
		messageRepository: messageRepository,
	}
}

func (s *messageServiceImpl) CreateMessage(ctx context.Context, userID uint64, messagePost model.MessagePost) (model.Message, error) {
	message := model.Message{
		UserID:    userID,
		PhotoID:   messagePost.PhotoID,
		Message:   messagePost.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.messageRepository.CreateMessage(ctx, message)
}

func (s *messageServiceImpl) GetMessagesByUserID(ctx context.Context, userID uint64) ([]model.Message, error) {
	return s.messageRepository.GetMessagesByUserID(ctx, userID)
}

func (s *messageServiceImpl) GetMessagesByPhotoID(ctx context.Context, photoID uint64) ([]model.Message, error) {
	return s.messageRepository.GetMessagesByPhotoID(ctx, photoID)
}

func (s *messageServiceImpl) UpdateMessage(ctx context.Context, id uint64, updatedMessage model.MessagePost) (model.Message, error) {
	// Check if the message exists
	message, err := s.messageRepository.GetMessageByID(ctx, id)
	if err != nil {
		return model.Message{}, err
	}

	// Update message fields
	message.Message = updatedMessage.Message
	message.UpdatedAt = time.Now()

	// Save the updated message
	return s.messageRepository.UpdateMessage(ctx, id, message)
}

func (s *messageServiceImpl) DeleteMessage(ctx context.Context, id uint64) error {
	// Check if the message exists
	_, err := s.messageRepository.GetMessageByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete the message
	return s.messageRepository.DeleteMessage(ctx, id)
}
