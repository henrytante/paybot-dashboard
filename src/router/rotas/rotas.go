package rotas

import (
	"net/http"
	"paybot-login/src/middleware"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI string
	Metodo string
	View string
	Funcao func(http.ResponseWriter, *http.Request)
	Auth bool
}

func Configurar(router *mux.Router) *mux.Router {
		
		rotas := dashRota
		for _, rota := range rotas{
			if rota.Auth{router.HandleFunc(rota.URI, middleware.VerifyToken(rota.View, rota.Funcao)).Methods(rota.Metodo)
			}else{
				router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
			}
		}
		fileServer := http.FileServer(http.Dir("./assets"))
		router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", fileServer))
		
		return router
}