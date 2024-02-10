package service

import "NYCard_Backend/model"

type Service struct {
	Hello
	User
	Card
	Poem
}

func New() *Service {
	if err := model.DB.Model(&model.Poem{}).Count(&poemNum).Error; err != nil {
		panic(err)
	}
	service := &Service{}
	return service
}
