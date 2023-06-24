package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/asifDurran/Golang-crud/initializers"
	"github.com/asifDurran/Golang-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth (c *gin.Context){

	tokenString, err := c.Cookie("Authorization")
   // get req data from body
   if err != nil {
	c.AbortWithStatus(http.StatusUnauthorized)
   }
   // decode and validate it
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	} 
	return []byte(os.Getenv("SECRET")), nil
   })

   if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	 // check expire
	 if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
	 }
	 // Find the user with token sub
	 var user models.User
	
	 initializers.DB.Find(&user, claims["sub"])
	 if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	 }
	 // attache to req
	 c.Set("user", user)
	 // Continue
	 c.Next()

   }else {
	c.AbortWithStatus(http.StatusUnauthorized)
   }
}