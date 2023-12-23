package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"paybot-login/src/controllers"
	"paybot-login/src/models"
	"paybot-login/src/utils"
	"strconv"
)

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	var data PageData

	jwt, err := r.Cookie("userJWT")
	if err != nil || jwt == nil || jwt.Value == "" {
		data.ErrorMessage = err.Error()
		utils.RunHtml(w, "login.html", data)
		return
	}

	var decodedJWT string
	if err := models.CookieHandler.Decode("userJWT", jwt.Value, &decodedJWT); err != nil {
		data.ErrorMessage = err.Error()
		utils.RunHtml(w, "login.html", data)
		return
	}

	err = r.ParseForm()
	if err != nil {
		data.ErrorMessage = err.Error()
		utils.RunHtml(w, "dashboard.html", data)
		return
	}

	pid := r.FormValue("product_id")
	if pid == "" {
		data.ErrorMessage = "Dados em branco"
		utils.RunHtml(w, "produtos.html", data)
		return
	}
	pidINT, err := strconv.Atoi(pid)
	if err != nil{
		data.ErrorMessage = "Erro ao tentar converter o id para int"
		utils.RunHtml(w, "produtos.html", data)
		return
	}
	postData := map[string]int{
		"pid": pidINT,
	}
	postDataJSON, err := json.Marshal(postData)
	if err != nil{
		data.ErrorMessage = "Erro ao tentar converter os dados para json"
		utils.RunHtml(w, "produtos.html", data)
		return
	}
	url := "http://localhost:8080/deletar-produto"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postDataJSON))
	if err != nil {
		data.ErrorMessage = "Erro na criação da requisição"
		utils.RunHtml(w, "produtos.html", data)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decodedJWT))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a API"
		utils.RunHtml(w, "produtos.html", data)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data.ErrorMessage = "Erro ao deletar produto"
		utils.RunHtml(w, "produtos.html", data)
		return
	} else {
		controllers.GetProdutos(w, r)
		
	}
}
