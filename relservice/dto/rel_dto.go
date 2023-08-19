package dto

//return dtos
type UserRelInfo struct {
	Id            uint   `json:"id"`
	OtherUsername string `json:"otherUsername"`
	Status        string `json:"status"`
	IsRequester   bool   `json:"isRequester"`
}

//Maybe this?
/*type UserRelInfo struct {
	Id            uint   `json:"id"`
	idRequester  uint      `json:"userRequester"`` (HASROLE ADMIN)
	idRequested  uint      `json:"userRequested"` (HASROLE ADMIN)
	usernameRequester  string `json:"usernameRequester"`
	usernameRequested  string `json:"usernameRequested"`
	Status        string `json:"status"`
	CreatedAt       time.Time `json:"created"` (HASROLE ADMIN)
	UpdatedAt       time.Time `json:"updated"` (HASROLE ADMIN)
}*/

//contracts
type AddUserRel struct {
	RequesterId uint `json:"requesterId"`
	RequestedId uint `json:"requestedId"`
}
