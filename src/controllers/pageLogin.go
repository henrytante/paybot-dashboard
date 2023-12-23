package controllers

import (
	"net/http"

	"paybot-login/src/utils"
)



func PageLogin(w http.ResponseWriter, r *http.Request) {
	utils.RunHtml(w, "login.html", nil)
}
