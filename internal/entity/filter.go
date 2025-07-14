package entity

import "time"

type Filter struct {
	BeginDate   time.Time
	EndDate     time.Time
	UserID      string
	ServiceName string
}
