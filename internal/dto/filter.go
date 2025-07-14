package dto

import "test/internal/entity"

type FilterRequest struct {
	BeginDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
	UserID      string `json:"user_id"`
	ServiceName string `json:"service_name"`
}

func (r *FilterRequest) ToEntity() (*entity.Filter, error) {
	begDate, endDate, err := parseDate(r.BeginDate, r.EndDate)
	if err != nil {
		return nil, err
	}

	return &entity.Filter{
		BeginDate:   *begDate,
		EndDate:     *endDate,
		UserID:      r.UserID,
		ServiceName: r.ServiceName,
	}, nil
}

type SumResponse struct {
	Sum int `json:"sum"`
}
