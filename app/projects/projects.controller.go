package projects

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type ProjectsController struct {
	root    *core.Controller
	service ProjectsServiceInterface
}

func NewProjectsController() *ProjectsController {
	return &ProjectsController{
		root:    core.GetController(),
		service: NewProjectsService(),
	}
}

// @tags     Projects
// @security Bearer
// @router   /api/v1/projects [get]
// @summary  get list of projects
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProjectsResponseType]
func (ctrl *ProjectsController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	projects, _ := ctrl.service.List(user.ID)

	// response projects data as clean transform data
	ctrl.root.Success(c, ProjectsResponse(projects))
}

// @tags     Projects
// @security Bearer
// @router   /api/v1/projects/{id} [get]
// @summary  get project by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProjectResponseType]
// @param    id path string true "Project ID"
func (ctrl *ProjectsController) Show(c *gin.Context) {
	projectId := c.Param("id")

	project, err := ctrl.service.Show(projectId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, ProjectResponse(project))
}

// @tags     Projects
// @security Bearer
// @router   /api/v1/projects [post]
// @summary  create new project api
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProjectResponseType]
// @param    request body ProjectCreateDto true "project inputs"
func (ctrl *ProjectsController) Create(c *gin.Context) {
	var dto ProjectCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	project, _ := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, ProjectResponse(project))
}

// @tags     Projects
// @security Bearer
// @router   /api/v1/projects/{id} [put]
// @summary  update project by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Project ID"
// @param    request body ProjectUpdateDto true "project inputs"
func (ctrl *ProjectsController) Update(c *gin.Context) {
	var dto ProjectUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	projectId := c.Param("id")

	ctrl.service.Update(projectId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Projects
// @security Bearer
// @router   /api/v1/projects/{id} [delete]
// @summary  get project by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Project ID"
func (ctrl *ProjectsController) Delete(c *gin.Context) {
	projectId := c.Param("id")

	err := ctrl.service.Delete(projectId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
