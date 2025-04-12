package User

import "github.com/waskull/gopharm/internal/ports"

type Handler struct {
	UserService ports.UserService
}
