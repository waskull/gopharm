package User

import (
	"github.com/gin-gonic/gin"
	"github.com/waskull/gopharm/cmd/core"
	"github.com/waskull/gopharm/internal/domain"
)

func (h Handler) GetPlayer(c *gin.Context) {
	playerIdParam := c.Param("id")
	player, err := h.UserService.Get(playerIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, player)
}

func (h Handler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		core.RespondError(c, err)
		return
	}
	c.JSON(200, users)
}

func (h Handler) CreateEmployee(c *gin.Context) {
	var employeeCreateParams core.EmployeeCreateParams
	if err := c.ShouldBindJSON(&employeeCreateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	user := &domain.User{
		Fullname: employeeCreateParams.Fullname,
		Email:    employeeCreateParams.Email,
		Password: employeeCreateParams.Password,
		Dni:      employeeCreateParams.Dni,
		Phone:    employeeCreateParams.Phone,
		Role:     employeeCreateParams.Role,
	}
	err := h.UserService.Create(user)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

func (h Handler) CreatePlayer(c *gin.Context) {
	var employeeCreateParams core.UserCreateParams
	if err := c.ShouldBindJSON(&employeeCreateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	user := &domain.User{
		Fullname: employeeCreateParams.Fullname,
		Email:    employeeCreateParams.Email,
		Password: employeeCreateParams.Password,
		Phone:    employeeCreateParams.Phone,
		Dni:      employeeCreateParams.Dni,
	}
	err := h.UserService.Create(user)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

func (h Handler) UpdatePlayer(c *gin.Context) {
	var userUpdateParams core.UserUpdateParams
	if err := c.ShouldBindJSON(&userUpdateParams); err != nil {
		core.RespondError(c, domain.NewAppError(domain.ErrCodeInvalidParams, err.Error()))
		return
	}

	user := &domain.User{
		Fullname: userUpdateParams.Fullname,
		Password: userUpdateParams.Password,
		Dni:      userUpdateParams.Dni,
		Phone:    userUpdateParams.Phone,
	}
	err := h.UserService.Update(user)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, user)
}

func (h Handler) DeletePlayer(c *gin.Context) {
	playerIdParam := c.Param("id")
	err := h.UserService.Delete(playerIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, nil)
}
