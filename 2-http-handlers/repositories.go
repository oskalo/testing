package handlers

type Comment struct {
	ID        int    `json:"id"`
	ProductID string `json:"product_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
}

type commentsRepository struct{}

func (c commentsRepository) SaveComment(comment Comment) error {
	return nil
}

func (c commentsRepository) GetComments(productID string) ([]Comment, error) {
	return []Comment{
		{
			ID:        1,
			ProductID: "de5ff4db-8196-4b22-b23b-2d44528ac725",
			Author:    "user",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		},
		{
			ID:        2,
			ProductID: "de5ff4db-8196-4b22-b23b-2d44528ac725",
			Author:    "user2",
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			ID:        3,
			ProductID: "de5ff4db-8196-4b22-b23b-2d44528ac725",
			Author:    "user1",
			Content:   "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		},
	}, nil
}

/*func (c commentsRepository) SaveComment(comment Comment) error {
	return fmt.Errorf("connection error")
}

func (c commentsRepository) GetComments(productID string) ([]Comment, error) {
	return nil, fmt.Errorf("connection error")
}*/
