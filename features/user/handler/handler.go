package handler

import (
	"firly/mytaskapp/apps/middlewares"
	"firly/mytaskapp/features/user"
	"firly/mytaskapp/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	dataLogin, token, err := h.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	response := MapCoreUserToLogRes(dataLogin, token)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "login successfully", response))
}
func (h *UserHandler) CreateUser(c echo.Context) error {
	NewUser := new(UserRequest)

	// mendapatkan data yang dikirim oleh FE melalui request body
	err := c.Bind(&NewUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//mapping dari struct request to struct core
	input := MapReqToCoreUser(*NewUser)
	_, err = h.userService.Create(input)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success create user", nil))
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	result, err := h.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var usersResponse []UserResponse
	for _, v := range result {
		usersResponse = append(usersResponse, MapCoreUserToRes(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", usersResponse))
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	// Memeriksa apakah ID pengguna yang diambil dari token sama dengan ID yang diminta
	result, err := h.userService.GetById(uint(userID))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
		}
	}

	resultResponse := UserResponse{
		Name:        result.Name,
		Email:       result.Email,
		PhoneNumber: result.PhoneNumber,
		CreatedAt:   result.CreatedAt,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (h *UserHandler) UpdateUserById(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//Mapping user reques to core user
	Core := MapReqToCoreUser(userInput)
	err := h.userService.UpdateById(uint(userID), Core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Get user data for response
	user, err := h.userService.GetById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "user not found", nil))
	}
	resultResponse := MapCoreUserToRes(user)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "user updated successfully", resultResponse))
}

func (h *UserHandler) DeleteUserById(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	err := h.userService.DeleteById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete user", nil))
}
