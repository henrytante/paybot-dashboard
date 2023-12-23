package rotas

import (
	"net/http"
	"paybot-login/src/controllers"
	"paybot-login/src/controllers/posts"
)

var dashRota = []Rota {
	{
		URI: "/",
		Metodo: http.MethodGet,
		View: "login.html",
		Funcao: controllers.PageLogin,
		Auth: false,
	},
	{
		URI: "/dashboard",
		Metodo: http.MethodGet,
		View: "dashboard.html",
		Funcao: controllers.PageDashboard,
		Auth: true,
	},
	{
		URI: "/",
		Metodo: http.MethodPost,
		View: "login.html",
		Funcao: posts.LoginBackend,
		Auth: false,
	},
	{
		URI: "/mensagem",
		Metodo: http.MethodGet,
		View: "message.html",
		Funcao: controllers.MensagemDashboard,
		Auth: true,
	},
	{
		URI: "/mensagem",
		Metodo: http.MethodPost,
		View: "message.html",
		Funcao: posts.MessageBack,
		Auth: false,
	},
	{
		URI: "/config",
		Metodo: http.MethodGet,
		View: "config.html",
		Funcao: controllers.PageConfig,
		Auth: true,
	},
	{
		URI: "/config",
		Metodo: http.MethodPost,
		Funcao: posts.ConfigBackend,
	},
	{
		URI: "/add-produtos",
		Metodo: http.MethodGet,
		View: "add-produto.html",
		Funcao: controllers.PageAddProduto,
		Auth: true,
	},
	{
		URI: "/add-produtos",
		Metodo: http.MethodPost,
		Funcao: posts.AddProduto,
	},
	{
		URI: "/produtos",
		Metodo: http.MethodGet,
		View: "produtos.html",
		Funcao: controllers.GetProdutos,
		Auth: false,
	},
	{
		URI: "/delete-product",
		Metodo: http.MethodPost,
		Funcao: posts.DeleteProduto,
	},
}