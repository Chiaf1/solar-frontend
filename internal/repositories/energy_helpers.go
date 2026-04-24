package repositories

import (
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
)

// Convert EnergyAPIPoint slice into chartdata
func mapAPIResponseToChartData(apiResp []EnergyAPIPoint) models.ChartData {
	loc, _ := time.LoadLocation("Europe/Rome")

	step := detectStep(apiResp, 5*time.Minute, loc)

	timeline := buildDailyTimeline(step, loc)
	index := indexAPIPoints(apiResp, step, loc)

	labels := make([]string, 0, len(timeline))
	production := make([]*float64, 0, len(timeline))
	consumption := make([]*float64, 0, len(timeline))

	for _, t := range timeline {
		labels = append(labels, t.Format("15:04"))

		if point, ok := index[t]; ok {
			production = append(production, &point.Production)
			consumption = append(consumption, &point.Consumption)
		} else {
			production = append(production, nil)
			consumption = append(consumption, nil)
		}

	}

	return models.ChartData{
		Labels:      labels,
		Production:  production,
		Consumption: consumption,
	}
}

// Build daily timeline returns the labels for a comlpete day with a set step
func buildDailyTimeline(step time.Duration, loc *time.Location) []time.Time {
	start := time.Date(0, 1, 1, 0, 0, 0, 0, loc)
	end := start.Add(24 * time.Hour)

	var timeline []time.Time
	for t := start; !t.After(end); t = t.Add(step) {
		timeline = append(timeline, t)
	}
	return timeline
}

// Creates a map for all energy points with the time stamp as key
func indexAPIPoints(apiPoints []EnergyAPIPoint, step time.Duration, loc *time.Location) map[time.Time]EnergyAPIPoint {
	index := make(map[time.Time]EnergyAPIPoint)

	for _, p := range apiPoints {
		t, err := time.Parse(time.RFC3339, p.Time)
		if err != nil {
			continue
		}

		bucket := normalizeToBacket(t, step, loc)
		index[bucket] = p
	}
	return index
}

// Detect step duration in a slice of EnergyAPIPoint
func detectStep(apiResp []EnergyAPIPoint, fallback time.Duration, loc *time.Location) time.Duration {
	if len(apiResp) < 2 {
		return fallback
	}

	t1, err1 := time.Parse(time.RFC3339, apiResp[0].Time)
	t2, err2 := time.Parse(time.RFC3339, apiResp[1].Time)

	if err1 != nil || err2 != nil {
		return fallback
	}

	// Conv to local time before step calculation
	t1 = t1.In(loc)
	t2 = t2.In(loc)

	step := t2.Sub(t1)

	// Check step value
	if step <= 0 || step > time.Hour {
		return fallback
	}

	return step
}

// Normalize timestamp to time stamp starting from midnight with fixed step
func normalizeToBacket(ts time.Time, step time.Duration, loc *time.Location) time.Time {
	local := ts.In(loc)

	// Midnight in local time of the smae day
	startOfDay := time.Date(0, 1, 1, 0, 0, 0, 0, loc)

	elapsed := time.Duration(local.Hour())*time.Hour +
		time.Duration(local.Minute())*time.Minute +
		time.Duration(local.Second())*time.Second +
		time.Duration(local.Nanosecond())

	bucketIndex := elapsed / step

	return startOfDay.Add(bucketIndex * step)
}

// Helper function returns empty chart data
func emptyDailyChart(step time.Duration) models.ChartData {
	loc, _ := time.LoadLocation("Europe/Rome")
	timeline := buildDailyTimeline(step, loc)

	labels := make([]string, 0, len(timeline))
	production := make([]*float64, 0, len(timeline))
	consumption := make([]*float64, 0, len(timeline))

	for _, t := range timeline {
		labels = append(labels, t.Format("15:04"))
		production = append(production, nil)
		consumption = append(consumption, nil)
	}

	return models.ChartData{
		Labels:      labels,
		Production:  production,
		Consumption: consumption,
	}
}
