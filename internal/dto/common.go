package dto

import (
	"test/internal/common"
	"time"
)

func parseDate(beg, end string) (*time.Time, *time.Time, error) {
	layout := "01-2006"
	begDate, err := time.Parse(layout, beg)
	if err != nil {
		return nil, nil, common.ErrDateFormat
	}

	var endDate *time.Time
	if end != "" {
		res, err := time.Parse(layout, end)
		if err != nil {
			return nil, nil, common.ErrDateFormat
		}
		endDate = &res

		if begDate.After(*endDate) {
			return nil, nil, common.ErrBeginDateAfterEndDate
		}
	}

	return &begDate, endDate, nil
}
