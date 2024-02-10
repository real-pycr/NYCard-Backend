package validator

import (
	ut "github.com/go-playground/universal-translator"
)

func timingTransZh(ut ut.Translator) error {
	return ut.Add("timing", "{0}输入的时间不符合要求", true)
}

func qqmailTransZh(ut ut.Translator) error {
	return ut.Add("qqmail", "{0}请输入qq邮箱", true)
}

func qqnumTransZh(ut ut.Translator) error {
	return ut.Add("qqnum", "请输入正确 qq 号", true)
}
