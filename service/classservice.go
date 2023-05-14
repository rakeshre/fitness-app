package service

import (
	"fmt"
	"time"
)

func GetFormatedDate(dateString string) time.Time {
	date, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		fmt.Println("Error parsing date \n", dateString, err)
	}
	return date
}

func GetFormatedTime(timestring string) time.Time {
	date, err := time.Parse("15:04", timestring)

	if err != nil {
		fmt.Println("Error parsing date \n", timestring, err)
	}
	return date
}
func GetStringDate(ts time.Time) string {
	return ts.Format("2006-01-02")
}
func GetStringTime(ts time.Time) string {
	return ts.Format("15:04")
}
