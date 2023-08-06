package services

import (
	"github.com/q10357/RelService/database/data/rel"
	"github.com/q10357/RelService/database/data/user"
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
	rels, err := r.relRepo.GetRelsByUserId(userId)
	var dtos = []*dto.UserRelDto{}

	for _, rel := range rels {
		dto, err := r.ToUserRelDto(rel, userId)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}

	return dtos, err
}

func (r *RelService) IsUserIsInRelation(relId uint, userId uint) (bool, error) {
	rel, err := r.relRepo.GetRelById(relId)

	if err != nil {
		return false, err
	}

	if userId != rel.UserIdRequested && userId != rel.UserIdRequester {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *RelService) SetRelStatusToAccepted(id uint, userId uint) (*dto.UserRelDto, error) {
	rel, err := r.relRepo.AcceptRel(id, userId)

	if err != nil {
		return nil, err
	}

	return r.ToUserRelDto(rel, userId)
}

func (r *RelService) ToUserRelDto(rel *rel.UserRel, userId uint) (*dto.UserRelDto, error) {
	var otherId uint
	if rel.UserIdRequester == userId {
		otherId = rel.UserIdRequested
	} else {
		otherId = rel.UserIdRequester
	}

	otherUser, err := r.userRepo.GetUserById(otherId)

	if err != nil {
		return nil, err
	}

	return &dto.UserRelDto{
		Id:            rel.ID,
		OtherUsername: otherUser.Username,
		Status:        rel.Status,
	}, nil
}
