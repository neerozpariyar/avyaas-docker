package utils

import (
	"errors"
	"time"
)

func ParseStringToTime(strTime string) (*time.Time, error) {
	var et time.Time
	var err error

	if et, err = time.Parse(time.RFC3339, strTime); err != nil {
		return nil, errors.New("error parsing invalid UTC time")
	}

	return &et, nil
}
