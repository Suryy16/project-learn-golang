package todoController

import (
	"encoding/json"
	"net/http"
	"todo-list-api/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodos(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	user := userInterface.(models.User)

	var todos []models.Todo

	if err := models.DB.Where("user_id = ?", user.ID).Find(&todos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return empty array instead of null
	if todos == nil {
		todos = []models.Todo{}
	}

	models.DB.Where("user_id=?", user.ID).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func GetTodo(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	var todo models.Todo

	id := c.Param("id")

	if err := models.DB.Where("user_id=?", user.ID).First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "terjadi kesalahan pada server"})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateTodo(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todo.UserID = user.ID
	models.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"todo": todo})

}

func UpdateTodo(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	var todo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&todo).Where("id= ? AND user_id=?", id, user.ID).Updates(&todo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "tidak dapat memperbarui data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

func DeleteTodo(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	var todo models.Todo

	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Where("user_id=?", user.ID).Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}

func ToggleTodo(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)
	var todo models.Todo

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if err := models.DB.Where("id=? AND user_id=?", id, user.ID).First(&todo).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "terjadi kesalahan pada server"})
			return
		}
	}

	todo.Completed = !todo.Completed
	models.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "todo completed status changed", "status": todo.Completed})
}
