package auth

import (
	"log"
	"net/http"

	"travel-manager/internal/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Repo user.Repository
}

func NewHandler(repo user.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Register(c echo.Context) error {
	var u user.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if u.Email == "" || u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "email and password are required"})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error hashing password"})
	}

	u.Password = string(hashed)
	if err := h.Repo.Create(&u); err != nil {
		log.Printf("Error creating user: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not create user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "user registered"})
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(c echo.Context) error {
	var creds Credentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if creds.Email == "" || creds.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "email and password are required"})
	}

	user, err := h.Repo.FindByEmail(creds.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
