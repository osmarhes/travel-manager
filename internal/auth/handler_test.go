package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/osmarhes/travel-manager/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestAuthHandler() *Handler {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to open test db")
	}
	db.AutoMigrate(&user.User{})
	repo := user.NewRepository(db)
	return NewHandler(repo)
}

func TestRegisterHandler(t *testing.T) {
	e := echo.New()
	handler := setupTestAuthHandler()

	payload := `{"name":"Test","email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "user registered")
	}
}

func TestLoginHandler(t *testing.T) {
	e := echo.New()
	handler := setupTestAuthHandler()

	regPayload := `{"name":"Test App","email":"login@example.com","password":"123456"}`
	req1 := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(regPayload))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	handler.Register(c1)

	loginPayload := `{"email":"login@example.com","password":"123456"}`
	req2 := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(loginPayload))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	if assert.NoError(t, handler.Login(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
		assert.Contains(t, rec2.Body.String(), "token")
	}
}
