package ports

import "github.com/waskull/gopharm/internal/domain"

type UserService interface {
	Create(user *domain.User) (err error)
	Get(id string) (user *domain.User, err error)
	Delete(id string) (err error)
	Update(id *domain.User) (err error)
	GetUsers() (players []*domain.User, err error)
}

type UserRepository interface {
	Insert(user *domain.User) (err error)
	Get(id string) (user *domain.User, err error)
	GetUsers() (players []*domain.User, err error)
	Delete(id string) (err error)
	Update(id *domain.User) (err error)
}
