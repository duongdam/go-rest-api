package userModels

import (
	"errors"
	"go-api/entities/userEntities"
)

var listUser = []*userEntities.User{
	{Id: "1-duong", Name: "Duong", Address: "Nga Thuy - Nga Son - Thanh Hoa"},
	{Id: "2-phi", Name: "Phi", Address: "Nga Trung - Nga Son - Thanh Hoa"},
	{Id: "3-thanh", Name: "Thanh", Address: "TT Nga Son - Thanh Hoa"},
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
