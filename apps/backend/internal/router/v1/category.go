package v1

import (
	"github.com/Pransh013/snaptask/internal/handler"
	"github.com/Pransh013/snaptask/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerCategoryRoutes(r *echo.Group, h *handler.CategoryHandler, auth *middleware.AuthMiddleware) {
	categories := r.Group("/categories")
	categories.Use(auth.RequireAuth)

	categories.POST("", h.CreateCategory)
	categories.GET("", h.GetCategories)

	dynamicCategory := categories.Group("/:id")
	dynamicCategory.PATCH("", h.UpdateCategory)
	dynamicCategory.DELETE("", h.DeleteCategory)
}
