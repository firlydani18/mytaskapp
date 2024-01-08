package handler

import (
	"firly/mytaskapp/apps/middlewares"
	"firly/mytaskapp/features/project"
	"firly/mytaskapp/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type projectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *projectHandler {
	return &projectHandler{
		projectService: service,
	}
}

func (h *projectHandler) CreateProject(c echo.Context) error {
	NewProject := new(ProjectRequest)
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	NewProject.UserID = middlewares.ExtractTokenUserId(c)
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&NewProject)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to CoreProject
	input := MapProjReqToCoreProject(*NewProject)
	_, err = h.projectService.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success create project", nil))
}

func (h *projectHandler) GetAllProject(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	UserID := middlewares.ExtractTokenUserId(c)
	result, err := h.projectService.GetAll(UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var projectsResponse []ProjectResponse
	for _, v := range result {
		projectsResponse = append(projectsResponse, MapCoreProjToProjRes(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", projectsResponse))
}

func (h *projectHandler) GetProjectById(c echo.Context) error {
	idParam := c.Param("project_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.projectService.GetById(uint(idConv), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := MapCoreProjToProjRes(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (h *projectHandler) UpdateProjectById(c echo.Context) error {
	idParam := c.Param("project_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	// Mengambil data input dari permintaan
	projectInput := ProjectRequest{}
	errBind := c.Bind(&projectInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//Mapping project request to core project
	coreProject := MapProjReqToCoreProject(projectInput)

	// Melakukan pembaruan data proyek di service
	err = h.projectService.UpdateById(uint(idConv), userID, coreProject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Mendapatkan data proyek yang telah diperbarui untuk respon
	updatedProject, err := h.projectService.GetById(uint(idConv), userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "project not found", nil))
	}
	// Mapping updated project to ProjectResponse
	resultResponse := MapCoreProjToProjRes(updatedProject)
	// Kirim respon JSON
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "project updated successfully", resultResponse))
}

func (h *projectHandler) DeleteProjectById(c echo.Context) error {
	idParam := c.Param("project_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	err = h.projectService.DeleteById(uint(idConv), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete project", nil))
}
