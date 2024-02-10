package controller

import (
	"NYCard_Backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Poem struct {
}

func (p *Poem) AddPoem(c *gin.Context) {
	var form struct {
		Title         string `json:"title" binding:"required"`
		ContentFirst  string `json:"content_first" binding:"required"`
		ContentSecond string `json:"content_second" binding:"required"`
		Author        string `json:"author" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.Poem.AddPoem(form.Title, form.ContentFirst, form.ContentSecond, form.Author)

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (p *Poem) GetPoem(c *gin.Context) {
	resp, err := srv.Poem.GetPoem()

	if err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
