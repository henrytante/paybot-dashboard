package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"paybot-login/src/models"
	"paybot-login/src/utils"
)

var data = PageData{
	ErrorMessage:   "",
	SuccessMessage: "",
}

func ConfigBackend(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		data.ErrorMessage = "Erro ao descodificar dados do post"
		utils.RunHtml(w, "config.html", data)
		return
	}
	opcao := r.FormValue("opcao")
	value := r.FormValue("input")
	if opcao == "" || value == "" {
		data.ErrorMessage = "Dados em branco"
		utils.RunHtml(w, "config.html", data)
		return
	} else {
		switch opcao {
		case "telegram":
			ChangeToken(w, r, value, opcao)
		case "banco":
			ChangeToken(w, r, value, opcao)
		case "chatID":
			ChangeToken(w, r, value, opcao)
		default:
			data.ErrorMessage = "Opção invalida"
			utils.RunHtml(w, "config.html", data)
			return
		}
	}
	
}

func ChangeToken(w http.ResponseWriter, r *http.Request, value string, opcao string) {
	jwt, err := r.Cookie("userJWT")
	if err != nil {
		data.ErrorMessage = err.Error()
		utils.RunHtml(w, "login.html", data)
		return
	}
	if jwt == nil || jwt.Value == "" {
		data.ErrorMessage = "Sem permissão!"
		utils.RunHtml(w, "login.html", data)
		return
	}
	var decodedJWT string
	if err := models.CookieHandler.Decode("userJWT", jwt.Value, &decodedJWT); err != nil {
		
		utils.RunHtml(w, "login.html", data)
		return
	}
	
	url := "http://localhost:8080/change-token"
	postData := map[string]interface{}{
		"token": value,
		"type":  opcao,
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		data.ErrorMessage = "Erro ao encodar em json"
		utils.RunHtml(w, "config.html", data)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		data.ErrorMessage = "Erro ao dar post na api"
		utils.RunHtml(w, "config.html", data)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decodedJWT))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "config.html", data)
		return
	}
	if resp.StatusCode != 200 {
		data.ErrorMessage = "Erro ao trocar o token, jwt invalido"
		utils.RunHtml(w, "config.html", data)
		return
	} else {
		data.SuccessMessage = "Token alterado com sucesso"
		utils.RunHtml(w, "config.html", data)
	}

}
