package posts

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"paybot-login/src/models"
	"paybot-login/src/utils"
)

func MessageBack(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		ErrorMessage:   "",
		SuccessMessage: "",
	}
	
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	msg := r.FormValue("message")
	if msg == "" {
		data.ErrorMessage = "Dados em branco"
		utils.RunHtml(w, "message.html", data)
		return
	}
	// Obter o token JWT do cookie
	jwt, err := r.Cookie("userJWT")
	if err != nil {
		data.ErrorMessage = err.Error()
		utils.RunHtml(w, "login.html", data)
		return
	}
	if jwt.Value == "" || jwt == nil {
		data.ErrorMessage = "Sem permissão!"
		utils.RunHtml(w, "login.html", data)
		return
	}
	var decodedJWT string
	if err := models.CookieHandler.Decode("userJWT", jwt.Value, &decodedJWT); err != nil {
		data.ErrorMessage = "Cookie inválido"
		utils.RunHtml(w, "login.html", data)
		return
	}
	
	// Escapar os espaços no parâmetro msg
	msgEscaped := url.QueryEscape(msg)

	url := fmt.Sprintf("http://localhost:8080/message?msg=%s", msgEscaped)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a API"
		utils.RunHtml(w, "message.html", data)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decodedJWT))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a API"
		utils.RunHtml(w, "message.html", data)
		return
	}
	fmt.Println("StatusCode:", resp.StatusCode)
	defer resp.Body.Close()
	
	if resp.StatusCode == 200 {
		data.SuccessMessage = "Mensagem enviada!"
		utils.RunHtml(w, "message.html", data)
	} else {
		data.ErrorMessage = "Erro ao enviar mensagem, token telegram invalido"
		utils.RunHtml(w, "message.html", data)
	}
}
