package controllers

import (
    "myapp/app"
    "myapp/database"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
    var photo app.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    photo.UserID = userID.(uint)
    photo.CreatedAt = time.Now()
    photo.UpdatedAt = time.Now()

    if err := database.DB.Create(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, photo)
}

func GetPhotos(c *gin.Context) {
    var photos []app.Photo
    if err := database.DB.Find(&photos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
    var photo app.Photo
    id := c.Param("photoId")

    if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    photo.UpdatedAt = time.Now()
    if err := database.DB.Save(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {
    var photo app.Photo
    id := c.Param("photoId")

    if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    if err := database.DB.Delete(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
