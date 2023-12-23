package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"paybot-login/src/models"
	"paybot-login/src/utils"
)

type PageDados struct {
	ErrorMessage   string
	SuccessMessage string
	UserID         string
}

var data = PageDados{
	ErrorMessage:   "",
	SuccessMessage: "",
}

var Produtos []Produto

type Produto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Price   int    `json:"price"`
	Success bool
}

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	jwt, err := r.Cookie("userJWT")
	if err != nil {
		data.ErrorMessage = "Erro ao coletar seu jwt"
		utils.RunHtml(w, "login.html", data)
		return
	}
	if jwt == nil || jwt.Value == "" {
		data.ErrorMessage = "Cookie invalido"
		utils.RunHtml(w, "login.html", data)
		return
	}

	var decodedJWT string
	if err = models.CookieHandler.Decode("userJWT", jwt.Value, &decodedJWT); err != nil {
		data.ErrorMessage = "Erro ao descriptografar o cookie"
		utils.RunHtml(w, "login.html", data)
		return
	}
	url := "http://localhost:8080/produtos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "login.html", data)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decodedJWT))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "login.html", data)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		data.ErrorMessage = "Erro ao ler corpo da resposta"
		utils.RunHtml(w, "login.html", data)
		return
	}
	
	if err = json.Unmarshal([]byte(string(body)), &Produtos); err != nil {
		data.ErrorMessage = "Erro ao ler corpo da resposta"
		utils.RunHtml(w, "login.html", data)
		return
	}
	if resp.StatusCode != 200 {
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "login.html", data)
		return
	} else {
		utils.RunHtml(w, "produtos.html", Produtos)
	}

}
