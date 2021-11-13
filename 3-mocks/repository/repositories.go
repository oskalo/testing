package repository

import "fmt"

type Comment struct {
	ID        int    `json:"id"`
	ProductID string `json:"product_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
}

type commentsRepository struct{}

type CommentsRepository interface {
	SaveComment(comment Comment) error
	GetComments(productID string) ([]Comment, error)
}

func NewCommentsRepository() CommentsRepository {
	return new(commentsRepository)
}

func (c commentsRepository) SaveComment(comment Comment) error {
	return fmt.Errorf("connection error")
}

func (c commentsRepository) GetComments(productID string) ([]Comment, error) {
	return nil, fmt.Errorf("connection error")
}
