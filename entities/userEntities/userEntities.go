package userEntities

import (
	"fmt"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (user User) ToString() string {
	return fmt.Sprintf("Id: %s\nName: %s\nAddress: %s\n", user.Id, user.Name, user.Address)
}
