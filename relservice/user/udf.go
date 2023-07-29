package user

var UserObjs = []User{}

var count = 1

// Queries
func GetUser(id int) User {
	for _, user := range UserObjs {
		if user.Id == id {
			return user
		}
	}
	/*if user == nil {
		//if no user found, return error not found
		errors.New("not found")
	} */
	return User{}
}
