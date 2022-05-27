package model

type Cliente struct {
	ID                 uint64 `json:"id,omitempty"`
	Cpf                string `json:"cpf,omitempty"`
	Private            string `json:"private,omitempty"`
	Incompleto         string `json:"incompleto,omitempty"`
	DataUltCompra      string `json:"data_ultima_compra"`
	TicketMedio        string `json:"ticket_medio"`
	TicketUltimaCompra string `json:"ticket_ultima_compra"`
	LojaMaisFreq       string `json:"loja_mais_frequente"`
	LojaUltCompra      string `json:"loja_ultima_compra"`
}
