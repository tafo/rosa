package todo

type AddRequest struct {
	Content string `json:"content"`
}

type AddResponse struct {
	Item
}

type CompleteRequest struct {
	Id int
}