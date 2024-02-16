package service

import (
	"AuthService/internal/consts"
	"AuthService/internal/model"
	"AuthService/internal/repository"
	se "AuthService/pkg/error"
	"AuthService/pkg/logger"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	GenerateTokens(guid string) (model.Tokens, error)
	RefreshTokens(refreshToken, guid string) (model.Tokens, error)
}

type AuthServiceImpl struct {
	repo   repository.AuthRepository
	logger *logger.Logger
}

func (a AuthServiceImpl) GenerateTokens(guid string) (model.Tokens, error) {

	accessToken, err := model.GenerateToken(guid, consts.AccessTokenTTL)

	if err != nil {
		return model.Tokens{}, &se.ResponseError{Message: consts.ErrGenToken, Code: http.StatusInternalServerError}
	}

	refreshToken := generateRefreshHash(guid)

	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 10)

	if err != nil {
		return model.Tokens{}, &se.ResponseError{Message: consts.ErrBcrypt, Code: http.StatusInternalServerError}
	}

	err = a.repo.SaveToken(guid, string(hash))

	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateRefreshHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	time.Sleep(1 * time.Millisecond)
	hash.Write([]byte(time.Now().String()))

	token := fmt.Sprintf("%x", hash.Sum(nil))

	return token[len(token)-30:]
}

func (a AuthServiceImpl) RefreshTokens(refreshToken, guid string) (model.Tokens, error) {
	res, err := a.repo.FindTokens(guid)
	if err != nil {
		return model.Tokens{}, err
	}

	doc := findToken(res, refreshToken)

	if len(doc.Token) == 0 {
		return model.Tokens{}, &se.ResponseError{Message: consts.ErrInvalidToken, Code: http.StatusBadRequest}
	}

	if err := a.repo.DeleteToken(doc.ID); err != nil {
		return model.Tokens{}, err
	}

	return a.GenerateTokens(guid)
}

func findToken(docs []model.MongoDoc, token string) model.MongoDoc {

	for _, doc := range docs {
		if bcrypt.CompareHashAndPassword([]byte(doc.Token), []byte(token)) == nil {
			return doc
		}
	}

	return model.MongoDoc{}
}
