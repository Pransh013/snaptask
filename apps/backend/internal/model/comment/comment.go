package comment

import "github.com/Pransh013/snaptask/internal/model"

type Comment struct {
	model.Base
	TodoID  string `json:"todoId" db:"todo_id"`
	UserID  string `json:"userId" db:"user_id"`
	Content string `json:"content" db:"content"`
}
