package process

import "time"

type Cache interface {
	GetHolidays(holidayType string, beginDate, endDate time.Time) []*Holiday
}

type Holiday struct {
	Name        string
	Date        time.Time
	Type        string
	PhoneNumber string
	Extra       interface{}
}
