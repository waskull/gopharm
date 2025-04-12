package User

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/waskull/gopharm/internal/domain"
)

func (s *Service) Create(user *domain.User) (err error) {
	now := time.Now().UTC()
	user.CreatedAt = &now

	err = s.Repo.Insert(user)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Error ese usuario ya ha sido registrado")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creando el usuario: Registro duplicado",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creando el usuario: %w", err)
	}

	return nil
}

// Delete implements ports.UserService.
func (s *Service) Delete(id string) (err error) {
	err = s.Repo.Delete(id)
	if err != nil {
		return domain.ManageError(err, "Error borrando el usuario")
	}
	return nil
}

// Get implements ports.UserService.
func (s *Service) Get(id string) (user *domain.User, err error) {
	if id == "" {
		return nil, errors.New("la id es requerida")
	}

	user, err = s.Repo.Get(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.NewAppError(
				domain.ErrCodeNotFound,
				fmt.Sprintf("El usuario con la id: '%s' no existe", id))
		}
		if errors.Is(err, domain.ErrTimeout) {
			return nil, domain.NewAppError(
				domain.ErrCodeTimeout,
				"Se acabo el tiempo intente mas tarde")
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("error obteniendo el usuario: %w", err)
	}

	return user, nil
}

// Update implements ports.UserService.
func (s *Service) Update(id *domain.User) (err error) {
	panic("unimplemented")
}

func (s *Service) GetUsers() (players []*domain.User, err error) {
	players, err = s.Repo.GetUsers()
	if err != nil {
		return nil, domain.ManageError(err, "Error obteniendo los jugadores")
	}
	return players, nil
}
