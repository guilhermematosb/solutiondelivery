package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/guilhermematosb/solutiondelivery/domain/model"
	"github.com/guilhermematosb/solutiondelivery/internal/infrastructure/storage/postgres"
)

// CriarCliente adiciona um novo cliente no banco de dados
func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var cliente model.Cliente
	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		log.Fatal(erro)
	}

	db, erro := postgres.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := postgres.NovoRepositorioDeCliente(db)
	clienteId, erro := repositorio.Criar(cliente)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", clienteId)))
}

// BuscarClientes busca todos os clientes no banco de dados
func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("BuscarClientes"))
	db, erro := postgres.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := postgres.NovoRepositorioDeCliente(db)
	clientes, erro := repositorio.BuscarClientes()

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(clientes); erro != nil {
		w.Write([]byte("Erro ao converter os clientes para JSON"))
		return
	}
}

// BuscarCliente busca um cliente no banco de dados
func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BuscarCliente"))
}

// AtualizarCliente atualiza um cliente no banco de dados
func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AtualizarCliente"))
}

// DeletarCliente exclui um cliente no banco de dados
func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletarCliente"))
}

func ImportarClientes() {

	fileScanned, err := readLines("./data/base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalData := splitData(fileScanned)

	db, erro := postgres.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := postgres.NovoRepositorioDeCliente(db)
	fmt.Printf("\nAguarde. Inserindo dados no banco de dados.\n")
	for i := range finalData {
		repositorio.Criar(finalData[i])
	}
	fmt.Printf("Dados Salvos!\n\n")
}
