package controllers

import (
	"net/http"
	"paybot-login/src/utils"
)


func PageAddProduto(w http.ResponseWriter, r *http.Request)  {
	utils.RunHtml(w, "add-produto.html", nil)
}