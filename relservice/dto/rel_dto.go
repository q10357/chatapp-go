package dto

import "time"

//return dtos
/*
type UserRelInfo struct {
	Id            uint   `json:"id"`
	OtherUsername string `json:"otherUsername"`
	Status        string `json:"status"`
	IsRequester   bool   `json:"isRequester"`
}*/

type UserRelInfo struct {
	Id                uint      `json:"id"`
	IdRequester       uint      `json:"userRequester"` //(HASROLE ADMIN)
	IdRequested       uint      `json:"userRequested"` //(HASROLE ADMIN)
	UsernameRequester string    `json:"usernameRequester"`
	UsernameRequested string    `json:"usernameRequested"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created"` //(HASROLE ADMIN)
	UpdatedAt         time.Time `json:"updated"` //(HASROLE ADMIN)
}

//contracts
type AddUserRel struct {
	RequesterId uint `json:"requesterId"`
	RequestedId uint `json:"requestedId"`
}
