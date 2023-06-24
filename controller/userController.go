package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/asifDurran/Golang-crud/initializers"
	"github.com/asifDurran/Golang-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



func CreateUser(c *gin.Context){
	// var body struct {
	// 	Email    string
	// 	Password string
	// }
   var body body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash password
	has, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create uer 
	user := models.User{Email: body.Email, Password: string(has)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create use",
		})
		return
	}
	// response

	c.JSON(http.StatusOK, gin.H{
		"User created : ": user,
	})
}

// Login function 
func LoginUser(c *gin.Context){
	var body body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Failed to read the body",
		})
		return
	}
	// Requested user

	var user models.User
	initializers.DB.First(&user, "email= ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})
		return
	}
	// Generate token 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Failed to create token",
		})

		return
	}
	// Send it back 
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 5, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Token generated": tokenString,
	})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"Return User : ": user,
	})
}