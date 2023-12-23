package controllers

import (
	"net/http"
	"paybot-login/src/utils"
)


func PageConfig(w http.ResponseWriter, r *http.Request)  {
	utils.RunHtml(w, "config.html", nil)
}