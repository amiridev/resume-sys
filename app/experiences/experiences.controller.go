package experiences

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type ExperiencesController struct {
	root    *core.Controller
	service ExperiencesServiceInterface
}

func NewExperiencesController() *ExperiencesController {
	return &ExperiencesController{
		root:    core.GetController(),
		service: NewExperiencesService(),
	}
}

// @tags     Experiences
// @security Bearer
// @router   /api/v1/experiences [get]
// @summary  get list of experiences
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ExperiencesResponseType]
func (ctrl *ExperiencesController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	experiences, _ := ctrl.service.List(user.ID)

	// response experiences data as clean transform data
	ctrl.root.Success(c, ExperiencesResponse(experiences))
}

// @tags     Experiences
// @security Bearer
// @router   /api/v1/experiences/{id} [get]
// @summary  get experience by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ExperienceResponseType]
// @param    id path string true "Experience ID"
func (ctrl *ExperiencesController) Show(c *gin.Context) {
	experienceId := c.Param("id")

	experience, err := ctrl.service.Show(experienceId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, ExperienceResponse(experience))
}

// @tags     Experiences
// @security Bearer
// @router   /api/v1/experiences [post]
// @summary  create new experience api
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ExperienceResponseType]
// @param    request body ExperienceCreateDto true "experience inputs"
func (ctrl *ExperiencesController) Create(c *gin.Context) {
	var dto ExperienceCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	experience, _ := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, ExperienceResponse(experience))
}

// @tags     Experiences
// @security Bearer
// @router   /api/v1/experiences/{id} [put]
// @summary  update experience by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Experience ID"
// @param    request body ExperienceUpdateDto true "experience inputs"
func (ctrl *ExperiencesController) Update(c *gin.Context) {
	var dto ExperienceUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	experienceId := c.Param("id")

	ctrl.service.Update(experienceId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Experiences
// @security Bearer
// @router   /api/v1/experiences/{id} [delete]
// @summary  delete experience by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Experience ID"
func (ctrl *ExperiencesController) Delete(c *gin.Context) {
	experienceId := c.Param("id")

	err := ctrl.service.Delete(experienceId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
