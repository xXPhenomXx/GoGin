package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GoGin/pkg/app"
	"GoGin/pkg/e"
	"GoGin/pkg/util"
	"GoGin/services/user_service"
)

// GetUsers godoc
// @Summary Get users
// @Description Returns an array of all User objects
// @Produce  json
// @Tags user
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/users/all" [get]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	users, err := userservice.All()

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, data)
	}

	data["status"] = true
	data["users"] = users
	appG.Response(http.StatusOK, e.SUCCESS, data)

}


func GetOrCreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	// user service to get or create the user record
	user, err := userservice.GetOrCreate(c)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_CREATE, err)
	}

	// generate JWT token
	token, err := util.GenerateToken(user.Email, user.First + " " + user.Last)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
	}

	// return token and user
	data["status"] = true
	data["user"] = user
	data["token"] = token

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GetUserByEmail(c *gin.Context) {

	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	user, err := userservice.GetByEmail(c)

	if err != nil {
		data["status"] = false
		data["error"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_EMAIL , data)
		return
	}

	token, err := util.GenerateToken(user.Email, user.First + " " + user.Last)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
	}

	data["status"] = true
	data["user"] = user
	data["token"] = token

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
