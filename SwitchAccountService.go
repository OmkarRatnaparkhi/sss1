package service

import (
	"fmt"

	"github.com/MACMREPO/switchaccount/helper"
	"github.com/MACMREPO/switchaccount/model"
)

// Service used add user record in database and redis. [OR] [07062023]
func AddSwitchAccountService(NewUser model.SwitchAccountModel) {
	EnqueuToDBChannel(NewUser)
	EnqueuToRedisChannel(NewUser)
}

// Service used delete user record in database and redis. [OR] [07062023]
func DeleteSwitchAccountService(DeleteUser model.AccountModel) error {
	EnqueuToDBChannel(DeleteUser)
	EnqueuToRedisChannel(DeleteUser)

	return nil
}

// Service used get user record from database. [OR] [07062023]
func GetSwitchAccountService(GetUser model.AccountModel) ([]model.SwitchAccountModel, error) {
	var objUser []model.SwitchAccountModel
	err := Db.SwitchAccount.Where("parent_id=?", GetUser.ParentId).Find(&objUser).Error
	if err != nil {
		return nil, err
	}
	return objUser, nil
}

// Service used get user record from redis. [OR] [07062023]
func GetRedisSwitchAccountService(GetUser model.AccountModel) (string, error) {
	redisgetmodel, err := RedisLib.Client.Get(helper.GenerateRedisKey(GetUser.ParentId)).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return redisgetmodel, nil
}
