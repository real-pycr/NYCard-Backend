package controller

import (
	"NYCard_Backend/common"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Card struct {
}

func (n *Card) AddCard(c *gin.Context) {
	session := SessionGet(c, "userinfo")
	if session == nil {
		c.Error(common.ErrNew(errors.New("用户未登录"), common.AuthErr))
		return
	}
	var form struct {
		From    string `json:"from" binding:"required"`
		To      string `json:"to" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.Card.AddCard(form.From, form.To, form.Content, int(session.(UserSession).ID))

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (n *Card) GetCard(c *gin.Context) {
	var form struct {
		Key string `form:"key" binding:"required"`
	}
	if err := c.ShouldBindQuery(&form); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.Card.GetCard(form.Key)

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (n *Card) DeleteCard(c *gin.Context) {
	var form struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&form); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.Card.DeleteCard(form.ID)

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (n *Card) GetCardList(c *gin.Context) {
	session := SessionGet(c, "userinfo")
	if session == nil {
		c.Error(common.ErrNew(errors.New("用户未登录"), common.AuthErr))
		return
	}

	resp, err := srv.Card.GetCardList(int(session.(UserSession).ID))

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
