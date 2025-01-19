package profiles

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	root    *core.Controller
	service ProfileServiceInterface
}

func NewProfileController() *ProfileController {
	return &ProfileController{
		root:    core.GetController(),
		service: NewProfileService(),
	}
}

// @tags     Profile
// @security Bearer
// @router   /api/v1/profile [get]
// @summary  get user profile
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProfileResponseType]
func (ctrl *ProfileController) GetProfile(c *gin.Context) {
	user := core.User(c) // احراز هویت از طریق توکن
	profile, err := ctrl.service.GetProfile(user.ID)
	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}
	ctrl.root.Success(c, ProfileResponse(profile))
}

// @tags     Profile
// @security Bearer
// @router   /api/v1/profile [post]
// @summary  create user profile
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProfileResponseType]
// @param    body body ProfileCreateDto true "Create Profile"
func (ctrl *ProfileController) CreateProfile(c *gin.Context) {
	var dto ProfileCreateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}
	user := core.User(c)
	profile, err := ctrl.service.CreateProfile(dto, user.ID)
	if err != nil {
		ctrl.root.InternalServerError(c, err)
		return
	}
	ctrl.root.Success(c, ProfileResponse(profile))
}

// @tags     Profile
// @security Bearer
// @router   /api/v1/profile [put]
// @summary  update user profile
// @accept   json
// @produce  json
// @success  200 {object} core.Response[ProfileResponseType]
// @param    body body ProfileUpdateDto true "Update Profile"
func (ctrl *ProfileController) UpdateProfile(c *gin.Context) {
	var dto ProfileUpdateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}
	user := core.User(c)
	profile, err := ctrl.service.UpdateProfile(dto, user.ID)
	if err != nil {
		ctrl.root.InternalServerError(c, err)
		return
	}
	ctrl.root.Success(c, ProfileResponse(profile))
}

// @tags     Profile
// @security Bearer
// @router   /api/v1/user-details [get]
// @summary  get all user details (profile, courses, projects, etc.)
// @accept   json
// @produce  json
// @success  200 {object} core.Response[UserDetailsResponseType]
func (ctrl *ProfileController) GetUserDetails(c *gin.Context) {
	user := core.User(c)
	userDetails, err := ctrl.service.GetUserDetails(user.ID)
	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}
	ctrl.root.Success(c, UserDetailsResponse(userDetails))
}
