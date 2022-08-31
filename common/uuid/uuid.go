package myUUID

import "github.com/google/uuid"

func GenID() (id string) {
	id = uuid.New().String()
	return id
}
