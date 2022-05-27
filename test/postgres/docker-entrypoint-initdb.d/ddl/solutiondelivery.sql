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

-- CREATE TABLE products
-- (
--     id    SERIAL,
--     name  TEXT           NOT NULL,
--     price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
--     CONSTRAINT products_pkey PRIMARY KEY (id)
-- );

-- insert into products (name, price)
-- values ('default_product', 1.00);
