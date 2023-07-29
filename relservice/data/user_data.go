package data

type User struct {
	Id       uint64 `json:"userId"`
	Username string `json:"name"`
}

// Faking data
var userList = []User{
	{
		Id:       1,
		Username: "issichik",
	},
	{
		Id:       2,
		Username: "checkers",
	},
}
