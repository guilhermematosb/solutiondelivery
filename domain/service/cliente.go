package service

import (
	"database/sql"
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

	// Verificar CPF/CNPJ
	cliente.Preparar()

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
	w.Write([]byte("Implementar busca de um clinete apenas"))
}

// AtualizarCliente atualiza um cliente no banco de dados
func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Implementar atualização de um cliente"))
}

// DeletarCliente exclui um cliente no banco de dados
func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Implementar deletar cliente"))
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

	// Cria o canal
	channel := make(chan []model.Cliente)
	fmt.Printf("\nAguarde. Inserindo dados no banco de dados.\n")
	// Passa como parametro a primeira metade do slice e o channel
	go salvandoDados(finalData[:len(finalData)/2], channel, db)
	// Passa como parametro a segunda metado do slice e o channel
	go salvandoDados(finalData[len(finalData)/2:], channel, db)

	// Recebe os resultados dos channels
	x, y := <-channel, <-channel

	fmt.Printf(x[0].ToString())
	fmt.Printf(y[0].ToString())
	fmt.Printf("Dados Salvos!\n\n")
}

// Funcao que soma todos os valores de um slice e envia o resultado por um channel
func salvandoDados(finalData []model.Cliente, channel chan []model.Cliente, db *sql.DB) {

	repositorio := postgres.NovoRepositorioDeCliente(db)

	for i := range finalData {
		repositorio.Criar(finalData[i])
	}

	// Envia o resultado da soma para o channel
	channel <- finalData
}
