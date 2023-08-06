package services

import (
	"fmt"

	"github.com/q10357/RelService/data/rel"
	"github.com/q10357/RelService/data/user"
	"github.com/q10357/RelService/dto"
)

type RelService struct {
	relRepo  rel.IRelRepo[rel.UserRel]
	userRepo user.IUserRepo[user.User]
}

func NewRelService(rr rel.IRelRepo[rel.UserRel], ur user.IUserRepo[user.User]) *RelService {
	return &RelService{relRepo: rr, userRepo: ur}
}

func (r *RelService) GetRelsByUserId(userId uint) ([]*dto.UserRelDto, error) {
	fmt.Printf("UserId: %d\n", userId)
	rels, err := r.relRepo.GetRelsByUserId(userId)
	var tmp = []*dto.UserRelDto{}

	for _, rel := range rels {
		var otherId uint
		if rel.UserIdRequester == userId {
			otherId = rel.UserIdRequester
		} else {
			otherId = rel.UserIdRequested
		}

		tmp = append(tmp, r.ToUserRelDto(rel, otherId))
	}

	return tmp, err
}

func (r *RelService) ToUserRelDto(userRel *rel.UserRel, otherUserId uint) *dto.UserRelDto {
	otherUser, err := r.userRepo.GetUserById(otherUserId)
	if err != nil {
		// handle error
	}

	return &dto.UserRelDto{
		Id:            userRel.ID,
		OtherUsername: otherUser.Username,
		Status:        userRel.Status,
	}
}
