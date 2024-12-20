package user

import (
	"errors"
	"net/http"

	"github.com/TrinityKnights/Backend/internal/delivery/http/handler"
	"github.com/TrinityKnights/Backend/internal/domain/model"
	"github.com/TrinityKnights/Backend/internal/service/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandlerImpl struct {
	Log  *logrus.Logger
	User user.UserService
}

func NewUserHandler(log *logrus.Logger, userService user.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		Log:  log,
		User: userService,
	}
}

// Register function is a handler to register a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "User data"
// @Success 201 {object} model.Response[model.UserResponse]
// @Failure 400 {object} model.Error
// @Failure 500 {object} model.Error
// @Router /users [post]
func (h *UserHandlerImpl) Register(ctx echo.Context) error {
	request := new(model.RegisterRequest)
	if err := ctx.Bind(request); err != nil {
		h.Log.Errorf("failed to bind request: %v", err)
		return handler.HandleError(ctx, 400, errors.New(http.StatusText(http.StatusBadRequest)))
	}

	response, err := h.User.Register(ctx.Request().Context(), request)
	if err != nil {
		h.Log.Errorf("failed to register: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		case errors.Is(err, errors.New(http.StatusText(http.StatusConflict))):
			return handler.HandleError(ctx, 409, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	return ctx.JSON(http.StatusCreated, model.NewResponse(response, nil))
}

// Login function is a handler to login user
// @Summary Login user
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "User data"
// @Success 200 {object} model.Response[model.TokenResponse]
// @Failure 400 {object} model.Error
// @Failure 401 {object} model.Error
// @Failure 500 {object} model.Error
// @Router /users/login [post]
func (h *UserHandlerImpl) Login(ctx echo.Context) error {
	request := new(model.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		h.Log.Errorf("failed to bind request: %v", err)
		return handler.HandleError(ctx, 400, errors.New(http.StatusText(http.StatusBadRequest)))
	}

	response, err := h.User.Login(ctx.Request().Context(), request)
	if err != nil {
		h.Log.Errorf("failed to login: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		case errors.Is(err, errors.New(http.StatusText(http.StatusUnauthorized))):
			return handler.HandleError(ctx, 401, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	return ctx.JSON(http.StatusOK, model.NewResponse(response, nil))
}

// Profile function is a handler to get user profile
// @Summary Get user profile
// @Description Get user profile
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} model.Response[model.UserResponse]
// @Failure 400 {object} model.Error
// @Failure 404 {object} model.Error
// @Failure 500 {object} model.Error
// @security ApiKeyAuth
// @Router /users [get]
func (h *UserHandlerImpl) Profile(ctx echo.Context) error {
	response, err := h.User.Profile(ctx.Request().Context())
	if err != nil {
		h.Log.Errorf("failed to get profile: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		case errors.Is(err, errors.New(http.StatusText(http.StatusNotFound))):
			return handler.HandleError(ctx, 404, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	return ctx.JSON(http.StatusOK, model.NewResponse(response, nil))
}

// Update function is a handler to update user profile
// @Summary Update user profile
// @Description Update user profile
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.UpdateRequest true "User data"
// @Success 200 {object} model.Response[model.UserResponse]
// @Failure 400 {object} model.Error
// @Failure 404 {object} model.Error
// @Failure 500 {object} model.Error
// @security ApiKeyAuth
// @Router /users [put]
func (h *UserHandlerImpl) Update(ctx echo.Context) error {
	request := new(model.UpdateRequest)
	if err := ctx.Bind(request); err != nil {
		h.Log.Errorf("failed to bind request: %v", err)
		return handler.HandleError(ctx, 400, errors.New(http.StatusText(http.StatusBadRequest)))
	}

	response, err := h.User.Update(ctx.Request().Context(), request)
	if err != nil {
		h.Log.Errorf("failed to update: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		case errors.Is(err, errors.New(http.StatusText(http.StatusNotFound))):
			return handler.HandleError(ctx, 404, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	return ctx.JSON(http.StatusOK, model.NewResponse(response, nil))
}

// RefreshToken function is a handler to refresh token
// @Summary Refresh token
// @Description Refresh token
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.RefreshTokenRequest true "User data"
// @Success 200 {object} model.Response[model.TokenResponse]
// @Failure 400 {object} model.Error
// @Failure 401 {object} model.Error
// @Failure 500 {object} model.Error
// @Router /users/refresh [post]
func (h *UserHandlerImpl) RefreshToken(ctx echo.Context) error {
	request := new(model.RefreshTokenRequest)
	if err := ctx.Bind(request); err != nil {
		h.Log.Errorf("failed to bind request: %v", err)
		return handler.HandleError(ctx, 400, errors.New(http.StatusText(http.StatusBadRequest)))
	}

	response, err := h.User.RefreshToken(ctx.Request().Context(), request)
	if err != nil {
		h.Log.Errorf("failed to refresh token: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		case errors.Is(err, errors.New(http.StatusText(http.StatusUnauthorized))):
			return handler.HandleError(ctx, 401, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	return ctx.JSON(http.StatusOK, model.NewResponse(response, nil))
}

// RequestResetPassword function is a handler to request reset password via email
func (h *UserHandlerImpl) RequestReset(ctx echo.Context) error {
	// Binding request to model
	request := new(model.RequestReset)
	if err := ctx.Bind(request); err != nil {
		h.Log.Errorf("failed to bind request: %v", err)
		return handler.HandleError(ctx, 400, errors.New(http.StatusText(http.StatusBadRequest)))
	}

	// Create a RequestResetPassword model from the email
	resetRequest := &model.RequestReset{
		Email: request.Email,
	}

	// Call the service method
	err := h.User.RequestReset(ctx.Request().Context(), resetRequest)
	if err != nil {
		h.Log.Errorf("failed to request reset password: %v", err)
		switch {
		case errors.Is(err, errors.New(http.StatusText(http.StatusBadRequest))):
			return handler.HandleError(ctx, 400, err)
		default:
			return handler.HandleError(ctx, 500, err)
		}
	}

	// Send success response
	return ctx.JSON(http.StatusOK, model.NewResponse("Reset password request successful. Please check your email.", nil))
}
