package controllers

import (
	"net/http"
	"paybot-login/src/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request)  {
	utils.RunHtml(w, "404.html", nil)
	
}