package controllers

import (
	"net/http"
	"paybot-login/src/utils"
)


func MensagemDashboard(w http.ResponseWriter, r *http.Request)  {
	utils.RunHtml(w, "message.html", nil)
}