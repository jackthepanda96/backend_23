package todos

import "apibe23/internal/models"

type TodoRequest struct {
	Title string `json:"title"`
}

func ToModelTodo(tr TodoRequest, userID uint) models.Todo {
	return models.Todo{
		UserID: userID,
		Title:  tr.Title,
	}
}
