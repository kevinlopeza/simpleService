package cache

import "time"

// Cache represents a caching implementation.
// It's been defined as an interface to hide implementation details.
type Cache interface {
	GetHolidays(holidayType string, beginDate, endDate time.Time) []*Holiday
}

// A Cache implementation returns a slice of Holidays when queried
type Holiday struct {
	Name        string
	Date        time.Time
	Type        string
	PhoneNumber string
	Extra       interface{}
}
