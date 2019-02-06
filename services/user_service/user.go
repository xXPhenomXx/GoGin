package userservice

import (
	"github.com/gin-gonic/gin"
	"GoGin/models"
)

// GetByEmail func queries for and returns a single user object by email address
func GetByEmail(c *gin.Context) (models.User, error) {
	email := c.Param("email")

	user, err := models.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

// GetOrCreate func
func GetOrCreate(c *gin.Context) (models.User, error){

	// fetch params
	email := c.PostForm("email")
	first := c.PostForm("first")
	last := c.PostForm("last")

	user, err := models.GetUserByEmail(email)

	if err != nil {

		if err.Error() == "not found" {
			user, err = models.CreateUser(email, first, last)
		} else {
			return models.User{}, c.Error(err)
		}
	}

	return user, nil
}

// All func
func All() ([]models.User, error) {

	return models.GetUsers()

}