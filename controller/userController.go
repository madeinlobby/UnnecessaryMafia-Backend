package controller

import (
	"UnnecessaryMafia-Backend/model"
	"log"
)

func InsertUser(name, password, phone_number, email, fname, lname, status string) { //GameUser is the name of the table //todo hashing
	db := model.GetDbConnection()
	insForm, err := db.Prepare("INSERT INTO mafia.gameuser(username, password, `phone number`, email, fname, lname, status) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insForm.Exec(name, password, phone_number, email, fname, lname, status)
	log.Println("INSERT: Name: " + name)

	defer db.Close()
	return
}
