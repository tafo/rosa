package todo

func (request AddRequest) ToEntity() Item {
	return Item{
		Content: request.Content,
		IsCompleted: false,
	}
}
