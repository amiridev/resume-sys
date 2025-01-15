package documents

import (
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

type DocumentsController struct {
	root    *core.Controller
	service DocumentsServiceInterface
}

func NewDocumentsController() *DocumentsController {
	return &DocumentsController{
		root:    core.GetController(),
		service: NewDocumentsService(),
	}
}

// @tags     Documents
// @security Bearer
// @router   /api/v1/Documents [get]
// @summary  get list of documents
// @accept   json
// @produce  json
// @success  200 {object} core.Response[DocumentsResponseType]
func (ctrl *DocumentsController) List(c *gin.Context) {
	// get authenticated user from gin Context
	user := core.User(c)

	// get list data from service & repository
	documents, _ := ctrl.service.List(user.ID)

	// response Documents data as clean transform data
	ctrl.root.Success(c, DocumentsResponse(documents))
}

// @tags     Documents
// @security Bearer
// @router   /api/v1/documents/{id} [get]
// @summary  get document by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[DocumentResponseType]
// @param    id path string true "Document ID"
func (ctrl *DocumentsController) Show(c *gin.Context) {
	documentId := c.Param("id")

	document, err := ctrl.service.Show(documentId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, DocumentResponse(document))
}

// @tags     Documents
// @security Bearer
// @router   /api/v1/Documents [post]
// @summary  create new document api
// @accept   json
// @produce  json
// @success  200 {object} core.Response[DocumentResponseType]
// @param    request body DocumentCreateDto true "document inputs"
func (ctrl *DocumentsController) Create(c *gin.Context) {
	var dto DocumentCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)

	document, _ := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, DocumentResponse(document))
}

// @tags     Documents
// @security Bearer
// @router   /api/v1/Documents/{id} [put]
// @summary  update document by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Document ID"
// @param    request body DocumentUpdateDto true "document inputs"
func (ctrl *DocumentsController) Update(c *gin.Context) {
	var dto DocumentUpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	documentId := c.Param("id")

	ctrl.service.Update(documentId, dto)

	ctrl.root.Success(c, nil)
}

// @tags     Documents
// @security Bearer
// @router   /api/v1/Documents/{id} [delete]
// @summary  get document by id
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Document ID"
func (ctrl *DocumentsController) Delete(c *gin.Context) {
	documentId := c.Param("id")

	err := ctrl.service.Delete(documentId)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
