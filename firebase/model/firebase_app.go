package model

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

type FirebaseApp struct {
	App *firebase.App
	DB  *db.Client
}
