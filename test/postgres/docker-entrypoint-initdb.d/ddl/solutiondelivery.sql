CREATE TABLE IF NOT EXISTS clientes (
    id    SERIAL,
    cpf varchar NOT NULL,
    private bool NOT NULL,
    incompleto bool NOT NULL,
    data_ultima_compra DATE,
    ticket_medio decimal(12, 2),
    ticket_ultima_compra decimal(12, 2),
    loja_mais_frequente varchar,
    loja_ultima_compra varchar
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
