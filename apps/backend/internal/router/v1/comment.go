package v1

import (
	"github.com/Pransh013/snaptask/internal/handler"
	"github.com/Pransh013/snaptask/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerCommentRoutes(r *echo.Group, h *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	comments := r.Group("/comments")
	comments.Use(auth.RequireAuth)

	dynamicComment := comments.Group("/:id")
	dynamicComment.PATCH("", h.UpdateComment)
	dynamicComment.DELETE("", h.DeleteComment)
}
