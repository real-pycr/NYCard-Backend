package service

import (
	"NYCard_Backend/common"
	"NYCard_Backend/model"
	"crypto/rand"
	"encoding/base64"
	"strconv"
)

type Card struct {
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func (n *Card) AddCard(from string, to string, content string, id int) (any, error) {
	key, err := GenerateRandomString(18)
	if err != nil {
		return "", common.ErrNew(err, common.SysErr)
	}
	card := model.Card{
		From:    from,
		To:      to,
		Content: content,
		UserId:  id,
		Key:     key + strconv.Itoa(id),
	}
	if err := model.DB.Create(&card).Error; err != nil {
		return nil, err
	}
	card.Content = ""
	return card, nil
}

func (n *Card) GetCard(key string) (any, error) {
	card := model.Card{}
	if err := model.DB.Model(&model.Card{}).Where("`key` = ?", key).First(&card).Error; err != nil {
		return "", common.ErrNew(err, common.SysErr)
	}
	return card, nil
}

func (n *Card) DeleteCard(id int) (any, error) {
	if err := model.DB.Where("id = ?", id).Delete(&model.Card{}).Error; err != nil {
		return "", common.ErrNew(err, common.SysErr)
	}
	return nil, nil
}
func (n *Card) GetCardList(id int) (any, error) {
	var cards []model.Card
	if err := model.DB.Where("user_id = ?", id).Find(&cards).Error; err != nil {
		return nil, common.ErrNew(err, common.SysErr)
	}
	return struct {
		Total int          `json:"total"`
		Cards []model.Card `json:"cards"`
	}{
		Total: len(cards),
		Cards: cards,
	}, nil
}
