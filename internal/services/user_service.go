package services

import (
	"context"

	"github.com/joaoCrulhas/ampel/internal/domain"
	"github.com/joaoCrulhas/ampel/internal/dtos"
	"github.com/joaoCrulhas/ampel/internal/repositories"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, input *dtos.CreateUserDTO) (*domain.User, error)
}

type UserService struct {
	userRepository       repositories.UserRepositoryInterface
	departmentRepository repositories.DepartmentRepositoryInterface
}

func (service *UserService) CreateUser(ctx context.Context, input *dtos.CreateUserDTO) (*domain.User, error) {

}

func NewUserService(userRepository repositories.UserRepositoryInterface, departmentRepository repositories.DepartmentRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository:       userRepository,
		departmentRepository: departmentRepository,
	}
}
