package repositories

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/rubichandrap/hello-go/dtos"
	"github.com/rubichandrap/hello-go/utils"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Get() ([]dtos.User, RepositoryResult) {
	var users []dtos.User

	err := r.db.Find(&users).Error

	if err != nil {
		return users, RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when fetching data from database: %v", err),
		}
	}

	return users, RepositoryResult{}
}

func (r *UserRepository) GetByID(user *dtos.User, id *int) RepositoryResult {
	err := r.db.Find(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return RepositoryResult{
				Error:      fmt.Errorf("no data found"),
				StatusCode: fiber.StatusNotFound,
			}
		}

		return RepositoryResult{
			Error:      fmt.Errorf("failed when fetching data from database: %v", err),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	return RepositoryResult{}
}

func (r *UserRepository) Store(user *dtos.User) RepositoryResult {
	err := r.db.Create(&user).Error

	if err != nil {
		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when creating data to database: %v", err),
		}
	}

	return RepositoryResult{}
}

func (r *UserRepository) Update(user *dtos.User, id *int, u *dtos.UpdateUser) RepositoryResult {
	err := r.db.Find(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return RepositoryResult{
				Error:      fmt.Errorf("no data found: %v", err),
				StatusCode: fiber.StatusNotFound,
			}
		}

		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when fetching data from database: %v", err),
		}
	}

	user.Name = u.Name
	user.Email = u.Email

	err = r.db.Save(&user).Error

	if err != nil {
		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when updating data to database: %v", err),
		}
	}

	return RepositoryResult{}
}

func (r *UserRepository) UpdatePassword(id *int, u *dtos.UpdateUserPassword) RepositoryResult {
	var user dtos.User

	err := r.db.Find(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return RepositoryResult{
				Error:      fmt.Errorf("no data found: %v", err),
				StatusCode: fiber.StatusNotFound,
			}
		}

		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when fetching data from database: %v", err),
		}
	}

	if !utils.CheckPassword(*user.Password, *u.CurrentPassword) {
		return RepositoryResult{
			StatusCode: fiber.StatusForbidden,
			Error:      fmt.Errorf("wrong credentials"),
		}
	}

	hashedPassword, err := utils.HashPassword(*u.Password)

	if err != nil {
		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("error when hashing password: %v", err),
		}
	}

	user.Password = &hashedPassword

	err = r.db.Save(&user).Error

	if err != nil {
		return RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when updating data to database: %v", err),
		}
	}

	return RepositoryResult{}
}

func (r *UserRepository) IsFieldExists(field string, fieldValue interface{}) (bool, RepositoryResult) {
	var user dtos.User

	result := r.db.Find(&user, fmt.Sprintf("%s = ?", field), fieldValue)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, RepositoryResult{}
		}

		return false, RepositoryResult{
			StatusCode: fiber.StatusInternalServerError,
			Error:      fmt.Errorf("failed when fetching data from database: %v", result.Error),
		}
	}

	return result.RowsAffected > 0, RepositoryResult{}
}
