package services

import "time"

// Returns current cached temperature value
func (s *ChartService) GetTemperature() (float64, error) {
	s.weatherMu.Lock()
	defer s.weatherMu.Unlock()

	// Check if cached value is valid and returns it
	if s.weatherCache != nil && time.Since(s.weatherCache.timestamp) < s.weatherCacheTTL {
		return s.weatherCache.temperature, nil
	}

	// If cache not valid call the api
	lat := 45.36286504575867
	lon := 10.158195263916129
	val, err := s.weatherRepo.GetCurrentTemperature(lat, lon)
	if err != nil {
		// fallback if valid cache
		if s.weatherCache != nil {
			return s.weatherCache.temperature, nil
		}
		return 0, nil
	}

	// update cache
	s.weatherCache = &cacheWeather{
		temperature: val,
		timestamp:   time.Now(),
	}

	return val, nil
}
