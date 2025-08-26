package service

import (
	"github.com/Pransh013/snaptask/internal/lib/job"
	"github.com/Pransh013/snaptask/internal/repository"
	"github.com/Pransh013/snaptask/internal/server"
)

type Services struct {
	Auth     *AuthService
	Job      *job.JobService
	Todo     *TodoService
	Category *CategoryService
	Comment  *CommentService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Todo:     NewTodoService(s, repos.Todo, repos.Category),
		Category: NewCategoryService(s, repos.Category),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
	}, nil
}
