package service

import (
	"bufio"
	"os"
	"strings"

	"github.com/guilhermematosb/solutiondelivery/domain/model"
)

var docInvalido []string

// lê o arquivo de texto e coloca em um slice de strings
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/* separa os dados e converte para os tipos
 * compatíveis com o banco de dados */
func splitData(data []string) []model.Cliente {
	parsedData := make([]model.Cliente, len(data)-1)
	for i1, v1 := range data {
		if i1 == 0 {
			continue
		}
		lines := strings.Fields(v1)
		newIndex := i1 - 1
		parsedData[newIndex].Cpf = lines[0]
		parsedData[newIndex].Private = lines[1]
		parsedData[newIndex].Incompleto = lines[2]
		parsedData[newIndex].DataUltCompra = lines[3]
		parsedData[newIndex].TicketMedio = commaToPeriod(lines[4])
		parsedData[newIndex].TicketUltimaCompra = commaToPeriod(lines[5])
		parsedData[newIndex].LojaMaisFreq = lines[6]
		parsedData[newIndex].LojaUltCompra = lines[7]
	}

	return parsedData
}

// substitui vírgula por ponto
func commaToPeriod(value string) string {
	if value == "NULL" {
		value = "0"
	}
	value = strings.Replace(value, ",", ".", 1)

	return value
}

/* checa se o elemento ja está contido no slice
 * para evitar CPFs/CNPJs inválidos duplicados */
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
