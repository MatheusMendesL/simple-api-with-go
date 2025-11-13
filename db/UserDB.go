package db

import (
	"log/slog"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Biography string
}

var users = map[int]User{
	1: {Id: 1, FirstName: "Matheus", LastName: "Lopes", Biography: "Matheus is cool"},
	2: {Id: 2, FirstName: "Ana", LastName: "Silva", Biography: "Ana likes Go"},
}

func GetAllUsers() (map[int]User, error) {
	return users, nil
}

func GetByID(id int) (User, error) {

	user, exists := users[id]
	if !exists {
		slog.Error("This user no exists")
		return User{}, nil
	}

	return user, nil
}

func AddUser(u User) (User, error) {
	lastId := 0
	for _, v := range users {
		if v.Id > lastId {
			lastId = v.Id
		}
	}

	u.Id = lastId + 1
	users[u.Id] = u

	return u, nil
}
