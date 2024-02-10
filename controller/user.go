package controller

import (
	"NYCard_Backend/common"
	"NYCard_Backend/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Register(c *gin.Context) {
	var form struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		QQNumber string `json:"qqnumber" binding:"required,qqnum"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.User.Register(form.Username, form.Password, form.QQNumber)

	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(err)
		return
	}
	SessionSet(c, "userinfo", UserSession{ID: resp.(model.User).ID, Username: resp.(model.User).Username})
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) Login(c *gin.Context) {
	var form struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}

	resp, err := srv.User.Login(form.Username, form.Password)

	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(err)
		return
	}
	SessionSet(c, "userinfo", UserSession{ID: resp.(model.User).ID, Username: resp.(model.User).Username})
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (u *User) Logout(c *gin.Context) {
	if session := SessionGet(c, "userinfo"); session == nil {
		c.Error(common.ErrNew(errors.New("用户未登录"), common.AuthErr))
		return
	}
	SessionClear(c)
	c.JSON(http.StatusOK, ResponseNew(c, nil))
}

func (u *User) GetUserStatus(c *gin.Context) {
	session := SessionGet(c, "userinfo")
	if session == nil {
		c.Error(common.ErrNew(errors.New("用户未登录"), common.AuthErr))
		return
	}

	resp, err := srv.User.GetUserStatus(session.(UserSession).ID)

	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
