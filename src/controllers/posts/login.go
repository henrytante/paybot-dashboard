package posts

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"paybot-login/src/db"
	"paybot-login/src/models"
	"paybot-login/src/utils"
	"time"
)

type PageData struct {
	ErrorMessage   string
	SuccessMessage string
	UserID string
}


func LoginBackend(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		ErrorMessage: "",
	}
	
	url := "http://localhost:8080/login"
	
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		data.ErrorMessage = "Dados em branco"
		utils.RunHtml(w, "login.html", data)
		return
	}
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var id int
	err = db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			data.ErrorMessage = "Usuario não encontrado"
			utils.RunHtml(w, "login.html", data)
			return
		} else {
			log.Fatal(err)
		}
	}
	postData := map[string]interface{}{
		"username": username,
		"password": password,
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		data.ErrorMessage = "Erro ao compilar o json"
		utils.RunHtml(w, "login.html", data)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		data.ErrorMessage = "Erro na conexão com a api"
		utils.RunHtml(w, "login.html", data)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ErrorMessage = "Erro ao enviar solicitação"
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
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		data.ErrorMessage = "Erro ao ler corpo da resposta"
		utils.RunHtml(w, "login.html", data)
		return
	}
	if mensagem, ok := responseData["token"]; ok {
		encodedID, err := models.CookieHandler.Encode("userID", id)
		if err == nil{
			cookie := &http.Cookie{
				Name: "userID",
				Value: encodedID,
				Expires: time.Now().Add(6 * time.Hour),
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
		encodedJWT, err := models.CookieHandler.Encode("userJWT", mensagem)
		if err == nil{
			cookie := &http.Cookie{
				Name: "userJWT",
				Value: encodedJWT,
				Expires: time.Now().Add(6 * time.Hour),
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
		
	} else {
		data.ErrorMessage = "Erro ao coletar seu jwt"
		utils.RunHtml(w, "login.html", data)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)
	
}
