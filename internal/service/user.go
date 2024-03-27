package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-mygram/internal/model"
	"go-mygram/internal/repository"
	"go-mygram/pkg/helper"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUsersById(ctx context.Context, id uint64) (model.User, error)
	UpdateUserByID(ctx context.Context, id uint64, updateUser model.UserUpdate) (model.User, error)
	DeleteUsersById(ctx context.Context, id uint64) (model.User, error)

	// activity
	SignUp(ctx context.Context, userSignUp model.UserSignUp) (model.User, error)
	SignIn(ctx context.Context, userSignIn model.UserSignIn) (model.User, error)

	// misc
	GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error)
}

type userServiceImpl struct {
	repo repository.UserQuery
}

func NewUserService(repo repository.UserQuery) UserService {
	return &userServiceImpl{repo: repo}
}

func (u *userServiceImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (u *userServiceImpl) GetUsersById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUsersByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (u *userServiceImpl) UpdateUserByID(ctx context.Context, id uint64, updateUser model.UserUpdate) (model.User, error) {
	// Get user by ID
	user, err := u.repo.GetUsersByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	// If user doesn't exist, return error
	if user.ID == 0 {
		return model.User{}, errors.New("user not found")
	}

	// Update user fields
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	// Save updated user
	updatedUser, err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (u *userServiceImpl) DeleteUsersById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUsersByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	// if user doesn't exist, return
	if user.ID == 0 {
		return model.User{}, nil
	}

	// delete user by id
	err = u.repo.DeleteUsersByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (u *userServiceImpl) SignUp(ctx context.Context, userSignUp model.UserSignUp) (model.User, error) {
	// assumption: semua user adalah user baru
	user := model.User{
		Username: userSignUp.Username,
		Email:    userSignUp.Email,
		Age:      userSignUp.Age,
	}

	pass, err := helper.GenerateHash(userSignUp.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = pass

	res, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return res, err
}

func (u *userServiceImpl) SignIn(ctx context.Context, userSignIn model.UserSignIn) (model.User, error) {
	// Retrieve user by email
	user, err := u.repo.FindByEmail(ctx, userSignIn.Email)
	if err != nil {
		return model.User{}, err
	}

	// Verify password
	if err := CompareHashAndPassword(user.Password, userSignIn.Password); err != nil {
		return model.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *userServiceImpl) GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error) {
	// generate claim
	now := time.Now()

	claim := model.StandardClaim{
		Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
		Iss: "go-middleware",
		Aud: "golang-006",
		Sub: "access-token",
		Exp: uint64(now.Add(time.Hour).Unix()),
		Iat: uint64(now.Unix()),
		Nbf: uint64(now.Unix()),
	}

	userClaim := model.AccessClaim{
		StandardClaim: claim,
		UserID:        user.ID,
		Username:      user.Username,
	}

	token, err = helper.GenerateToken(userClaim)
	return
}
