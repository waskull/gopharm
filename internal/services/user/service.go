package User

import (
	"github.com/waskull/gopharm/internal/ports"
)

// Make sure Service implements the PlayerService interface
// at compile time.
var _ ports.UserService = &Service{}

type Service struct {
	Repo ports.UserRepository
}
