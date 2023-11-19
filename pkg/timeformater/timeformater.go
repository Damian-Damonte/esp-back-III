package timeformater

import (
	"errors"
	"time"
)

const (
	timeFormat = "2006-01-02"
)

var (
	ErrParceStringToTime = errors.New("error al convertir string a time.time")
)

func StringToTime(timeString string) (*time.Time, error) {
	var fecha time.Time

  fecha, err := time.Parse(timeFormat, timeString)
  if err != nil {
    return nil, ErrParceStringToTime
  }

	return &fecha, nil
}

func TimeToString(time time.Time) (string) {
  return time.Format(timeFormat)
}