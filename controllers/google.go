package controllers

import (
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleController struct {
	beego.Controller
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_SECRET_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_CLIENT"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	//TODO randomize it
	randomState = "random"
)

func (this *GoogleController) Get() {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, url, http.StatusTemporaryRedirect)
}
