package service

import (
	db "Gym-backend/db/sqlc"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidateUser(ctx *gin.Context, Email string, Password string, store db.Store) (string, string, string) {
	objecttype := "user"
	if strings.Contains(Email, "@gym.com") {
		objecttype = "employee"
	}
	object := ""
	if objecttype == "employee" {
		employee, err := store.GetEmployeeFromEmail(ctx, Email)
		if err != nil {
			return "Error Reading Employee form DB", "", ""
		}
		if Password != employee.Hashedpassword {
			return "Employee Password Mismatch", "", ""
		}

		b, err := json.Marshal(employee)
		if err != nil {
			return "Error Marshalling Employee object", "", ""
		}
		object = string(b)
		return "", objecttype, object
	}

	user, err := store.GetUserFromEmail(ctx, Email)
	if err != nil {
		return "Error Reading User form DB", "", ""
	}
	if Password != user.Hashedpassword {
		return "Employee Password Mismatch", "", ""
	}
	b, err := json.Marshal(user)
	if err != nil {
		return "Error Marshalling User object", "", ""
	}
	object = string(b)
	return "", objecttype, object
}
