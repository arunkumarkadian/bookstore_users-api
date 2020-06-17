package users

import (
	"encoding/json"
	"github.com/arunkumarkadian/bookstore_users-api/domain/users"
	"github.com/arunkumarkadian/bookstore_users-api/services"
	"github.com/arunkumarkadian/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user Id")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return

	}

	result, createUsrErr := services.CreateUser(user)

	if createUsrErr != nil {
		c.JSON(createUsrErr.Status, createUsrErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "Implement me !! ")
}
