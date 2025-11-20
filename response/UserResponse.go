package response

import (
	"_046_project/db"
	"_046_project/helper"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var _, queries, errDB = db.Conn()
var ctx = context.Background()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		slog.Error("Error with the method", "Error", "The method needs to be GET")
		return
	}

	if errDB != nil {
		slog.Error("Error to enter in the db", "Error", errDB)
		return
	}

	res, err := queries.ListUser(ctx)

	if err != nil {
		slog.Error("Error to do the query", "Error", err)
	}

	helper.Response(helper.Response_struct{Data: res}, w, http.StatusOK)
}

func GetByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("failed to parse id", "err", err, "raw_id", id)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	slog.Info("PARSED_ID", "idInt", idInt)
	res, err := queries.GetUser(ctx, idInt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.Response(helper.Response_struct{Data: "This user is not on the db"}, w, http.StatusNotFound)
		}
		slog.Error("Error to enter in the db", "Error", err)
		return
	}

	helper.Response(helper.Response_struct{Data: res}, w, http.StatusOK)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

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

	res, err := queries.CreateUser(ctx, db.CreateUserParams{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Biography: user.Biography})
	if err != nil {
		helper.Response(helper.Response_struct{Error: "failed to add user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: res}, w, http.StatusCreated)
}

func EditUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		slog.Error("Error with the method", "Error", "The method needs to be PUT")
		return
	}

	idStr := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(idStr)

	slog.Info("ids", idStr, idInt)

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

	if err := queries.UpdateUser(ctx, db.UpdateUserParams{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Biography: user.Biography,
		ID:        int64(idInt),
	}); err != nil {
		helper.Response(helper.Response_struct{Error: "failed to edit user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: "User edited"}, w, http.StatusCreated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		slog.Error("Error with the method", "Error", "The method needs to be DELETE")
		return
	}

	idStr := chi.URLParam(r, "id")
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	if err := queries.DeleteUser(ctx, idInt); err != nil {
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

	data, err := queries.SearchByName(ctx, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helper.Response(helper.Response_struct{Data: "This user is not on the db"}, w, http.StatusNotFound)
			return
		}
		helper.Response(helper.Response_struct{Error: "failed to search user"}, w, http.StatusInternalServerError)
		return
	}

	helper.Response(helper.Response_struct{Data: data}, w, http.StatusOK)
}
