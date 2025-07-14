package common

import "errors"

var (
	ErrSubscriptionNotFound  = errors.New("subscription not found")
	ErrDateFormat            = errors.New("wrong date format")
	ErrBeginDateAfterEndDate = errors.New("the start date cannot be later than the end date")
	ErrBadURL                = errors.New("bad url")
	ErrBadType               = errors.New("bad type")
)
