package services

import (
	"fmt"
	"strings"

	"godago-rest-api/internal/db"
	"godago-rest-api/internal/dto"
	"godago-rest-api/internal/errors"
	"godago-rest-api/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	userDB *db.UserDB
}

func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		userDB: db.NewUserDB(database),
	}
}

func userToResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (s *UserService) CreateUser(request *dto.CreateUserRequest) (*dto.UserResponse, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, errors.NewBadRequestError("Name cannot be empty")
	}

	if strings.TrimSpace(request.Email) == "" {
		return nil, errors.NewBadRequestError("Email cannot be empty")
	}

	user, err := s.userDB.CreateUser(request.Name, request.Email)
	if err != nil {
		return nil, errors.NewDatabaseError(err.Error())
	}

	return userToResponse(user), nil
}

func (s *UserService) GetUser(id int64) (*dto.UserResponse, error) {
	user, err := s.userDB.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
		}
		return nil, errors.NewDatabaseError(err.Error())
	}

	return userToResponse(user), nil
}

func (s *UserService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.userDB.GetAllUsers()
	if err != nil {
		return nil, errors.NewDatabaseError(err.Error())
	}

	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *userToResponse(&user)
	}

	return responses, nil
}

func (s *UserService) UpdateUser(id int64, request *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	if request.Name != nil && strings.TrimSpace(*request.Name) == "" {
		return nil, errors.NewBadRequestError("Name cannot be empty")
	}

	if request.Email != nil && strings.TrimSpace(*request.Email) == "" {
		return nil, errors.NewBadRequestError("Email cannot be empty")
	}

	user, err := s.userDB.UpdateUser(id, request.Name, request.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
		}
		return nil, errors.NewDatabaseError(err.Error())
	}

	return userToResponse(user), nil
}

func (s *UserService) DeleteUser(id int64) error {
	rowsAffected, err := s.userDB.DeleteUser(id)
	if err != nil {
		return errors.NewDatabaseError(err.Error())
	}

	if rowsAffected == 0 {
		return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
	}

	return nil
}
