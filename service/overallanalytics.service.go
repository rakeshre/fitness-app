package service

import (
	db "Gym-backend/db/sqlc"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type overallAnalyticsResponse struct {
	Previous int64 `json:"previous"`
	Current  int64 `json:"current"`
}

func GetOverallAnalytics(ctx *gin.Context, store db.Store, querytype int) (string, error) {

	switch querytype {
	case 1:
		{
			res, err := GetOverallAnalyticsDaily(ctx, store)
			if err != nil {
				return "", err
			}
			return res, err
		}
	case 2:
		{
			res, err := GetOverallAnalyticsWeekly(ctx, store)
			if err != nil {
				return "", err
			}
			return res, err
		}
	case 3:
		{
			res, err := GetOverallAnalyticsMonthly(ctx, store)
			if err != nil {
				return "", err
			}
			return res, err
		}
	}
	return "Invalid Type", nil
}

func GetOverallAnalyticsDaily(ctx *gin.Context, store db.Store) (string, error) {

	currentWeekCheckinCount, err := store.GetCurrentWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekCheckinCount, err := store.GetPreviousWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}

	response := []overallAnalyticsResponse{}

	checkin := overallAnalyticsResponse{
		Previous: previousWeekCheckinCount,
		Current:  currentWeekCheckinCount,
	}
	currentWeekClassEnrolmentCount, err := store.GetCurrentWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekClassEnrolmentCount, err := store.GetPreviousWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}

	enrolment := overallAnalyticsResponse{
		Previous: previousWeekClassEnrolmentCount,
		Current:  currentWeekClassEnrolmentCount,
	}

	currentWeekMembershipsCount, err := store.GetCurrentWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekMembershipsCount, err := store.GetPreviousWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}

	member := overallAnalyticsResponse{
		Previous: previousWeekMembershipsCount,
		Current:  currentWeekMembershipsCount,
	}

	response = append(response, checkin, member, enrolment)

	responseString, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(responseString), nil
}

func GetOverallAnalyticsWeekly(ctx *gin.Context, store db.Store) (string, error) {

	currentWeekCheckinCount, err := store.GetCurrentWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekCheckinCount, err := store.GetPreviousWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}

	response := []overallAnalyticsResponse{}

	checkin := overallAnalyticsResponse{
		Previous: previousWeekCheckinCount,
		Current:  currentWeekCheckinCount,
	}
	currentWeekClassEnrolmentCount, err := store.GetCurrentWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekClassEnrolmentCount, err := store.GetPreviousWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}

	enrolment := overallAnalyticsResponse{
		Previous: previousWeekClassEnrolmentCount,
		Current:  currentWeekClassEnrolmentCount,
	}

	currentWeekMembershipsCount, err := store.GetCurrentWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekMembershipsCount, err := store.GetPreviousWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}

	member := overallAnalyticsResponse{
		Previous: previousWeekMembershipsCount,
		Current:  currentWeekMembershipsCount,
	}

	response = append(response, enrolment, member, checkin)

	responseString, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(responseString), nil
}

func GetOverallAnalyticsMonthly(ctx *gin.Context, store db.Store) (string, error) {

	currentWeekCheckinCount, err := store.GetCurrentWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekCheckinCount, err := store.GetPreviousWeekCheckinCount(ctx)
	if err != nil {
		return "", err
	}

	response := []overallAnalyticsResponse{}

	checkin := overallAnalyticsResponse{
		Previous: previousWeekCheckinCount,
		Current:  currentWeekCheckinCount,
	}
	currentWeekClassEnrolmentCount, err := store.GetCurrentWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekClassEnrolmentCount, err := store.GetPreviousWeekClassEnrolmentCount(ctx)
	if err != nil {
		return "", err
	}

	enrolment := overallAnalyticsResponse{
		Previous: previousWeekClassEnrolmentCount,
		Current:  currentWeekClassEnrolmentCount,
	}

	currentWeekMembershipsCount, err := store.GetCurrentWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}
	previousWeekMembershipsCount, err := store.GetPreviousWeekMembershipsCount(ctx)
	if err != nil {
		return "", err
	}

	member := overallAnalyticsResponse{
		Previous: previousWeekMembershipsCount,
		Current:  currentWeekMembershipsCount,
	}

	response = append(response, enrolment, member, checkin)

	responseString, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(responseString), nil
}
