package rel

import (
	"errors"
	"fmt"
	"time"
)

type RelRepo struct {
	UserRels []*UserRel
}

func NewRelRepo() *RelRepo {
	return &RelRepo{
		UserRels: []*UserRel{
			NewUserRel(1, 2, 1),
		},
	}
}

func (r *RelRepo) GetRelsByUserId(userId uint) ([]*UserRel, error) {
	var rels []*UserRel
	for _, rel := range r.UserRels {
		if rel.UserIdRequester == userId || rel.UserIdRequested == userId {
			rels = append(rels, rel)
		}
	}
	return rels, nil
}

func (r *RelRepo) CreateRel(partial *UserRel) *UserRel {
	newItem := partial
	newItem.ID = uint(len(r.UserRels)) + 1
	r.UserRels = append(r.UserRels, newItem)
	return newItem
}

func (r *RelRepo) GetAllRels() []*UserRel {
	return r.UserRels
}

func (r *RelRepo) GetRelById(id uint) (*UserRel, error) {
	for _, it := range r.UserRels {
		if it.ID == id {
			return it, nil
		}
	}
	return nil, fmt.Errorf("key '%d' not found", id)
}

func (r *RelRepo) UpdateRel(id uint, amended *UserRel) (*UserRel, error) {
	for i, it := range r.UserRels {
		if it.ID == id {
			amended.ID = id
			r.UserRels[i] = amended
			return amended, nil
		}
	}
	return nil, fmt.Errorf("key '%d' not found", id)
}

func (r *RelRepo) AcceptRel(id uint, userId uint) (*UserRel, error) {
	for _, it := range r.UserRels {
		if it.ID == id {
			if it.UserIdRequested != userId {
				return nil, errors.New("Unauthorized")
			}
			it.Status = "ACCEPTED"
			it.UpdatedAt = time.Now()
			return it, nil
		}
	}
	return nil, fmt.Errorf("key '%d' not found", id)
}

func (r *RelRepo) DeleteRelById(id uint) (bool, error) {
	for i, it := range r.UserRels {
		if it.ID == id {
			r.UserRels = append(r.UserRels[:i], r.UserRels[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}
