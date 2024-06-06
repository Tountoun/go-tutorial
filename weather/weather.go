package weather

import (
  "log"
  owm "github.com/briandowns/openweathermap"
)

func Current(unit string, language string, apiKey string) *owm.CurrentWeatherData {
  w, err := owm.NewCurrent(unit, language, apiKey)

  if err != nil {
    log.Fatalln(err)
  }
  
  return w
}
