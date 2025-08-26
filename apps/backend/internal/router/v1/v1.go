package v1

import (
	"github.com/Pransh013/snaptask/internal/handler"
	"github.com/Pransh013/snaptask/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {
	registerTodoRoutes(router, handlers.Todo, handlers.Comment, middleware.Auth)
	registerCategoryRoutes(router, handlers.Category, middleware.Auth)
	registerCommentRoutes(router, handlers.Comment, middleware.Auth)
}
