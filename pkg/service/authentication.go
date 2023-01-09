package service

import (
	"errors"

	"github.com/xybor/xyauth/internal/database"
	"github.com/xybor/xyauth/internal/models"
	"github.com/xybor/xyauth/internal/utils"
	"github.com/xybor/xyauth/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Authenticate(email, password string) error {
	var cred = models.UserCredential{}
	result := database.Get().Where("email=?", email).Take(&cred)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return NotFoundError.Newf("could not find the email %s", email)
		}
		utils.GetLogger().Event("query-user-failed").
			Field("email", email).Field("password", password).
			Field("error", result.Error).Warning()
		return ServiceError.New("failed to authenticate")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(password)); err != nil {
		utils.GetLogger().Event("password-compare-failed").
			Field("email", email).Field("password", password).
			Field("error", err).Warning()
		return CredentialError.New("failed to authenticate")
	}

	return nil
}

func CreateToken(email string) (string, error) {
	var user = models.User{}
	result := database.Get().Where("email=?", email).Take(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", NotFoundError.Newf("could not find the email %s", email)
		}
		utils.GetLogger().Event("query-user-failed").
			Field("email", email).Field("error", result.Error).Warning()
		return "", ServiceError.New("failed to authenticate")
	}

	token, err := token.CreateAccessToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
