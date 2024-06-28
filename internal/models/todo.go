package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string
	UserID uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) InsertTodo(newData Todo) error {
	qry := tm.db.Create(&newData)
	if err := qry.Error; err != nil {
		return err
	}

	return nil
}

func (tm *TodoModel) GetTodoByUser(userID uint) ([]Todo, error) {
	var result []Todo
	qry := tm.db.Where("user_id = ?", userID).Find(&result)
	if err := qry.Error; err != nil {
		return nil, err
	}

	return result, nil
}
