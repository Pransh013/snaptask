package v1

import (
	"github.com/Pransh013/snaptask/internal/handler"
	"github.com/Pransh013/snaptask/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerTodoRoutes(r *echo.Group, h *handler.TodoHandler, ch *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	todos := r.Group("/todos")
	todos.Use(auth.RequireAuth)

	todos.POST("", h.CreateTodo)
	todos.GET("", h.GetTodos)
	todos.GET("/stats", h.GetTodoStats)

	dynamicTodo := todos.Group("/:id")
	dynamicTodo.GET("", h.GetTodoByID)
	dynamicTodo.PATCH("", h.UpdateTodo)
	dynamicTodo.DELETE("", h.DeleteTodo)

	todoComments := dynamicTodo.Group("/comments")
	todoComments.POST("", ch.AddComment)
	todoComments.GET("", ch.GetCommentsByTodoID)

	todoAttachments := dynamicTodo.Group("/attachments")
	todoAttachments.POST("", h.UploadTodoAttachment)
	todoAttachments.DELETE("/:attachmentId", h.DeleteTodoAttachment)
	todoAttachments.GET("/:attachmentId/download", h.GetAttachmentPresignedURL)
}
