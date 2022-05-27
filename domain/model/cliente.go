package model

import "github.com/klassmann/cpfcnpj"

type Cliente struct {
	ID                  uint64 `json:"id,omitempty"`
	Cpf                 string `json:"cpf,omitempty"`
	Private             string `json:"private,omitempty"`
	Incompleto          string `json:"incompleto,omitempty"`
	DataUltCompra       string `json:"data_ultima_compra"`
	TicketMedio         string `json:"ticket_medio"`
	TicketUltimaCompra  string `json:"ticket_ultima_compra"`
	LojaMaisFreq        string `json:"loja_mais_frequente"`
	LojaUltCompra       string `json:"loja_ultima_compra"`
	CpfValido           string `json:"cpf_valido"`
	CnpjValido          string `json:"cnpj_valido"`
	CnpjUltCompraValido string `json:"cnpj_ult_compra_valido"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (cliente *Cliente) Preparar() error {
	return cliente.validar()
}

func (cliente *Cliente) validar() error {
	if !cpfcnpj.ValidateCPF(cpfcnpj.Clean(cliente.Cpf)) {
		cliente.CpfValido = "1"
	} else {
		cliente.CpfValido = "0"
	}
	if cliente.LojaMaisFreq != "NULL" && !cpfcnpj.ValidateCNPJ(cpfcnpj.Clean(cliente.LojaMaisFreq)) {
		cliente.CnpjValido = "1"
	} else {
		cliente.CnpjValido = "0"
	}
	if cliente.LojaUltCompra != "NULL" && !cpfcnpj.ValidateCNPJ(cpfcnpj.Clean(cliente.LojaUltCompra)) {
		cliente.CnpjUltCompraValido = "1"
	} else {
		cliente.CnpjUltCompraValido = "0"
	}

	return nil
}

func (cliente *Cliente) ToString() string {
	return ""
}
