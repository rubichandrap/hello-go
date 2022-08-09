package repositories

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/rubichandrap/hello-go/dtos"
	"github.com/rubichandrap/hello-go/utils"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Login(payload *dtos.Login) RepositoryResult {
	var user dtos.User

	err := r.db.Find(&user, "email = ?", &payload.Email).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return RepositoryResult{
				Error:      fmt.Errorf("no data found: %v", err),
				StatusCode: fiber.StatusNotFound,
			}
		}

		return RepositoryResult{
			Error:      fmt.Errorf("failed when fetching data from database: %v", err),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	isMatched := utils.CheckPassword(*user.Password, *payload.Password)

	if !isMatched {
		return RepositoryResult{
			Error:      fmt.Errorf("password does not match"),
			StatusCode: fiber.StatusBadRequest,
		}
	}

	t, err := utils.GenerateNewAccessToken()

	if err != nil {
		return RepositoryResult{
			Error:      fmt.Errorf("failed generating token: %v", err),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	return RepositoryResult{
		Data: t,
	}
}
