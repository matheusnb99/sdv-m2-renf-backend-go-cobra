package controllers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(Email string, Password string) string {

	if Email == "" {
		panic("Email is required")
	}

	if Password == "" {
		panic("Password is required")
	}



	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 10)

	if err != nil {
		panic("error hashing password")
	}

	

	// Create user
	user := models.User{ Email: Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		panic("error creating user")
	}

	return "User created successfully"

}


func Login(Email string, Password string) string {
	if Email == "" {
		panic("Email is required")
	}

	if Password == "" {
		panic("Password is required")
	}

	// Find user
	var user models.User
	result := initializers.DB.First(&user, "email = ?", Email)

	if result.Error != nil {
		panic("user not found")
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))

	if err != nil {
		panic("invalid password")
	}

	// return user logged
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		panic("error generating token")
	}

	return tokenString
}

