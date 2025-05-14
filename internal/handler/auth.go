package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // Ginフレームワークのコアパッケージ

	"golang-catch-up/internal/db"
	"golang-catch-up/internal/model"
	"golang-catch-up/internal/repository"
)

func LoginHandler(c *gin.Context) {
	var loginData model.LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Printf("Received login request: %s, %s\n", loginData.Username, loginData.Password)

	repo := repository.NewUserRepository(db.Conn)
	users, err := repo.GetAll(c.Request.Context())
	if err != nil {
		fmt.Printf("Error fetching users from DB: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー情報の取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
