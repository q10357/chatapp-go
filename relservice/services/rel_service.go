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

// This function will return the rels by userid.
func (r *RelService) GetRelsByUserId(userId uint) ([]*dto.UserRelInfo, error) {
	rels, err := r.relRepo.GetRelsByUserId(userId)
	var dtos = []*dto.UserRelInfo{}

	for _, rel := range rels {
		dto, err := r.ToUserRelResponse(rel, userId)
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

func (r *RelService) SetRelStatusToAccepted(id uint, userId uint) (*dto.UserRelInfo, error) {
	rel, err := r.relRepo.AcceptRel(id, userId)

	if err != nil {
		return nil, err
	}

	return r.ToUserRelResponse(rel, userId)
}

func (r *RelService) ToUserRelResponse(rel *rel.UserRel, userId uint) (*dto.UserRelInfo, error) {
	requester, err := r.userRepo.GetUserById(rel.UserIdRequester)
	requested, err := r.userRepo.GetUserById(rel.UserIdRequested)

	if err != nil {
		return nil, err
	}

	return &dto.UserRelInfo{
		Id:                rel.ID,
		IdRequester:       requester.Id,
		IdRequested:       requested.Id,
		UsernameRequester: requester.Username,
		UsernameRequested: requested.Username,
		Status:            rel.Status,
		CreatedAt:         rel.CreatedAt,
		UpdatedAt:         rel.UpdatedAt,
	}, nil
}
