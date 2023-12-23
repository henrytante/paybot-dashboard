package middleware

import (
	"encoding/json"
	"fmt"
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
	Id             int
	Time           string
	Now            string
	Username       string
	Vendas         int `json:"vendas"`
	TotalV         int `json:"valor_total_vendas"`
	TelegramToken  string
	BancoToken     string
}

func VerifyToken(page string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data PageData
		jwtCookie, err := r.Cookie("userJWT")
		if err != nil || jwtCookie == nil || jwtCookie.Value == "" {
			data.ErrorMessage = "Cookie em branco, faça o login"
			utils.RunHtml(w, "login.html", data)
			return
		}

		var jwt string
		if err := models.CookieHandler.Decode("userJWT", jwtCookie.Value, &jwt); err != nil {
			data.ErrorMessage = "Sem permissão!"
			utils.RunHtml(w, "login.html", data)
			return
		}

		url := "http://localhost:8080/verify"
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwt))
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			data.ErrorMessage = "Erro ao verificar o token"
			utils.RunHtml(w, "login.html", data)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			var id int
			userID, err := r.Cookie("userID")
			if err != nil || userID == nil || userID.Value == "" {
				data.ErrorMessage = "Cookie em branco, faça o login"
				utils.RunHtml(w, "login.html", data)
				return
			}
			if err := models.CookieHandler.Decode("userID", userID.Value, &id); err != nil {
				data.ErrorMessage = "Sem permissão!"
				utils.RunHtml(w, "login.html", data)
				return
			}
			data.Id = id
			currentHour := time.Now().Hour()
			if currentHour >= 5 && currentHour < 12 {
				data.Time = "Bom dia"
			} else if currentHour >= 12 && currentHour < 19 {
				data.Time = "Boa tarde"
			} else {
				data.Time = "Boa noite"
			}
			
			var user string
			dbConn, err := db.ConnectDB() // Renomeando db para dbConn para evitar confusão
			if err != nil {
				log.Fatal(err)
			}
			err = dbConn.QueryRow("SELECT username FROM users WHERE id = ?", data.Id).Scan(&user)
			if err != nil {
				data.ErrorMessage = "Erro ao coletar seu username"
				utils.RunHtml(w, "login.html", data)
				return
			}
			data.Username = user

			// Renomeando client e req para evitar conflito com variáveis já existentes
			userURL := fmt.Sprintf("http://localhost:8080/user?id=%d", data.Id)
			userClient := &http.Client{}
			userReq, err := http.NewRequest("GET", userURL, nil)
			if err != nil {
				data.ErrorMessage = err.Error()
				utils.RunHtml(w, "login.html", data)
				return
			}
			userReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
			userResp, err := userClient.Do(userReq)
			if err != nil {
				data.ErrorMessage = err.Error()
				utils.RunHtml(w, "login.html", data)
				return
			}
			defer userResp.Body.Close()
			body, err := ioutil.ReadAll(userResp.Body)
			if err != nil {
				data.ErrorMessage = err.Error()
				utils.RunHtml(w, "login.html", data)
				return
			}
			if err := json.Unmarshal(body, &data); err != nil {
				data.ErrorMessage = err.Error()
				utils.RunHtml(w, "login.html", data)
				return
			}
			
			utils.RunHtml(w, page, data)
		} else {
			data.ErrorMessage = "Faça o login antes de acessar esta página"
			utils.RunHtml(w, "login.html", data)
			return
		}
	}
}
