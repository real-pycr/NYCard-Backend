package service

import (
	"NYCard_Backend/common"
	"NYCard_Backend/model"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
}

func (u *User) Register(username string, password string, qqnum string) (any, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	resp, err := http.Get("https://api.oioweb.cn/api/qq/info?qq=" + qqnum)
	if err != nil {
		return "", common.ErrNew(err, common.SysErr)
	}
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败：", err)
		return "", common.ErrNew(err, common.SysErr)
	}
	var data struct {
		Code   int `json:"code"`
		Result struct {
			User_id  int    `json:"user_id"`
			Nickname string `json:"nickname"`
			Gender   string `json:"sex"`
			Age      int    `json:"age"`
			Area     string `json:"area"`
		} `json:"result"`
		Msg string `json:"msg"`
	}
	if err := json.Unmarshal(respbody, &data); err != nil {
		fmt.Println("JSON解析失败", err)
		return "", common.ErrNew(err, common.SysErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", common.ErrNew(errors.New("QQ号码不存在"), common.ParamErr)
	}

	user := model.User{
		Username: username,
		Password: encryptedPassword,
		QQNumber: qqnum,
		QQName:   data.Result.Nickname,
		Age:      data.Result.Age,
		Area:     data.Result.Area,
		Gender:   data.Result.Gender,
	}
	if user.Gender == "" {
		user.Gender = "unknown"
	}
	if user.Area == "" {
		user.Area = "unknown"
	}
	if err := model.DB.
		Where("username = ? OR qq_number = ?", username, user.QQNumber).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := model.DB.Create(&user).Error; err != nil {
				return "", common.ErrNew(err, common.SysErr)
			}
			return user, nil
		}
		return "", common.ErrNew(err, common.SysErr)
	}
	return "", common.ErrNew(errors.New("用户名/QQ已注册"), common.ParamErr)
}

func (u *User) Login(name string, password string) (any, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	var user model.User
	if err := model.DB.Where("qq_number = ? AND password = ?", name, encryptedPassword).First(&user).Error; err != nil {
		if err := model.DB.Where("username = ? AND password = ?", name, encryptedPassword).First(&user).Error; err != nil {
			return "", common.ErrNew(errors.New("用户不存在或密码错误"), common.SysErr)
		}
	}
	return user, nil
}

func (u *User) GetUserStatus(id int64) (any, error) {
	var user model.User
	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return "", common.ErrNew(errors.New("数据库出现问题"), common.SysErr)
	}
	return user, nil
}
