package user

import (
	"github.com/gin-gonic/gin"
	"main/domain/users"
	"main/services"
	"main/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err!= nil {
	//	fmt.Println("err ", err)
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil{
	//	//TODO: Handle Json error
	//	fmt.Println("err ", err)
	//	return
	//}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadReqError("invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr!= nil{
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64); {
		if userErr != nil {
			err := errors.NewBadReqError("user should be number")
			c.JSON(err.Status, err)
			return
		}
	}
	user, getErr := services.GetUser(userId)
	if getErr!= nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}