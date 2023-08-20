package rel

import "time"

type UserRel struct {
	ID              uint
	UserIdRequester uint
	UserIdRequested uint
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewUserRel(id uint, userIdRequester uint, userIdRequested uint) *UserRel {
	return &UserRel{
		ID:              id,
		UserIdRequester: userIdRequester,
		UserIdRequested: userIdRequested,
		Status:          "PENDING",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
