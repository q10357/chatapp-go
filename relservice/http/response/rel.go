package response

//return dtos
type GetUserRel struct {
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
