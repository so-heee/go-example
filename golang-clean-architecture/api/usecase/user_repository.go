package usecase

import "github.com/so-heee/golang-example/golang-clean-architecture/api/domain"

type UserRepository interface {
	FindById(id int) (domain.User, error)
	FindAll() (domain.Users, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}
