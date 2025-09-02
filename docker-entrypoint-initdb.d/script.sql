
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
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) >= 2),
    email VARCHAR(255) UNIQUE NOT NULL CHECK (position('@' IN email) > 1),
    active BOOLEAN NOT NULL DEFAULT FALSE,
    password TEXT NOT NULL
);

INSERT INTO "user" (name, email, password) VALUES
('Anaelise Cristiane Nunes Ribeiro Silveiro', 'anaelise.silveiro@app.com', 'Anaelise@2025'),
('João Pedro da Silva Júnior', 'joao.pedro@app.com', 'JoaoSilva@123'),
('Maria Clara de Souza Lima', 'maria.clara@app.com', 'ClaraLima@456'),
('Carlos Henrique Monteiro', 'carlos.henrique@app.com', 'CarlosH@789'),
('Fernanda Beatriz Ramos', 'fernanda.ramos@app.com', 'FernandaR@321');
/*('Lucas Eduardo Martins', 'lucas.martins@app.com', 'LucasM@654'),
('Patrícia Gomes Oliveira', 'patricia.oliveira@app.com', 'PatriciaO@987'),
('Rafael Augusto Nunes', 'rafael.nunes@app.com', 'RafaelN@159'),
('Juliana Tavares Costa', 'juliana.costa@app.com', 'JulianaC@753'),
('Bruno César Almeida', 'bruno.almeida@app.com', 'BrunoA@852'),
('Camila Rocha Fernandes', 'camila.fernandes@app.com', 'CamilaF@741'),
('Eduardo Vinícius Teixeira', 'eduardo.teixeira@app.com', 'EduardoT@369'),
('Larissa Helena Duarte', 'larissa.duarte@app.com', 'LarissaD@258'),
('Thiago Moura Cardoso', 'thiago.cardoso@app.com', 'ThiagoC@147'),
('Beatriz Antunes Figueiredo', 'beatriz.figueiredo@app.com', 'BeatrizF@963'),
('Gabriel Lopes Santana', 'gabriel.santana@app.com', 'GabrielS@852'),
('Isabela Martins Freitas', 'isabela.freitas@app.com', 'IsabelaF@741'),
('Felipe Cunha Moreira', 'felipe.moreira@app.com', 'FelipeM@159'),
('Natália Ribeiro Pires', 'natalia.pires@app.com', 'NataliaP@357'),
('Rodrigo Azevedo Lima', 'rodrigo.lima@app.com', 'RodrigoL@951'),
('Vanessa Cristina Moura', 'vanessa.moura@app.com', 'VanessaM@753'),
('André Luiz Barbosa', 'andre.barbosa@app.com', 'AndreB@456'),
('Tatiane Souza Mendes', 'tatiane.mendes@app.com', 'TatianeM@654'),
('Marcelo Henrique Duarte', 'marcelo.duarte@app.com', 'MarceloD@852'),
('Aline Ferreira Costa', 'aline.costa@app.com', 'AlineC@147'),
('Renato Oliveira Silva', 'renato.silva@app.com', 'RenatoS@369'),
('Débora Lima Andrade', 'debora.andrade@app.com', 'DeboraA@258'),
('Gustavo Pereira Rocha', 'gustavo.rocha@app.com', 'GustavoR@741'),
('Simone Batista Nogueira', 'simone.nogueira@app.com', 'SimoneN@963'),
('Daniela Castro Tavares', 'daniela.tavares@app.com', 'DanielaT@159'),
('Leonardo Almeida Pinto', 'leonardo.pinto@app.com', 'LeonardoP@357'),
('Bruna Santos Ribeiro', 'bruna.ribeiro@app.com', 'BrunaR@951'),
('Fábio Costa Martins', 'fabio.martins@app.com', 'FabioM@753'),
('Elaine Teixeira Lopes', 'elaine.lopes@app.com', 'ElaineL@456'),
('Ricardo Moura Fernandes', 'ricardo.fernandes@app.com', 'RicardoF@654'),
('Sabrina Rocha Cunha', 'sabrina.cunha@app.com', 'SabrinaC@852'),
('Henrique Lima Cardoso', 'henrique.cardoso@app.com', 'HenriqueC@147'),
('Luana Ribeiro Antunes', 'luana.antunes@app.com', 'LuanaA@369'),
('Pedro Vinícius Duarte', 'pedro.duarte@app.com', 'PedroD@258'),
('Jéssica Almeida Souza', 'jessica.souza@app.com', 'JessicaS@741'),
('Murilo Teixeira Barbosa', 'murilo.barbosa@app.com', 'MuriloB@963'),
('Tatiana Lopes Martins', 'tatiana.martins@app.com', 'TatianaM@159'),
('Diego Castro Oliveira', 'diego.oliveira@app.com', 'DiegoO@357'),
('Monique Freitas Lima', 'monique.lima@app.com', 'MoniqueL@951'),
('Wesley Henrique Pires', 'wesley.pires@app.com', 'WesleyP@753'),
('Renata Moura Fernandes', 'renata.fernandes@app.com', 'RenataF@456'),
('Vinícius Costa Andrade', 'vinicius.andrade@app.com', 'ViniciusA@654'),
('Carolina Batista Nunes', 'carolina.nunes@app.com', 'CarolinaN@852'),
('Alexandre Rocha Silva', 'alexandre.silva@app.com', 'AlexandreS@147');*/


-- Ver todas as tabelas do bando de dados.
-- SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';