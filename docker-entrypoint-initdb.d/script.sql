
--DROP DATABASE app;

--CREATE DATABASE app;

-- Crie o banco de dados 'app' com o proprietário 'postgres'
/*CREATE DATABASE app;
WITH OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    LC_COLLATE = 'pt_BR.UTF-8'
    LC_CTYPE = 'pt_BR.UTF-8'
    CONNECTION LIMIT = 25;*/

-- Crie o esquema 'app' dentro do banco de dados
CREATE SCHEMA IF NOT EXISTS app;

-- Altera o número de conexões limite para um único usuário.
-- -1 deixa ilimitado.
--ALTER DATABASE app WITH CONNECTION LIMIT = -1;
ALTER DATABASE app WITH CONNECTION LIMIT = 25;

\c app

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY, 
    email VARCHAR(100) UNIQUE NOT NULL, 
    password VARCHAR(100) NOT NULL
);

INSERT INTO "user" (email, password) VALUES ('test@test.com', '123');
INSERT INTO "user" (email, password) VALUES ('user1@test.com', 'password1');
INSERT INTO "user" (email, password) VALUES ('user2@test.com', 'password2');
INSERT INTO "user" (email, password) VALUES ('user3@test.com', 'password3');
INSERT INTO "user" (email, password) VALUES ('user4@test.com', 'password4');
INSERT INTO "user" (email, password) VALUES ('user5@test.com', 'password5');
INSERT INTO "user" (email, password) VALUES ('user6@test.com', 'password6');
INSERT INTO "user" (email, password) VALUES ('user7@test.com', 'password7');
INSERT INTO "user" (email, password) VALUES ('user8@test.com', 'password8');
INSERT INTO "user" (email, password) VALUES ('user9@test.com', 'password9');

-- Ver todas as tabelas do bando de dados.
-- SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';