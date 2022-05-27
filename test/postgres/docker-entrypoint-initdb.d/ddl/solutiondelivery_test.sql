-- TEST DB

CREATE TABLE IF NOT EXISTS clientes (
    cpf varchar NOT NULL,
    private bool NOT NULL,
    incompleto bool NOT NULL,
    ultima_compra DATE,
    ticket_medio decimal(12, 2),
    ticket_ultimo decimal(12, 2),
    loja_mais_frequente varchar,
    loja_ultima_compra varchar
);