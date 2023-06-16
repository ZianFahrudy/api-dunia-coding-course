package formatter

import (
	"api-dunia-coding/domain/entity"
	"time"
)

type InformationFormatter struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
}

func Information(information entity.Information) InformationFormatter {
	informationFormatter := InformationFormatter{
		ID:    information.ID,
		Title: information.Title,
		Body:  information.Body,
		Date:  information.Date,
	}

	return informationFormatter
}

func Informations(informations []entity.Information) []InformationFormatter {
	informationsFormatter := []InformationFormatter{}

	for _, information := range informations {
		informationFormatter := Information(information)
		informationsFormatter = append(informationsFormatter, informationFormatter)
	}

	return informationsFormatter
}
