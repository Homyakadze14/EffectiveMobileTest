package dto

import (
	"test/internal/entity"
)

type SubscritpionRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	BeginDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

func (r *SubscritpionRequest) ToEntity() (*entity.Subscritpion, error) {
	begDate, endDate, err := parseDate(r.BeginDate, r.EndDate)
	if err != nil {
		return nil, err
	}

	return &entity.Subscritpion{
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      r.UserID,
		BeginDate:   *begDate,
		EndDate:     endDate,
	}, nil
}

type SubscritpionResponse struct {
	ID          int     `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       uint    `json:"price"`
	UserID      string  `json:"user_id"`
	BeginDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}

func ToSubscriptionResponse(r *entity.Subscritpion) SubscritpionResponse {
	layout := "01-2006"
	var endDate *string
	if r.EndDate != nil {
		res := r.EndDate.Format(layout)
		endDate = &res
	}

	return SubscritpionResponse{
		ID:          r.ID,
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      r.UserID,
		BeginDate:   r.BeginDate.Format(layout),
		EndDate:     endDate,
	}
}

type SubscritpionsResponse struct {
	Subscriptions []SubscritpionResponse `json:"subscriptions"`
}

func ConvertToSubscriptionsResponse(subs []entity.Subscritpion) SubscritpionsResponse {
	res := make([]SubscritpionResponse, len(subs))
	for i, el := range subs {
		res[i] = ToSubscriptionResponse(&el)
	}

	return SubscritpionsResponse{
		Subscriptions: res,
	}
}
