package services

import (
	"fmt"
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
)

// Update date and time for header
func (s *ChartService) GetDateAndTime() models.DateAndTime {
	loc, _ := time.LoadLocation("Europe/Rome")
	now := time.Now().In(loc)

	dayName := models.Giorni[int(now.Weekday())]
	date := fmt.Sprintf("%d %s %d", now.Day(), models.Mesi[int(now.Month())], now.Year())
	timeStr := now.Format("15:04")

	return models.DateAndTime{
		DayName: dayName,
		Date:    date,
		Time:    timeStr,
	}
}
