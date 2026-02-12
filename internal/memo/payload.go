package memo

type MemoRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type MemoUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
