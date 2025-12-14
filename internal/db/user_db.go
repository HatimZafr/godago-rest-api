package db

import (
	"godago-rest-api/internal/models"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) CreateUser(name, email string) (*models.User, error) {
	user := &models.User{
		Name:  name,
		Email: email,
	}

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDB) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := u.db.Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserDB) UpdateUser(id int64, name, email *string) (*models.User, error) {
	user, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if name != nil {
		updates["name"] = *name
	}
	if email != nil {
		updates["email"] = *email
	}

	if len(updates) > 0 {
		if err := u.db.Model(user).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	return u.GetUserByID(id)
}

func (u *UserDB) DeleteUser(id int64) (int64, error) {
	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
