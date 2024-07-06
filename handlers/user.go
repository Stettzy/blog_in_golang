package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Stettzy/blog_in_golang/pkg/user"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PasswordCompare(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateToken(subject string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func LoginUser(c echo.Context) error {
	var creds Credentials

	err := c.Bind(&creds)
	if err != nil {
		return echo.ErrInternalServerError
	}

	u, err := user.GetByEmail(creds.Email)
	if err != nil {
		return echo.ErrNotFound
	}

	match, err := PasswordCompare(u.Password, creds.Password)
	if err != nil {
		log.Printf("Password comparison error: %v", err)
	}

	if match {
		token, err := GenerateToken(u.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}

	return echo.ErrUnauthorized
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func RegisterUser(c echo.Context) error {
	var ur RegisterRequest

	if err := c.Bind(&ur); err != nil {
		return echo.ErrBadRequest
	}

	u := user.NewUser()
	u.AssignRegisterData(ur.Username, ur.Email, ur.Password)

	id, err := u.CreateUser()
	if err != nil {
		return echo.ErrInternalServerError
	}

	u.ID = id

	return c.JSON(201, u)
}

func getUserIDFromRequest(c echo.Context) (int, error) {
	var body struct {
		ID int `json:"id"`
	}
	if err := c.Bind(&body); err != nil {
		return 0, err
	}
	return body.ID, nil
}

func RemoveUser(c echo.Context) error {
	userID, err := getUserIDFromRequest(c)
	if err != nil {
		return echo.ErrBadRequest
	}

	u, err := user.GetById(userID)
	if err != nil {
		return echo.ErrNotFound
	}

	u.DeleteUser()

	return c.JSON(http.StatusOK, map[string]int{"deleted_user_id": userID})
}

func UpdateUser(c echo.Context) (int, error) {
	return 0, nil
}
