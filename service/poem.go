package service

import (
	"NYCard_Backend/common"
	"NYCard_Backend/model"
	"math/rand"
)

type Poem struct {
}

var poemNum int64

func (p *Poem) AddPoem(title string, contentFirst string, contentSecond string, author string) (any, error) {
	poem := model.Poem{
		Title:         title,
		ContentFirst:  contentFirst,
		ContentSecond: contentSecond,
		Author:        author,
	}
	if err := model.DB.Create(&poem).Error; err != nil {
		return nil, err
	}
	poemNum++
	return poem, nil
}

func (p *Poem) GetPoem() (any, error) {
	randomNum := rand.Intn(int(poemNum)) + 1
	poem := model.Poem{}
	if err := model.DB.Where("id = ?", randomNum).First(&poem).Error; err != nil {
		return nil, common.ErrNew(err, common.SysErr)
	}
	return poem, nil
}
