package User

import (
	"github.com/gin-gonic/gin"
	"github.com/waskull/gopharm/cmd/core"
	"github.com/waskull/gopharm/internal/domain"
)

func (h Handler) CreateEmployee(c *gin.Context) {
	var loginParams core.LoginParams
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	var fakeDataBasePassword string = "1234"

	user := &domain.User{
		Email:    loginParams.Email,
		Password: loginParams.Password,
	}

	if user.Password != fakeDataBasePassword {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidPassword, "Invalid password"))
		return
	}

	err := h.UserService.Create(user)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}
