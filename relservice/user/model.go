package user

type User struct {
	Id       int    `json:"userId"`
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
