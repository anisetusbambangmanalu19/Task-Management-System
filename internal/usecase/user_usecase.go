package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
	"github.com/anisetusbambangmanalu19/task-management/internal/repository"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *UserUsecase) Register(name, email, password string) (*entity.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	err = u.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

var jwtSecret = []byte("supersecretkey")

func (u *UserUsecase) Login(email, password string) (string, error) {

	user, err := u.UserRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
