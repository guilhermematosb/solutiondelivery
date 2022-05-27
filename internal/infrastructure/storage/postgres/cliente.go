package postgres

import (
	"database/sql"
	"log"

	"github.com/guilhermematosb/solutiondelivery/domain/model"
)

// Cliente representa um repositório de usuarios
type Cliente struct {
	db *sql.DB
}

// NovoRepositorioDeCliente cria um repositório de clientes
func NovoRepositorioDeCliente(db *sql.DB) *Cliente {
	return &Cliente{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Cliente) Criar(cliente model.Cliente) (uint64, error) {
	erro := repositorio.db.QueryRow(
		"INSERT INTO clientes(cpf,private,incompleto,data_ultima_compra,ticket_medio,ticket_ultima_compra,loja_mais_frequente,loja_ultima_compra) VALUES($1, $2, $3, NULLIF($4,'NULL')::date, NULLIF($5,'')::numeric, NULLIF($6,'')::numeric, $7, $8) RETURNING id",
		cliente.Cpf, cliente.Private, cliente.Incompleto, cliente.DataUltCompra, cliente.TicketMedio, cliente.TicketUltimaCompra, cliente.LojaMaisFreq, cliente.LojaUltCompra).Scan(&cliente.ID)

	if erro != nil {
		log.Printf("ERROR: inserindo novo cliente: %q\n", erro)
		return 0, erro
	}
	return cliente.ID, nil
}

func (repositorio Cliente) BuscarClientes() ([]model.Cliente, error) {
	linhas, erro := repositorio.db.Query("select ID, cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra from clientes")
	if erro != nil {
		log.Printf("ERROR: Listando clientes: %q\n", erro)
		return nil, erro
	}
	defer linhas.Close()

	var clientes []model.Cliente
	for linhas.Next() {
		var cliente model.Cliente

		linhas.Scan(&cliente.ID, &cliente.Cpf, &cliente.Private, &cliente.Incompleto, &cliente.DataUltCompra, &cliente.TicketMedio, &cliente.TicketUltimaCompra, &cliente.LojaMaisFreq, &cliente.LojaUltCompra)

		clientes = append(clientes, cliente)
	}

	return clientes, nil

}
