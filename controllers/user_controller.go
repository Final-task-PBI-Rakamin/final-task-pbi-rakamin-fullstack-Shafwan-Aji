package controllers

import (
    "myapp/app"
    "myapp/database"
    "myapp/helpers"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func RegisterUser(c *gin.Context) {
    var user app.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
    var user app.User
    var input app.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := helpers.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
    var user app.User
    id := c.Param("userId")

    if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.UpdatedAt = time.Now()
    if err := database.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    var user app.User
    id := c.Param("userId")

    if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
