package skills

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type SkillsController struct {
	root    *core.Controller
	service SkillsServiceInterface
}

func NewSkillsController() *SkillsController {
	return &SkillsController{
		root:    core.GetController(),
		service: NewSkillsService(),
	}
}

// @tags     Skills
// @security Bearer
// @router   /api/v1/skills [get]
// @summary  get list of skills
// @accept   json
// @produce  json
// @success  200 {skill} core.Response[SkillsResponseType]
func (ctrl *SkillsController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	skills := ctrl.service.List(user.ID)

	// response projects data as clean transform data
	ctrl.root.Success(c, SkillsResponse(skills))
}

// @tags     Skills
// @security Bearer
// @router   /api/v1/skills/{id} [get]
// @summary  get skill by id
// @accept   json
// @produce  json
// @success  200 {skill} core.Response[SkillResponseType]
// @param    id path string true "Skill ID"
func (ctrl *SkillsController) Show(c *gin.Context) {
	skillId := c.Param("id")

	skill, err := ctrl.service.Show(skillId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, SkillResponse(skill))
}

// @tags     Skills
// @security Bearer
// @router   /api/v1/skills [post]
// @summary  create new skill api
// @accept   json
// @produce  json
// @success  200 {skill} core.Response[SkillResponseType]
// @param    request body SkillCreateDto true "skill inputs"
func (ctrl *SkillsController) Create(c *gin.Context) {
	var dto SkillCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	skill := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, SkillResponse(skill))
}

// @tags     Skills
// @security Bearer
// @router   /api/v1/skills/{id} [put]
// @summary  update skill by id
// @accept   json
// @produce  json
// @success  200 {skill} core.SuccessResponse
// @param    id path string true "Skill ID"
// @param    request body SkillUpdateDto true "skill inputs"
func (ctrl *SkillsController) Update(c *gin.Context) {
	var dto SkillUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	skillId := c.Param("id")

	ctrl.service.Update(skillId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Skills
// @security Bearer
// @router   /api/v1/skills/{id} [delete]
// @summary  get skill by id
// @accept   json
// @produce  json
// @success  200 {skill} core.SuccessResponse
// @param    id path string true "Skill ID"
func (ctrl *SkillsController) Delete(c *gin.Context) {
	skillId := c.Param("id")

	err := ctrl.service.Delete(skillId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
