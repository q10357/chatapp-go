package rel

import "time"

type UserRel struct {
	ID              uint      `json:"id"`
	UserIdRequester uint      `json:"userRequester"`
	UserIdRequested uint      `json:"userRequested"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created"`
	UpdatedAt       time.Time `json:"updated"`
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
