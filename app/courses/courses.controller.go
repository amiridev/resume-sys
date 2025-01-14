package courses

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	root    *core.Controller
	service CoursesServiceInterface
}

func NewCoursesController() *CoursesController {
	return &CoursesController{
		root:    core.GetController(),
		service: NewCoursesService(),
	}
}

// @tags		Courses
// @security	Bearer
// @router	/api/v1/courses [get]
// @summary	get list of courses
// @accept	json
// @produce	json
// @success 200 {object} core.Response[CoursesResponseType]
func (ctrl *CoursesController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	courses, _ := ctrl.service.List(user.ID)

	// response courses data as clean transform data
	ctrl.root.Success(c, CoursesResponse(courses))
}

// @tags     Courses
// @security Bearer
// @router   /api/v1/courses/{id} [get]
// @summary  get course by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CourseResponseType]
// @param    id path string true "Course ID"
func (ctrl *CoursesController) Show(c *gin.Context) {
	courseId := c.Param("id")

	course, err := ctrl.service.Show(courseId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, CourseResponse(course))
}

// @tags     Courses
// @security Bearer
// @router   /api/v1/courses [post]
// @summary  create new course api
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CourseResponseType]
// @param    request body CourseCreateDto true "Courses inputs"
func (ctrl *CoursesController) Create(c *gin.Context) {
	var dto CourseCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	course, _ := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, CourseResponse(course))
}

// @tags     Courses
// @security Bearer
// @router   /api/v1/courses/{id} [put]
// @summary  update course by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Course ID"
// @param    request body CourseUpdateDto true "course inputs"
func (ctrl *CoursesController) Update(c *gin.Context) {
	var dto CourseUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	courseId := c.Param("id")

	ctrl.service.Update(courseId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Courses
// @security Bearer
// @router   /api/v1/courses/{id} [delete]
// @summary  get Course by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Course ID"
func (ctrl *CoursesController) Delete(c *gin.Context) {
	courseId := c.Param("id")

	err := ctrl.service.Delete(courseId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
