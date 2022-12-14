package controllers

import (
	"fmt"
	"net/http"

	config "github.com/md/go-auth/internal/configs"
	"github.com/md/go-auth/internal/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Authentication using OAuth2.0")
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create oauthState cookie
	oauthState := utils.GenerateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user
		from CSRF attacks. You must always provide a non-empty string
		and validate that it matches the the state query parameter
		on your redirect callback.
	*/

	u := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}
