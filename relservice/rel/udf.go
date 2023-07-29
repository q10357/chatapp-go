package rel

import "fmt"

// Faking data
var relList = []UserRel{
	{
		Id:      1,
		UserId1: 1,
		UserId2: 2,
	},
}

var count = 1

// Queries
func GetRel(id int) UserRel {
	for _, rel := range relList {
		if rel.Id == id {
			return rel
		}
	}

	return UserRel{}
}

func GetRelsByUserId(userId int) []UserRel {
	fmt.Printf("UserId: %d\n", userId)
	rels := []UserRel{}
	for _, rel := range relList {
		if rel.UserId1 == userId || rel.UserId2 == userId {
			rels = append(rels, rel)
		}
	}
	return rels
}

func AddRel(userId1 int, userId2 int) *UserRel {
	tmp := &UserRel{
		Id:      count,
		UserId1: userId1,
		UserId2: userId2,
	}

	relList = append(relList, *tmp)

	count++

	return tmp
}
