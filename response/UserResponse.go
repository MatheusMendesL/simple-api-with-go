package response

import (
	"_046_project/db"
	"_046_project/helper"
	"context"
	"log/slog"
	"net/http"
)

var _, queries, errDB = db.Conn()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		slog.Error("Error with the method", "Error", "The method needs to be GET")
		return
	}

	if errDB != nil {
		slog.Error("Error to enter in the db", "Error", errDB)
		return
	}

	ctx := context.Background()

	res, err := queries.ListUser(ctx)

	if err != nil {
		slog.Error("Error to do the query", "Error", err)
	}

	slog.Info("chegou aq")

	helper.Response(helper.Response_struct{Data: res}, w, http.StatusOK)
}

/* // tlvz seja inutil
func GetByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if r.Method != "GET" {
		slog.Error("Error with the method", "Error", "The method needs to be GET")
		return
	}

	idInt, _ := strconv.Atoi(id)
	idInt = idInt
	 res, err := (idInt)

	if err != nil {
		slog.Error("Error to enter in the db", "Error", err)
		return
	}

	if (res == db.User{}) {
		helper.Response(helper.Response_struct{Error: "This user not exists on DB"}, w, http.StatusNotFound)
		return
	}

	helper.Response(helper.Response_struct{Data: /* res  "aa"}, w, http.StatusOK)
} */

/* func AddUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		slog.Error("Error with the method", "Error", "The method needs to be POST")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1000)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			helper.Response(helper.Response_struct{Error: "Body too large"}, w, http.StatusRequestEntityTooLarge)
			return
		}
		slog.Error("failed to read request body", "error", err)
		helper.Response(helper.Response_struct{Error: "something went wrong"}, w, http.StatusInternalServerError)
		return
	}

	var user db.User
	if err := json.Unmarshal(data, &user); err != nil {
		helper.Response(helper.Response_struct{Error: "invalid JSON body"}, w, http.StatusUnprocessableEntity)
		return
	}

	newUser, err := db.AddUser(user)
	if err != nil {
		helper.Response(helper.Response_struct{Error: "failed to add user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: newUser}, w, http.StatusCreated)
}

func EditUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		slog.Error("Error with the method", "Error", "The method needs to be PUT")
		return
	}

	idStr := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(idStr)

	r.Body = http.MaxBytesReader(w, r.Body, 1000)
	data, err := io.ReadAll(r.Body)

	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			helper.Response(helper.Response_struct{Error: "Body too large"}, w, http.StatusRequestEntityTooLarge)
			return
		}
		slog.Error("failed to read request body", "error", err)
		helper.Response(helper.Response_struct{Error: "something went wrong"}, w, http.StatusInternalServerError)
		return
	}

	var user db.User
	if err := json.Unmarshal(data, &user); err != nil {
		helper.Response(helper.Response_struct{Error: "invalid JSON body"}, w, http.StatusUnprocessableEntity)
		return
	}

	EditedUser, err := db.EditUser(idInt, user)
	if err != nil {
		helper.Response(helper.Response_struct{Error: "failed to add user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: EditedUser}, w, http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		slog.Error("Error with the method", "Error", "The method needs to be DELETE")
		return
	}

	idStr := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(idStr)

	err := db.DeleteUser(idInt)
	if err != nil {
		helper.Response(helper.Response_struct{Error: "failed to delete user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: "User deleted"}, w, http.StatusCreated)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		slog.Error("Error with the method", "Error", "The method needs to be GET")
		return
	}

	name := chi.URLParam(r, "name")

	data, err := db.SearchUser(name)

	if err != nil {
		helper.Response(helper.Response_struct{Error: "failed to search user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: data}, w, http.StatusOK)
}
*/
