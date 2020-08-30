package server

import (
	"UnnecessaryMafia-Backend/controller"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	controller.InsertUser(
		r.FormValue("username"),
		r.FormValue("password"),
		r.FormValue("phone_number"),
		r.FormValue("email"),
		r.FormValue("fname"),
		r.FormValue("lname"),
		r.FormValue("status"))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	gameUser := controller.GetUser(
		r.FormValue("username"),
		r.FormValue("password"))
	jsonResp, err := json.MarshalIndent(gameUser, "", "	")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("could not write response: %s", r.RequestURI)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
