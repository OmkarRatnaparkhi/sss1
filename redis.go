package service

import (
	"encoding/json"
	"fmt"

	"github.com/MACMREPO/libgochannel"
	"github.com/MACMREPO/switchaccount/helper"
	"github.com/MACMREPO/switchaccount/model"
)

var (
	RedisChannel = libgochannel.CreateQueue()
)

func redisEnqueue(ch chan<- interface{}, mod libgochannel.ChanMod) {
	DBChannel <- mod
}

func enqueueredis(objmodel any) {
	redisEnqueue(DBChannel, libgochannel.ChanMod{Model: objmodel, Func: func(m any) {
		DequeueRedisChannel(m)
	}})
}

func EnqueuToRedisChannel(objmodel any) {
	enqueueredis(objmodel)
}

func DequeueRedisChannel(objmodel any) {
	switch value := objmodel.(type) {
	case model.SwitchAccountModel:
		AddUserInSwitchAccountRedis(value)
	case model.AccountModel:
		DeleteSwitchAccountRedis(value)
	}
}
// Function used add user record in redis. [OR] [07062023]
func AddUserInSwitchAccountRedis(NewUser model.SwitchAccountModel) {

	arrobj := []model.SwitchAccountModel{}

	helper.GenerateRedisKey(NewUser.ParentId)

	redisgetdata, err := RedisLib.Client.Get(helper.GenerateRedisKey(NewUser.ParentId)).Result()
	fmt.Println(redisgetdata)
	if err != nil {
		data, err := json.Marshal(NewUser)
		if err != nil {
			return
		}
		err = RedisLib.Client.Set(helper.GenerateRedisKey(NewUser.ParentId), data, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		redisgetmodel, err := RedisLib.Client.Get(helper.GenerateRedisKey(NewUser.ParentId)).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		userdatamodel := model.SwitchAccountModel{}

		err1 := json.Unmarshal([]byte(redisgetmodel), &userdatamodel)
		if err1 != nil {
			fmt.Println(err1)
		}

		if userdatamodel.ParentId == NewUser.ParentId {
			arrobj = append(arrobj, userdatamodel)
			arrobj = append(arrobj, NewUser)

			adddata, err := json.Marshal(arrobj)
			if err != nil {
				return
			}
			err = RedisLib.Client.Set(helper.GenerateRedisKey(NewUser.ParentId), adddata, 0).Err()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}

// Function used delete user record in database. [OR] [07062023]
func DeleteSwitchAccountRedis(DeleteUser model.AccountModel) error {
	err := RedisLib.Client.Del(helper.GenerateRedisKey(DeleteUser.ParentId)).Err()
	if err != nil {
		Zerologs.Error().Msg(" ParentId redis key not deleted for user=" + DeleteUser.ParentId + "_SA" + err.Error())
		return err
	}
	return nil
}
