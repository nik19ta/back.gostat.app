package http

import (
	auth "gostat/services/auth"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// SignIn @Summary Sign Up
// @Description Log in to the admin panel, get a token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} auth.SingUpResp
// @Failure 400
// @Failure 401
// @Router /api/auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	inp := new(auth.SignIn)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(inp.Mail, inp.Password)

	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, auth.SingUpResp{Token: *token})
}
