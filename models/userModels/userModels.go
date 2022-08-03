package userModels

import (
	"errors"
	"go-api/entities/userEntities"
)

var listUser = []*userEntities.User{
	{Id: "1-kenzo", Name: "Kenzo Hashima", Address: "Tokyo - Japan"},
	{Id: "2-hinata", Name: "Hinata Yuri", Address: "Tokyo - Japan"},
	{Id: "3-yumi", Name: "Yumi Sakura", Address: "Tokyo - Japan"},
}

func CreateUser(user *userEntities.User) bool {
	if user.Id != "" && user.Name != "" && user.Address != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return false
}

func UpdateUser(eUser *userEntities.User) bool {
	for index, user := range listUser {
		if user.Id == eUser.Id {
			listUser[index] = eUser
			return true
		}
	}
	return false
}

func FindUser(id string) (*userEntities.User, error) {
	for _, user := range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("User doesn't exist")
}

func DeleteUser(id string) bool {
	for index, user := range listUser {
		if user.Id == id {
			copy(listUser[index:], listUser[index+1:])
			listUser[len(listUser)-1] = &userEntities.User{}
			listUser = listUser[:len(listUser)-1]
			return true
		}
	}
	return false
}

func GetAllUser() []*userEntities.User {
	return listUser
}
