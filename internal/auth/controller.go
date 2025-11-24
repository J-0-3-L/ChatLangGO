package auth

import (
	"log"
	"net/http"
	"time"

	"go000/internal/config"
	"go000/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SecretJwt = []byte("SECRETO_PASS")

// Ruta register

// @Summary Registro de usuarios
// @Description Crea nuevos usuarios en el API
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.RegisterInput true "Datos del usuario"
// @Success 200 {object} auth.UserResponse
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func Register(c *gin.Context) {

	var input RegisterInput

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

	// c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado!!!"})
	// Crear respuesta segura
	response := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		//Password:  user.Password,
		AvatarURL: user.AvatarURL,
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary Inicio de sesión
// @Description Permite a un usuario autenticarse y obtener un token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body auth.LoginInput true "Credenciales del usuario"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Comparar contraseña
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Crear JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(SecretJwt)
	if err != nil {
		log.Println("ERROR_JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   tokenString,
	})
}
