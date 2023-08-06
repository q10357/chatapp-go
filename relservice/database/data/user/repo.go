package user

import "fmt"

type UserRepo struct {
	userObjs []*User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		userObjs: []*User{
			{Id: 1, Username: "issichik"},
			{Id: 2, Username: "checkers"},
		},
	}
}

func (u *UserRepo) CreateUser(partial *User) *User {
	newItem := partial
	newItem.Id = uint(len(u.userObjs) + 1)
	u.userObjs = append(u.userObjs, newItem)
	return newItem
}

func (u *UserRepo) GetAllUsers() []*User {
	return u.userObjs
}

func (u *UserRepo) GetUserById(id uint) (*User, error) {
	for _, it := range u.userObjs {
		if it.Id == id {
			return it, nil
		}
	}
	return nil, fmt.Errorf("key '%d' not found", id)
}

func (u *UserRepo) UpdateUser(id uint, amended *User) (*User, error) {
	for i, it := range u.userObjs {
		if it.Id == id {
			amended.Id = id
			u.userObjs = append(u.userObjs[:i], u.userObjs[i+1:]...)
			u.userObjs = append(u.userObjs, amended)
			return amended, nil
		}
	}
	return nil, fmt.Errorf("key '%d' not found", amended.Id)
}

func (u *UserRepo) DeleteUserById(id uint) (bool, error) {
	for i, it := range u.userObjs {
		if it.Id == id {
			u.userObjs = append(u.userObjs[:i], u.userObjs[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}
