package services

import "github.com/chiaf1/solar-frontend/internal/models"

// Returns chart data for today's chart
func GetTodayChart() models.ChartData {
	todayData := models.ChartData{
		Labels: []string{
			"00:00", "02:00", "04:00", "06:00",
			"08:00", "10:00", "12:00", "14:00",
			"16:00", "18:00", "20:00", "22:00",
		},
		Values: []float64{
			0, 0, 0, 0.5,
			1.2, 2.8, 4.1, 3.9,
			2.6, 1.1, 0.3, 1,
		},
	}
	return todayData
}

// Returns chart data for history's charts
func GetHistoryCharts() map[string]models.ChartData {
	hystoryData := map[string]models.ChartData{
		"chart-yesterday": {
			Labels: []string{"00:00", "04:00", "08:00", "12:00", "16:00", "20:00"},
			Values: []float64{0.5, 1.2, 2.3, 1.8, 0.9, 0.3},
		},
		"chart-minus-2": {
			Labels: []string{"00:00", "06:00", "12:00", "18:00"},
			Values: []float64{0.2, 1.1, 2.0, 0.8},
		},
		"chart-minus-3": {
			Labels: []string{"00:00", "06:00", "12:00", "18:00"},
			Values: []float64{0, 0.9, 1.7, 0.4},
		},
		"chart-minus-4": {
			Labels: []string{"00:00", "06:00", "12:00", "18:00"},
			Values: []float64{0, 0.9, 1.7, 0.4},
		},
		"chart-minus-5": {
			Labels: []string{"00:00", "06:00", "12:00", "18:00"},
			Values: []float64{0, 0.9, 1.7, 0.4},
		},
		"chart-minus-6": {
			Labels: []string{"00:00", "06:00", "12:00", "18:00"},
			Values: []float64{0, 0.9, 1.7, 0.4},
		},
	}
	return hystoryData
}
