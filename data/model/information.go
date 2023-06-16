package model

type DeleteInformationInput struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteAllInformationInput struct {
	ID []int `json:"ids" binding:"required"`
}
