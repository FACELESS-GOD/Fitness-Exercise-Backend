package Util

import (
	"errors"
	"time"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	"github.com/golang-jwt/jwt/v5"
)

type TokenProcessor struct {
}

func (TokenProc *TokenProcessor) CreateToken(UserDT StructStore.UserData, Secret string) (string, error) {

	if UserDT.Designation <= 0 {
		err := errors.New("Invalid Data")
		return "", err
	} else if UserDT.UserName == "" {
		err := errors.New("Invalid Data")
		return "", err
	} else if Secret == "" {
		err := errors.New("Invalid Data")
		return "", err
	}

	claims := jwt.MapClaims{
		"UserName": UserDT.UserName,
		"RoleID":   UserDT.Designation,
		"Exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tkn.SignedString(ConfigSetup.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
