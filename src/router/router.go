package router

import (
	"net/http"
	"paybot-login/src/controllers"
	"paybot-login/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)
	return rotas.Configurar(r)
}