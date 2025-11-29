package service

import (
	"errors"
	"workshop-pit/entity"
	"workshop-pit/internal/repository"
	"workshop-pit/model"
	"workshop-pit/pkg/bcrypt"
	"workshop-pit/pkg/database/mariadb"
	"workshop-pit/pkg/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	RegisterUser(param model.UserRegisterParam) (*model.UserRegisterResponse, error)
	LoginUser(param model.UserLoginParam) (*model.UserLoginResponse, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
	db             *gorm.DB
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		userRepository: userRepository,
		db:             mariadb.Connection,
		bcrypt:         bcrypt,
		jwtAuth:        jwtAuth,
	}
}

func (u *UserService) RegisterUser(param model.UserRegisterParam) (*model.UserRegisterResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	_, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err == nil {
		return nil, errors.New("email already exists")
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:      userID,
		RoleID:      2,
		Name:        param.Name,
		Email:       param.Email,
		Password:    hashPassword,
		PhoneNumber: param.PhoneNumber,
	}

	err = u.userRepository.CreateUser(tx, user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	response := &model.UserRegisterResponse{
		Name:        param.Name,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
	}

	return response, nil
}

func (u *UserService) LoginUser(param model.UserLoginParam) (*model.UserLoginResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	user, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return nil, err
	}

	err = u.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.UserID, false)
	if err != nil {
		return nil, err
	}

	response := &model.UserLoginResponse{
		Token: token,
	}

	return response, nil
}
func (u *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return u.userRepository.GetUser(param)
}
