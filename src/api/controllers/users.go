package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
