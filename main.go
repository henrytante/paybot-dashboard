package main

import (
	"fmt"
	"net/http"
	"paybot-login/src/router"
	"paybot-login/src/utils"
)

func main() {
	fmt.Println("rodando login")
	r := router.Gerar()
	utils.LoadHtml()
	http.ListenAndServe(":3001", r)

}
