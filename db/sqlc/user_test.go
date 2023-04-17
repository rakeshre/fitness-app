package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Name:           "J1",
		Hashedpassword: "Password123",
		Email:          "jaswanth@gmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Hashedpassword, user.Hashedpassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

//func TestGetUser(t *testing.T) {
//	user1 := createRandomUser(t)
//	user2, err := testQueries.GetUser(context.Background(), user1.Name)
//	require.NoError(t, err)
//	require.NotEmpty(t, user2)
//
//	require.Equal(t, user1.Name, user2.Name)
//	require.Equal(t, user1.Hashedpassword, user2.Hashedpassword)
//	require.Equal(t, user1.Email, user2.Email)
//	//require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
//}

//
//func TestUpdateUserOnlyFullName(t *testing.T) {
//	oldUser := createRandomUser(t)
//
//	newFullName := util.RandomOwner()
//	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
//		Username: oldUser.Username,
//		FullName: string{
//			String: newFullName,
//			Valid:  true,
//		},
//	})
//
//	require.NoError(t, err)
//	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
//	require.Equal(t, newFullName, updatedUser.FullName)
//	require.Equal(t, oldUser.Email, updatedUser.Email)
//	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
//}
//
//func TestUpdateUserOnlyEmail(t *testing.T) {
//	oldUser := createRandomUser(t)
//
//	newEmail := util.RandomEmail()
//	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
//		Username: oldUser.Username,
//		Email: string{
//			String: newEmail,
//			Valid:  true,
//		},
//	})
//
//	require.NoError(t, err)
//	require.NotEqual(t, oldUser.Email, updatedUser.Email)
//	require.Equal(t, newEmail, updatedUser.Email)
//	require.Equal(t, oldUser.FullName, updatedUser.FullName)
//	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
//}
//
//func TestUpdateUserOnlyPassword(t *testing.T) {
//	oldUser := createRandomUser(t)
//
//	newPassword := util.RandomString(6)
//	newHashedPassword, err := util.HashPassword(newPassword)
//	require.NoError(t, err)
//
//	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
//		Username: oldUser.Username,
//		HashedPassword: string{
//			String: newHashedPassword,
//			Valid:  true,
//		},
//	})
//
//	require.NoError(t, err)
//	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
//	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
//	require.Equal(t, oldUser.FullName, updatedUser.FullName)
//	require.Equal(t, oldUser.Email, updatedUser.Email)
//}
//
//func TestUpdateUserAllFields(t *testing.T) {
//	oldUser := createRandomUser(t)
//
//	newFullName := util.RandomOwner()
//	newEmail := util.RandomEmail()
//	newPassword := util.RandomString(6)
//	newHashedPassword, err := util.HashPassword(newPassword)
//	require.NoError(t, err)
//
//	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
//		Username: oldUser.Username,
//		FullName: string{
//			String: newFullName,
//			Valid:  true,
//		},
//		Email: string{
//			String: newEmail,
//			Valid:  true,
//		},
//		HashedPassword: string{
//			String: newHashedPassword,
//			Valid:  true,
//		},
//	})
//
//	require.NoError(t, err)
//	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
//	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
//	require.NotEqual(t, oldUser.Email, updatedUser.Email)
//	require.Equal(t, newEmail, updatedUser.Email)
//	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
//	require.Equal(t, newFullName, updatedUser.FullName)
//}
