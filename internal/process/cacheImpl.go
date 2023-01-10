package process

import (
	log "github.com/sirupsen/logrus"
	"simpleService/internal"
	"strings"
	"time"
)

type cache struct {
	cachedHolidays []*Holiday
}

func NewCacheImpl() (Cache, error) {
	c := &cache{}
	log.Info("Initializing cache: fetching data")
	c.cachedHolidays = []*Holiday{
		&Holiday{
			Name:        "Año nuevo",
			Type:        "Civil",
			PhoneNumber: "X",
			Extra:       nil,
		},
		&Holiday{
			Name:        "Independencia nacional",
			Type:        "Civil",
			PhoneNumber: "Y",
		},
		&Holiday{
			Name:        "Asunción de la virgen",
			Type:        "Religioso",
			PhoneNumber: "Z",
		},
	}
	newYear, _ := time.Parse(internal.Layout, "01-01")
	c.cachedHolidays[0].Date = newYear

	independenceDay, _ := time.Parse(internal.Layout, "09-18")
	c.cachedHolidays[1].Date = independenceDay

	assumptionDay, _ := time.Parse(internal.Layout, "08-15")
	c.cachedHolidays[2].Date = assumptionDay
	log.Info("Initializing cache: fetched data", c.cachedHolidays)
	return c, nil
}

func (c cache) GetHolidays(holidayType string, beginDate, endDate time.Time) (results []*Holiday) {
	log.Info("CacheImpl.GetHolidays has been invoked with params ", holidayType, beginDate, endDate)
	for _, holiday := range c.cachedHolidays {

		// Compare holiday type
		if strings.EqualFold(holiday.Type, holidayType) {
			results = append(results, holiday)
		}

		// TODO: Compare dates
	}
	log.Info("CacheImpl.GetHolidays is returning with results ", results)
	return
}

func IsBetween(date, beginDate, endDate time.Time) bool {
	return false
}
