package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//User is the model
type User struct {
	ID       string
	UserName string
	Email    string
	Password string
}

var (
	users []*User
)

//GetUsers - function for retrieving available users
func GetUsers() []*User {
	return users
}

//AddUser - function for adding new users
func AddUser(u User) (User, error) {
	if u.UserName == "" {
		return User{}, errors.New("New User must have a username")
	} else if u.Email == "" {
		return User{}, errors.New("New User must have a valid email")
	} else if len(u.Password) < 5 {
		return User{}, errors.New("Password must contain minimum of 6 characters")
	}
	u.ID = uuid.New().String()
	users = append(users, &u)
	return u, nil
}

//GetUserByUserName - function will retrieve specific user by it's ID
func GetUserByUserName(username string) (User, error) {
	for _, u := range users {
		if u.UserName == username {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with USERNAME '%v' not found", username)
}

//UpdateUser - function will update already existing user based on it's ID
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUserByID - function will remove user from collection if ID exists
func RemoveUserByID(id string) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
