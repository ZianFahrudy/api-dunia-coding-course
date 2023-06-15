package model

// request
type GetEventDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type JoinEventInput struct {
	EventID  int `json:"event_id" binding:"required"`
	MemberID int `json:"member_id" binding:"required"`
}
