package controller

import (
	"fmt"
	"net/http"

	"github.com/MACMREPO/switchaccount/model"
	"github.com/MACMREPO/switchaccount/service"
	"github.com/gin-gonic/gin"
)

// Controller used to access AddSwitchAccountService  [OR] [07062023]
func AddSwitchAccountController(c *gin.Context) {

	var NewUser model.SwitchAccountModel
	err := c.BindJSON(&NewUser)
	if err != nil {
		// Zerologs.Error().Msg("AddMPIN(): Error in c.BindJSON is " + err.Error())
		c.JSON(http.StatusBadRequest, "")
		return
	}

	service.AddSwitchAccountService(NewUser)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	c.JSON(http.StatusOK, "Account successfully added")
}

// Controller used to access DeleteSwitchAccountController  [OR] [07062023]
func DeleteSwitchAccountController(c *gin.Context) {

	var DeleteUser model.AccountModel
	err := c.BindJSON(&DeleteUser)
	if err != nil {
		// Zerologs.Error().Msg("AddMPIN(): Error in c.BindJSON is " + err.Error())
		c.JSON(http.StatusBadRequest, "")
		return
	}

	err = service.DeleteSwitchAccountService(DeleteUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Account successfully deleted")
}

// Controller used to access GetSwitchAccountController  [OR] [07062023]
func GetSwitchAccountController(c *gin.Context) {

	var GetUser model.AccountModel
	err := c.BindJSON(&GetUser)
	if err != nil {
		// fmt.Println("Bind Data Failed")
		// Zerologs.Error().Msg("AddMPIN(): Error in c.BindJSON is " + err.Error())
		c.JSON(http.StatusBadRequest, "")
		return
	}

	//response, err := service.GetSwitchAccountService(GetUser)
	response, err := service.GetRedisSwitchAccountService(GetUser)
	if err != nil {
		dbresponse, err := service.GetSwitchAccountService(GetUser)
		fmt.Println(dbresponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		for _, v := range dbresponse {
			service.AddUserInSwitchAccountRedis(v)
		}

		c.JSON(http.StatusOK, dbresponse)
		return
	} else {
		//todo: unmarshalll pending response
		c.JSON(http.StatusOK, response)
		return
	}

}

// Controller used to access GetSwitchAccountRedisController  [OR] [07062023]
// func GetSwitchAccountRedisController(c *gin.Context) {

// 	var GetUser model.AccountModel
// 	err := c.BindJSON(&GetUser)
// 	if err != nil {
// 		// fmt.Println("Bind Data Failed")
// 		// Zerologs.Error().Msg("AddMPIN(): Error in c.BindJSON is " + err.Error())
// 		c.JSON(http.StatusBadRequest, "")
// 		return
// 	}

// 	response, err := service.GetRedisSwitchAccountService(GetUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	var objUser []model.SwitchAccountModel
// 	err = json.Unmarshal([]byte(response), &objUser)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, objUser)
// }
