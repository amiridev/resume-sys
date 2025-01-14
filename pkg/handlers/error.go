package handlers

import (
	"resume-sys/core"
	"resume-sys/pkg/messages"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func InternalErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	c.AbortWithStatusJSON(
		500,
		core.ToResponse(
			messages.MessageInternalError,
			map[string]any{},
			map[string]any{
				"reason": goErr.Error(),
			},
		),
	)
}
