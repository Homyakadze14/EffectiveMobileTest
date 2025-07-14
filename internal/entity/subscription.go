package entity

import "time"

type Subscritpion struct {
	ID          int
	ServiceName string
	Price       uint
	UserID      string
	BeginDate   time.Time
	EndDate     *time.Time
}
