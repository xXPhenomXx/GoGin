package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xxphenomxx/GoGin/pkg/app"
	"github.com/xxphenomxx/GoGin/pkg/e"
	"github.com/xxphenomxx/GoGin/pkg/util"
	"github.com/xxphenomxx/GoGin/services/user_service"
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


// @Summary Get or Create a User object - returns a User object and JWT token
// @Description Queries for a single User object by the email address provided, if it cannot locate one then it will create a new User record and return it
// @Produce  json
// @Param email path string true "Email"
// @Param gID path string true "Google ID"
// @Param image path string true "Image URL"
// @Param full path string true "Users full name"
// @Param fn path string true "First name"
// @Param ln path string true "Last name"
// @Tags user
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user/get-or-create" [post]
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

// @Summary Fetch a user record by email address
// @Description Queries for a single User object by the email address provided otherwise will return an error
// @Produce  json
// @Tags user
// @Param email path string true "Email"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user/email/{email} [get]
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
