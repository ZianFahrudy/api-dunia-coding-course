package myevent

type PresenceInput struct {
	EventID int `uri:"id" binding:"required"`
}
