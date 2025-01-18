package educations

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type EducationsController struct {
	root    *core.Controller
	service EducationsServiceInterface
}

func NewEducationsController() *EducationsController {
	return &EducationsController{
		root:    core.GetController(),
		service: NewEducationsService(),
	}
}

// @tags     Educations
// @security Bearer
// @router   /api/v1/educations [get]
// @summary  get list of educations
// @accept   json
// @produce  json
// @success  200 {object} core.Response[EducationsResponseType]
func (ctrl *EducationsController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	educations, _ := ctrl.service.List(user.ID)

	// response educations data as clean transform data
	ctrl.root.Success(c, EducationsResponse(educations))
}

// @tags     Educations
// @security Bearer
// @router   /api/v1/educations/{id} [get]
// @summary  get education by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[EducationResponseType]
// @param    id path string true "Education ID"
func (ctrl *EducationsController) Show(c *gin.Context) {
	educationId := c.Param("id")

	education, err := ctrl.service.Show(educationId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, EducationResponse(education))
}

// @tags     Educations
// @security Bearer
// @router   /api/v1/educations [post]
// @summary  create new education api
// @accept   json
// @produce  json
// @success  200 {object} core.Response[EducationResponseType]
// @param    request body EducationCreateDto true "education inputs"
func (ctrl *EducationsController) Create(c *gin.Context) {
	var dto EducationCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	education, _ := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, EducationResponse(education))
}

// @tags     Educations
// @security Bearer
// @router   /api/v1/educations/{id} [put]
// @summary  update education by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Education ID"
// @param    request body EducationUpdateDto true "education inputs"
func (ctrl *EducationsController) Update(c *gin.Context) {
	var dto EducationUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	educationId := c.Param("id")

	ctrl.service.Update(educationId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Educations
// @security Bearer
// @router   /api/v1/educations/{id} [delete]
// @summary  delete education by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Education ID"
func (ctrl *EducationsController) Delete(c *gin.Context) {
	educationId := c.Param("id")

	err := ctrl.service.Delete(educationId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
