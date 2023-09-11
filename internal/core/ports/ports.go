package ports

import "github.com/AntonyIS/notelify-users-service/internal/core/domain"

type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	ReadUserWithEmail(email string) (*domain.User, error)
	ReadUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) (string, error)
	DeleteAllUsers() (string, error)
}

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	ReadUserWithEmail(email string) (*domain.User, error)
	ReadUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) (string, error)
	DeleteAllUsers() (string, error)
}
