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

type CommentRepo struct {
	DB *gorm.DB
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Get all comments in mygram
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Comment "Get all comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comments Not Found"
// @Router /comments/getcomments [get]
func (o *CommentRepo) GetComment(c *gin.Context) {
	Comments := []models.Comment{}

	if err := o.DB.Debug().Preload("User").Preload("Photo").Find(&Comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}

// GetComment godoc
// @Summary Get comment
// @Description Get comment identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {object} models.Comment "Get comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment by Id Not Found"
// @Router /comments/getcommentbyid/{commentId} [get]
func (o *CommentRepo) GetCommentById(c *gin.Context) {

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := o.DB.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// CreateComment godoc
// @Summary Create comment
// @Description Create comment for photo identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201 {object} models.Comment "Create comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /comments/uploadcomment [post]
func (o *CommentRepo) UploadComment(c *gin.Context) {
	contentType := pkg.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	if err := o.DB.Debug().Find(&models.Comment{}, Comment.Photo_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}
	Comment.User_id = uint(userId)
	Comment.Created_at = time.Now()
	Comment.Updated_at = time.Now()

	if err := o.DB.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":   "failed to upload comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.Photo_id,
		"user_id":    Comment.User_id,
		"created_at": Comment.Created_at,
	})
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {object} models.Comment "Update comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comments/updatecomment/{commentId} [put]
func (o *CommentRepo) UpdateComment(c *gin.Context) {
	contentType := pkg.GetContentType(c)
	Comment := models.Comment{}
	OldComment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.User_id = uint(userId)
	Comment.Updated_at = time.Now()

	if err := o.DB.Debug().First(&OldComment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Comment not found",
			"messsage": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Model(&OldComment).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldComment.Id,
		"message":    OldComment.Message,
		"user_id":    OldComment.User_id,
		"updated_at": OldComment.Updated_at,
	})
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment identified by given ID
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {string} string "Delete comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comments/deletecomment/{commentId} [delete]
func (o *CommentRepo) DeleteComment(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	if err := o.DB.Debug().First(&Comment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Delete(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})

}
