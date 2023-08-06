package dto

//return dtos
type UserRelDto struct {
	Id            uint   `json:"id"`
	OtherUsername string `json:"otherUsername"`
	Status        string `json:"status"`
	IsRequester   bool   `json:"isRequester"`
}

type AdminUserRelDto struct {
	Id              uint   `json:"id"`
	UserIdRequester uint   `json:"userRequester"`
	UserIdRequested uint   `json:"userRequested"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created"`
	UpdatedAt       string `json:"updated"`
}

//contracts
type AddUserRelDto struct {
	RequesterId uint `json:"requesterId"`
	RequestedId uint `json:"requestedId"`
}
