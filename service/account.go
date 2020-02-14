package service

import (
	"demoproject/conf"
	"demoproject/util"
)

type Demo struct {
	UserId     int    `json:"user_id"`

}

func (a *Demo) DemoService() util.Errno {
	var pwd string
	sql := "select user_pwd from user2b where user_id=?"
	err := conf.DB.Get(&pwd, sql, a.UserId)
	var result = "it's ok"
	if err != nil {
		return util.FAILURE
	}

	return util.Errno{
		200,
		"success",
		result,
	}
}
