package initModel

import (
	myUUID "golang-api-demo/common/uuid"
	model "golang-api-demo/modules/auth/struct"
	"time"
)

func Register(register *model.Register) map[string]interface{} {
	id := myUUID.GenID()
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	doc := make(map[string]interface{})

	doc["id"] = id
	doc["email"] = register.Email
	doc["phoneNumber"] = register.PhoneNumber
	doc["password"] = register.Password
	doc["createdAt"] = createdAt
	doc["updatedAt"] = updatedAt

	return doc
}
