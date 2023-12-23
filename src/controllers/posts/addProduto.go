package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"paybot-login/src/models"
	"paybot-login/src/utils"
	"strconv"
)

func AddProduto(w http.ResponseWriter, r *http.Request)  {
	jwt, err := r.Cookie("userJWT")
	if err != nil{
		data.ErrorMessage = "Erro ao coletar o JWT"
		utils.RunHtml(w, "login.html", data)
		return
	}
	if jwt == nil || jwt.Value == ""{
		data.ErrorMessage = "Cookie invalido"
		utils.RunHtml(w, "login.html", data)
		return
	}
	var decodedJWT string
	err = models.CookieHandler.Decode("userJWT", jwt.Value, &decodedJWT)
	if err != nil{
		data.ErrorMessage = "Cookie invalido"
		utils.RunHtml(w, "login.html", data)
		return
	}
	data := PageData{
		ErrorMessage: "",
		SuccessMessage: "",
	}

	url := "http://localhost:8080/add-produto"
	err = r.ParseForm()
	if err != nil{
		data.ErrorMessage = "Erro ao ler corpo do post"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	name := r.FormValue("nome")
	content := r.FormValue("conteudo")
	price := r.FormValue("valor")
	if name == "" || content == "" || price == ""{
		data.ErrorMessage = "Dados em branco"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	priceINT, err := strconv.Atoi(price)
	if err != nil{
		data.ErrorMessage = "Erro ao converter para int"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	postData := map[string]interface{}{
		"name": name,
		"content": content,
		"price": priceINT,
	}
	postDataJSON, err := json.Marshal(postData)
	if err != nil{
		data.ErrorMessage = "Erro ao converter dados em json"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postDataJSON))
	if err != nil{
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decodedJWT))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	if resp.StatusCode != 200{
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "add-produto.html", data)
		return
	}
	data.SuccessMessage = "Produto adicionado com sucesso"
	utils.RunHtml(w, "add-produto.html", data)
}