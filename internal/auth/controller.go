package auth

import (
	"net/http"
	//"time"

	"go000/internal/config"
	"go000/internal/models"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SecretJwt = []byte("SECRETO_PASS")

// Ruta register

// @Summary Endpoint register
// @Description Registro de nuevos usuarios
// @Tags Register
// @Produce json
// Success 200 {object} map[string]string
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required, min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hasheo de constraseña
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPw),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario o email ya existe"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado!!!"})
}
