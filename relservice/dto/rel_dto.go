package dto

//return dtos
type UserRelInfo struct {
	Id            uint   `json:"id"`
	OtherUsername string `json:"otherUsername"`
	Status        string `json:"status"`
	IsRequester   bool   `json:"isRequester"`
}

//contracts
type AddUserRel struct {
	RequesterId uint `json:"requesterId"`
	RequestedId uint `json:"requestedId"`
}
