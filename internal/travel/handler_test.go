package travel

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestTravelHandler() *Handler {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to open test db")
	}
	if err := db.AutoMigrate(&TravelRequest{}); err != nil {
		panic("failed to migrate")
	}
	repo := NewRepository(db)
	return NewHandler(repo)
}

func TestCreateTravelRequest(t *testing.T) {
	e := echo.New()
	handler := setupTestTravelHandler()

	payload := `{
        "requester": "Alice",
        "destination": "Recife",
        "departure": "2025-08-01T10:00:00Z",
        "return": "2025-08-05T18:00:00Z"
    }`

	req := httptest.NewRequest(http.MethodPost, "/travels", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userID", uint(1))

	if assert.NoError(t, handler.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res TravelRequest
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "Alice", res.Requester)
		assert.Equal(t, "solicitado", res.Status)
	}
}

func TestUpdateTravelStatus(t *testing.T) {
	e := echo.New()
	handler := setupTestTravelHandler()

	payload := `{
        "requester": "Osmar",
        "destination": "São Paulo",
        "departure": "2025-09-01T10:00:00Z",
        "return": "2025-09-05T18:00:00Z"
    }`
	req := httptest.NewRequest(http.MethodPost, "/travels", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userID", uint(1))
	handler.Create(c)

	updatePayload := `{"status":"aprovado"}`
	req2 := httptest.NewRequest(http.MethodPut, "/travels/1/status", strings.NewReader(updatePayload))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.Set("userID", uint(2))
	c2.SetParamNames("id")
	c2.SetParamValues("1")

	if assert.NoError(t, handler.UpdateStatus(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
		assert.Contains(t, rec2.Body.String(), "status updated")
	}
}
