package initGroup

import (
	myUUID "golang-api-demo/common/uuid"
	groupModel "golang-api-demo/modules/groups/model"
	"time"
)

func InitGroup(group *groupModel.Group) (map[string]interface{}, string) {
	id := myUUID.GenID()
	doc := make(map[string]interface{})

	// init default skill sender
	skillSender := make(map[string]interface{})
	skillSender["manager"] = 3
	skillSender["members"] = 1

	doc["id"] = id
	doc["name"] = group.Name
	doc["shortName"] = group.ShortName
	doc["icon"] = group.Icon
	doc["landId"] = group.LandId
	doc["owner"] = group.Owner
	doc["skillSender"] = skillSender
	doc["memberIds"] = group.MemberIds
	doc["createdAt"] = time.Now()
	doc["updatedAt"] = time.Now()
	doc["isDeleted"] = false

	return doc, id
}

func ParseGroup(group *groupModel.Group) (map[string]interface{}, string) {
	id := group.Id
	doc := make(map[string]interface{})

	doc["id"] = group.Id
	doc["name"] = group.Name
	doc["shortName"] = group.ShortName
	doc["icon"] = group.Icon
	doc["landId"] = group.LandId
	doc["owner"] = group.Owner
	doc["skillSender"] = group.SkillSender
	doc["memberIds"] = group.MemberIds
	doc["createdAt"] = group.CreatedAt
	doc["updatedAt"] = time.Now()
	doc["isDeleted"] = false

	return doc, id
}

func InitMemberGroup(member *groupModel.MemberGroup) (map[string]interface{}, string) {
	document := make(map[string]interface{})
	userId := member.UserId
	// init default skill sender
	skillSender := make(map[string]interface{})
	skillSender["manager"] = 3
	skillSender["members"] = 1

	document["id"] = userId
	document["userId"] = userId
	document["avatar"] = member.Avatar
	document["firstname"] = member.FirstName
	document["lastname"] = member.LastName
	document["role"] = member.Role
	document["tags"] = member.Tags
	document["skill"] = member.Skill
	document["sumori"] = member.Sumori
	document["komori"] = member.Komori
	document["createdAt"] = time.Now()
	document["updatedAt"] = time.Now()
	document["isDeleted"] = false

	return document, userId
}
