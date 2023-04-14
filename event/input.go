package event

type GetEventDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
