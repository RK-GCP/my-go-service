package models

import (
	"errors"
	"fmt"
)

//User type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers - all users
func GetUsers() []*User {
	return users
}

//AddUser to the collection
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("Newly created user id must be 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID - Returns User By Id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User Id '%v' not found", id)
}

//UpdateUser - updates a User
func UpdateUser(u User) (User, error) {
	for indx, candidate := range users {
		if candidate.ID == u.ID {
			users[indx] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with Id '%v' not found", u.ID)
}

//RemoveUserByID - removes a user by Id
func RemoveUserByID(id int) error {
	for indx, candidate := range users {
		if candidate.ID == id {
			users = append(users[:indx], users[indx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with Id '%v' not found", id)
}
