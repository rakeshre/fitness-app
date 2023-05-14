package service

import (
	db "Gym-backend/db/sqlc"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetPreviousRecord(ctx *gin.Context, store db.Store, userid int64) (db.Checkinrecord, error) {
	previous, err := store.GetLatestCheckinRecord(ctx, userid)
	if err != nil {
		return db.Checkinrecord{}, err
	}
	return previous, nil
}

func CreateCheckOutRecord(ctx *gin.Context, store db.Store, arg db.CreateCheckinRecordsParams) (db.Checkinactivity, error) {
	previous, err := GetPreviousRecord(ctx, store, arg.Userid)
	if err != nil {
		return db.Checkinactivity{}, err
	}
	if previous.Type != 1 {
		return db.Checkinactivity{}, errors.New("No Previous Checkin")
	}
	current, err := store.CreateCheckinRecords(ctx, arg)
	if err != nil {
		return db.Checkinactivity{}, err
	}
	activity, err := CreateCheckinActivity(ctx, store, previous, current)
	if err != nil {
		return db.Checkinactivity{}, err
	}
	return activity, nil

}

func CreateCheckInRecord(ctx *gin.Context, store db.Store, arg db.CreateCheckinRecordsParams) (db.Checkinrecord, error) {
	previous, err := GetPreviousRecord(ctx, store, arg.Userid)
	if err != nil {
		return db.Checkinrecord{}, err
	}
	if previous.Type == 1 {
		return db.Checkinrecord{}, errors.New("Already Checkdin")
	}
	current, err := store.CreateCheckinRecords(ctx, arg)
	if err != nil {
		return db.Checkinrecord{}, err
	}

	return current, nil

}

func CreateCheckinActivity(ctx *gin.Context, store db.Store, checkin db.Checkinrecord, checkout db.Checkinrecord) (db.Checkinactivity, error) {
	arg := db.CreateCheckinActivityParams{
		Checkin:    checkin.Time,
		Checkout:   checkout.Time,
		Userid:     checkin.Userid,
		Employeeid: checkout.Employeeid,
		Locationid: checkout.Locationid,
	}

	activity, err := store.CreateCheckinActivity(ctx, arg)
	if err != nil {

		return db.Checkinactivity{}, err
	}
	return activity, nil

}
