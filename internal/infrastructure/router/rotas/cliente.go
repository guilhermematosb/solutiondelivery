package rotas

import (
	"net/http"

	"github.com/guilhermematosb/solutiondelivery/domain/service"
)

var rotasClientes = []Rota{
	{
		URI:                "/cliente",
		Metodo:             http.MethodPost,
		Funcao:             service.CriarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/clientes",
		Metodo:             http.MethodGet,
		Funcao:             service.BuscarClientes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/cliente/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             service.BuscarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/cliente/{clienteId}",
		Metodo:             http.MethodPut,
		Funcao:             service.AtualizarCliente,
		RequerAutenticacao: false,
	},
	{
		URI:                "/cliente/{clienteId}",
		Metodo:             http.MethodDelete,
		Funcao:             service.DeletarCliente,
		RequerAutenticacao: false,
	},
}
