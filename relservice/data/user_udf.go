package data

var UserObjs = []User{}

var count = 1

// Queries
func GetUser(id uint64) User {
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
