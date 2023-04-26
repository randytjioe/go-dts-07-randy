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

type PhotoRepo struct {
	DB *gorm.DB
}

// GetAllPhotos godoc
// @Summary Get all photos
// @Description Get all existing photos
// @Tags photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Photo{} "Get all photos success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photos Not Found"
// @Router /photo/getphotos [get]
func (p *PhotoRepo) GetPhoto(c *gin.Context) {
	Photos := []models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Preload("User").Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})

}

// GetPhoto godoc
// @Summary Get photo
// @Description Get photo by ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Get photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/getphotobyid/{photoId} [get]
func (p *PhotoRepo) GetPhotoById(c *gin.Context) {

	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := p.DB.First(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// CreatePhoto godoc
// @Summary Create photo
// @Description Create photo to post in mygram
// @Tags photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201 {object} models.Photo "Create photo success"
// @Failure 401 "Unauthorized"
// @Router /photo/uploadphoto [post]
func (p *PhotoRepo) UploadPhoto(c *gin.Context) {
	Photo := models.Photo{}
	contextType := pkg.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_id = uint(userId)
	Photo.Created_at = time.Now()
	Photo.Updated_at = time.Now()

	if err := p.DB.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to updload photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.User_id,
		"created_at": Photo.Created_at,
	})
}

// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Update photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/updatephoto/{photoId} [put]
func (p *PhotoRepo) UpdatePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"]

	contextType := pkg.GetContentType(c)
	Photo := models.Photo{}
	OldPhoto := models.Photo{}

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.Updated_at = time.Now()
	Photo.User_id = uint(userId.(float64))

	if err := p.DB.Debug().First(&OldPhoto, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Model(&OldPhoto).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldPhoto.Id,
		"title":      OldPhoto.Title,
		"caption":    OldPhoto.Caption,
		"photo_url":  OldPhoto.Photo_url,
		"user_id":    OldPhoto.User_id,
		"updated_at": OldPhoto.Updated_at,
	})

}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {string} string "Delete photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photo/deletephoto/{photoId} [delete]
func (p *PhotoRepo) DeletePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := p.DB.Debug().First(&Photo, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Delete(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete photo",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been succesfully deleted",
	})
}
