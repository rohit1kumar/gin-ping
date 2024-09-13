package controllers

import (
	"errors"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohit1kumar/pgo/config"
	"github.com/rohit1kumar/pgo/models"
	"gorm.io/gorm"
)

// @Summary Create a new post
// @Description Create a new blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param body body object true "Post object"
// @Success 201 {object} models.Post "Successfully created post"
// @Failure 400 {object} object "Bad request"
// @Failure 500 {object} object "Internal server error"
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Bad request, check your body json",
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := config.DB.Create(&post)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// @Summary Get paginated posts
// @Description Retrieve a paginated list of posts
// @Tags posts
// @Accept json
// @Produce json
// @Param limit query int false "Number of items per page" default(10) maximum(20)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} map[string]interface{} "Successfully retrieved posts"
// @Failure 404 {object} map[string]interface{} "Page not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if limit > 20 {
		limit = 20
	}

	var totalCount int64
	if err := config.DB.Model(&models.Post{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "Failed to get total count of posts",
		})
		return
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	// Check if requested page is out of range
	if page > totalPages {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "Page not found",
		})
		return
	}

	offset := (page - 1) * limit

	var posts []models.Post
	config.DB.Limit(limit).Offset(offset).Find(&posts)

	paginationInfo := struct {
		CurrentPage int   `json:"current_page"`
		TotalPages  int   `json:"total_pages"`
		TotalCount  int64 `json:"total_count"`
		Limit       int   `json:"limit"`
	}{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalCount:  totalCount,
		Limit:       limit,
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "posts fetched",
		"data": gin.H{
			"posts": posts,
			"meta":  paginationInfo,
		},
	})
}

// @Summary Get a post by ID
// @Description Retrieve a specific post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} object "Successfully retrieved post"
// @Failure 404 {object} object "Post not found"
// @Router /posts/{id} [get]
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := config.DB.First(&post, id)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Got it",
		"data":  post,
	})
}

// @Summary Update a post
// @Description Update an existing post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param body body object true "Updated post object"
// @Success 200 {object} object "Successfully updated post"
// @Failure 400 {object} object "Bad request"
// @Failure 404 {object} object "Post not found"
// @Failure 500 {object} object "Internal server error"
// @Router /posts/{id} [put]
func UpdatePostById(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  *string `json:"body"`
		Title *string `json:"title"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Bad request, check your body json",
		})
		return
	}

	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "Post not found",
		})
		return
	}

	if body.Title != nil {
		post.Title = *body.Title
	}
	if body.Body != nil {
		post.Body = *body.Body
	}

	result = config.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "Failed to update post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Post updated successfully",
	})
}

// @Summary Delete a post
// @Description Delete an existing post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} object "Successfully deleted post"
// @Failure 404 {object} object "Post not found"
// @Failure 500 {object} object "Internal server error"
// @Router /posts/{id} [delete]
func DeletePostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "Post not found",
		})
		return
	}

	result = config.DB.Delete(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "Failed to delete post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Post deleted",
	})
}
