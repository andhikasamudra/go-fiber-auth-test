package services

import (
	"github.com/form3tech-oss/jwt-go"
	"go-fiber-auth-test/config"
	"go-fiber-auth-test/dto"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(request dto.AuthLoginRequest) (*dto.AuthToken, error) {
	var response dto.AuthToken
	secret := []byte(config.Config("jwtSecret"))

	user, err := GetUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	// create access token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	response.AccessToken = t

	// create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(secret)
	if err != nil {
		return nil, err
	}

	response.RefreshToken = rt

	return &response, nil
}
