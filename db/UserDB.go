package db

import (
	"log/slog"
)

func GetAllUsers() (map[int]User, error) {
	return nil, nil
}

func GetByID(id int) /* User */ error {
	return nil
}

func AddUser(u User) /* User, */ error {
	return nil
}

func EditUser(id int, u User) /* User, */ error {
	return nil
}

func DeleteUser(id int) error {

	return nil
}

func SearchUser(name string) (User, error) {

	slog.Error("User not found")
	return User{}, nil
}
