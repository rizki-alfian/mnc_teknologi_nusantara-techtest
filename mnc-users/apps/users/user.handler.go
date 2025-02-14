package users

import (
	"net/http"
	"mnc-users/apps/users/dto"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req dto.UserRegisterDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	userResponse, err := h.Service.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": userResponse,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	var req dto.UserLoginDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	accessToken, refreshToken, err := h.Service.Login(req.PhoneNumber, req.PIN)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Phone Number and PIN doesn’t match."})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var req dto.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	profile, err := h.Service.UpdateProfile(userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": profile,
	})
}