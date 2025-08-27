package service

import (
	"fmt"

	"github.com/Pransh013/snaptask/internal/lib/aws"
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

	awsClient, err := aws.NewAWS(s)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS client: %w", err)
	}

	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Todo:     NewTodoService(s, repos.Todo, repos.Category, awsClient),
		Category: NewCategoryService(s, repos.Category),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
	}, nil
}
