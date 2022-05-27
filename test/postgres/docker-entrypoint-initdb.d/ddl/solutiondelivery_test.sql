-- TEST DB

CREATE TABLE IF NOT EXISTS clientes (
    id    SERIAL,
    cpf varchar NOT NULL,
    private bool NOT NULL,
    incompleto bool NOT NULL,
    data_ultima_compra DATE,
    ticket_medio varchar,
    ticket_ultima_compra varchar,
    loja_mais_frequente varchar,
    loja_ultima_compra varchar,
    cpf_valido bool,
    cnpj_valido bool,
    cnpj_ult_compra_valido bool
);