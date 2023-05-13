package service

import (
	db "Gym-backend/db/sqlc"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetPreviousUserActivityRecord(ctx *gin.Context, store db.Store, userid int64) (db.Activityrecord, error) {
	previous, err := store.GetLatestActivityRecord(ctx, userid)
	if err != nil {
		return db.Activityrecord{}, err
	}
	return previous, nil
}

func CreateEndActivityRecord(ctx *gin.Context, store db.Store, arg db.CreateActivityRecordsParams) (db.Useractivity, error) {
	previous, err := GetPreviousUserActivityRecord(ctx, store, arg.Userid)
	if err != nil {
		return db.Useractivity{}, err
	}
	if previous.Type != 1 {
		return db.Useractivity{}, errors.New("No Previous Checkin")
	}
	current, err := store.CreateActivityRecords(ctx, arg)
	if err != nil {
		return db.Useractivity{}, err
	}
	activity, err := CreateUserActivity(ctx, store, previous, current)
	if err != nil {
		return db.Useractivity{}, err
	}
	return activity, nil

}

func CreateUserActivity(ctx *gin.Context, store db.Store, checkin db.Activityrecord, checkout db.Activityrecord) (db.Useractivity, error) {
	arg := db.CreateUserActivityParams{
		Start:      checkin.Time,
		End:        checkout.Time,
		Userid:     checkin.Userid,
		Locationid: checkout.Locationid,
		Deviceid:   checkout.Deviceid,
	}

	activity, err := store.CreateUserActivity(ctx, arg)
	if err != nil {

		return db.Useractivity{}, err
	}
	return activity, nil

}
