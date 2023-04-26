package controllers

import (
	"fmt"
	"net/http"
	"project-my-gram/models"
	"project-my-gram/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

// UserRegister godoc
// @Summary Register user
// @Description Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param email query string true "email"
// @Param password query string true "password"
// @Param age query int true "age"
// @Success 201 {object} models.User "Register success response"
// @Router /users/register [post]
func (u *UserRepo) UserRegister(c *gin.Context) {
	contentType := pkg.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Created_at = time.Now()
	User.Updated_at = time.Now()

	if err := u.DB.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to create user data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":    User.GormModel.Id,
		"email": User.Email,
		"name":  User.Name,
	})
}

// UserLogin godoc
// @Summary Login user
// @Description Login user by email
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {object} interface{} "Login response"
// @Failure 401 "Unauthorized"
// @Router /users/login [post]
func (u *UserRepo) UserLogin(c *gin.Context) {
	contentType := pkg.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password := User.Password

	if err := u.DB.Debug().Where("email=?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	fmt.Println((User.Password), (password))
	if comparePass := pkg.ComparePass([]byte(User.Password), []byte(password)); !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	token := pkg.GenerateToken(uint(User.GormModel.Id), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
