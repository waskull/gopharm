package User

import (
	"errors"
	"fmt"

	"github.com/waskull/gopharm/internal/domain"
	"github.com/waskull/gopharm/internal/ports"
	"gorm.io/gorm"
)

// Make sure Repository implements ports.PlayerRepository
// at compile time
var _ ports.UserRepository = &Repository{}

type Repository struct {
	DB *gorm.DB
}

// Delete implements ports.UserRepository.
func (r *Repository) Delete(id string) (err error) {
	panic("unimplemented")
}

// Get implements ports.UserRepository.
func (r *Repository) Get(id string) (user *domain.User, err error) {
	var u domain.User
	err = r.DB.Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}

// GetPlayers implements ports.UserRepository.
func (r *Repository) GetUsers() (users []*domain.User, err error) {
	var u []*domain.User
	err = r.DB.Find(&u).Error
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los usuarios: %w", err)
	}

	return u, nil
}

// Insert implements ports.UserRepository.
func (r *Repository) Insert(user *domain.User) (err error) {
	panic("unimplemented")
}

// Update implements ports.UserRepository.
func (r *Repository) Update(id *domain.User) (err error) {
	panic("unimplemented")
}
