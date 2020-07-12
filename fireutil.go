package fireutil

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func InitializeAppWithServiceAccount(serviceAccountKeyPath string, databaseURL string) *firebase.App {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	config := &firebase.Config{
		DatabaseURL: databaseURL,
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

func GetAccessToken(app *firebase.App, serviceAccountKeyPath string) string {
	data, _ := ioutil.ReadFile(serviceAccountKeyPath)

	scope := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database",
	}

	config, err := google.JWTConfigFromJSON(data, scope...)
	if err != nil {
		return ""
	}
	tokenSource := config.TokenSource(context.Background())

	token, err := tokenSource.Token()
	if err != nil {
		return ""
	}

	return token.AccessToken
}

func GetDatabaseSize(appName string, accessToken string) float32 {

	// https://trelva-app.firebaseio.com/.json?print=pretty&format=export&download=trelva-app-export.json&auth=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQ1ODM3Nzg2LCJpYXQiOjE1NDU4MzQxODYsInYiOjB9.3f_KrA3llLChq-eVQKpv7mWcgFe2SjH7FrftQZR1szM

	url := fmt.Sprintf("https://%s.firebaseio.com/.json?format=export&download=%s-export.json&access_token=%s", appName, appName, accessToken)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Println(err)
		return 0
	}

	log.Println(url)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return 0
	}

	length := len(body)
	size := float32(float32(length) / 1024)

	return size
}
