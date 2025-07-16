package travel

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(req *TravelRequest) error
	FindByID(id uint, userID uint) (*TravelRequest, error)
	List(userID uint, filters map[string]interface{}) ([]TravelRequest, error)
	UpdateStatus(id uint, newStatus string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(req *TravelRequest) error {
	return r.db.Create(req).Error
}

func (r *repository) FindByID(id uint, userID uint) (*TravelRequest, error) {
	var req TravelRequest
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&req).Error
	return &req, err
}

func (r *repository) List(userID uint, filters map[string]interface{}) ([]TravelRequest, error) {
	var requests []TravelRequest
	query := r.db.Where("user_id = ?", userID)

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if dest, ok := filters["destination"]; ok {
		query = query.Where("destination ILIKE ?", "%"+dest.(string)+"%")
	}

	if from, ok := filters["from"]; ok {
		query = query.Where("departure >= ?", from.(time.Time))
	}
	if to, ok := filters["to"]; ok {
		query = query.Where("return <= ?", to.(time.Time))
	}

	err := query.Order("created_at DESC").Find(&requests).Error
	return requests, err
}

func (r *repository) UpdateStatus(id uint, newStatus string) error {
	return r.db.Model(&TravelRequest{}).Where("id = ?", id).Update("status", newStatus).Error
}
