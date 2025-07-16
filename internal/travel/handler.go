package travel

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var req TravelRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	req.UserID = userID
	req.Status = "solicitado"

	if err := h.Repo.Create(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create request"})
	}

	return c.JSON(http.StatusCreated, req)
}

func (h *Handler) GetByID(c echo.Context) error {
	userID := c.Get("userID").(uint)
	id, _ := strconv.Atoi(c.Param("id"))

	req, err := h.Repo.FindByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "request not found"})
	}

	return c.JSON(http.StatusOK, req)
}

func (h *Handler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)

	filters := make(map[string]interface{})
	if status := c.QueryParam("status"); status != "" {
		filters["status"] = status
	}
	if destination := c.QueryParam("destination"); destination != "" {
		filters["destination"] = destination
	}

	if from := c.QueryParam("from"); from != "" {
		t, _ := time.Parse("2006-01-02", from)
		filters["from"] = t
	}
	if to := c.QueryParam("to"); to != "" {
		t, _ := time.Parse("2006-01-02", to)
		filters["to"] = t
	}

	list, err := h.Repo.List(userID, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to list"})
	}

	return c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateStatus(c echo.Context) error {
	userID := c.Get("userID").(uint)
	id, _ := strconv.Atoi(c.Param("id"))

	var body struct {
		Status string `json:"status"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	req, err := h.Repo.FindByID(uint(id), userID)
	if err == nil && req.UserID == userID {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "cannot update own request"})
	}

	if err := h.Repo.UpdateStatus(uint(id), body.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "update failed"})
	}

	log.Printf("Notificação: Pedido %d alterado para status: %s", id, body.Status)

	return c.JSON(http.StatusOK, echo.Map{"message": "status updated"})
}
