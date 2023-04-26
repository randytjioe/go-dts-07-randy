package controllers

import (
	"net/http"
	"project-my-gram/models"
	"project-my-gram/pkg"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MediaRepo struct {
	DB *gorm.DB
}

// CreateSocialMedia godoc
// @Summary Create social media
// @Description Create social media of the user
// @Tags social media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201 {object} models.Media "Create social media success"
// @Failure 401 "Unauthorized"
// @Router /socialmedia/uploadsocmed [post]
func (m *MediaRepo) UploadMedia(c *gin.Context) {
	contentType := pkg.GetContentType(c)
	Media := models.Media{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.User_id = uint(userId)
	Media.Created_at = time.Now()
	Media.Updated_at = time.Now()

	if err := m.DB.Debug().Create(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to upload social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Media.Id,
		"name":             Media.Name,
		"social_media_url": Media.Social_media_url,
		"user_id":          Media.User_id,
		"created_at":       Media.Created_at,
	})
}

// GetAllSocialMedia godoc
// @Summary Get all social media
// @Description Get all social media in mygram
// @Tags social media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Media "Get all social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/getsocmeds [get]
func (m *MediaRepo) GetMedia(c *gin.Context) {
	Medias := []models.Media{}

	if err := m.DB.Debug().Find(&Medias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "can't find media",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"media": Medias,
	})
}

// GetSocialMedia godoc
// @Summary Get social media
// @Description Get social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.Media "Get social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/getsocmedbyid/{socialMediaId} [get]
func (m *MediaRepo) GetSocialMediabyId(c *gin.Context) {

	SocialMedia := models.Media{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := m.DB.First(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.Media "Update social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/updatesocmed/{socialMediaId} [put]
func (m *MediaRepo) UpdateMedia(c *gin.Context) {
	contentType := pkg.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("socialMediaId"))

	Media := models.Media{}
	OldMedia := models.Media{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.Updated_at = time.Now()
	Media.User_id = uint(userId)

	if err := m.DB.Debug().First(&OldMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Model(&OldMedia).Updates(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               OldMedia.Id,
		"name":             OldMedia.Name,
		"social_media_url": OldMedia.Social_media_url,
		"user_id":          OldMedia.User_id,
		"updated_at":       OldMedia.Updated_at,
	})

}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media identified by given ID
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {string} string "Delete social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/deletesocmed/{socialMediaId} [delete]
func (m *MediaRepo) DeleteMedia(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("socialMediaId"))
	Media := models.Media{}

	if err := m.DB.Debug().First(&Media, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Delete(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete media",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
