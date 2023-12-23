package controllers

import (
	"fmt"
	"net/http"
	"paybot-login/src/utils"
	
)

type PageData struct {
	ID string
}

func PageDashboard(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("userID")
	if err != nil {
		if id == nil {
			erro := "Logue-se"
			utils.RunHtml(w, "login.html", erro)
			return
		}else if id.Value == ""{
			erro := "Cookie invalido"
			utils.RunHtml(w, "login.html", erro)
			return
		}
	}

	data := PageData{
		ID: id.Value,
	}
	fmt.Println("ID: ", data.ID)
	utils.RunHtml(w, "dashboard.html", data)

}
