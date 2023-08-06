package user

type IUserRepo[T any] interface {
	CreateUser(*T) *T
	GetAllUsers() []*T
	GetUserById(uint) (*T, error)
	//GetUserByUsername(string) (*T, error)
	UpdateUser(uint, *T) (*T, error)
	DeleteUserById(uint) (bool, error)
}
