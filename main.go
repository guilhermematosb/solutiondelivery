package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guilhermematosb/solutiondelivery/domain/service"
	"github.com/guilhermematosb/solutiondelivery/internal/infrastructure/config"
	"github.com/guilhermematosb/solutiondelivery/internal/infrastructure/router"
	_ "github.com/lib/pq"
)

func main() {
	config.Carregar()

	// Iniciar importação
	service.ImportarClientes()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
