# Aplicativo de Manipulação de Banco de Dados

Playlist do Projeto no Youtube - [Golang - Programação com GO](https://www.youtube.com/watch?v=Q6xQ3wD427Q&list=PL7AlK3EF-9TMZ2Upgk8nXC3I-QC_WWD7S)

Aplicação **Go** para manipulação de banco de dados com foco em **confiabilidade**, **monitoramento** e **utilização de recursos modernos da linguagem** como **goroutines** e **generics**.  
O projeto implementa CRUD de usuários, serialização JSON genérica, monitoramento de conexões e reconexão automática.

Com base nos conhecimentos adquiridos e na vontade de aprofundar meus estudos em Go, este projeto reúne práticas e aprendizados sobre a linguagem e a manipulação de dados. A partir dele, outros projetos foram desenvolvidos e podem ser encontrados aqui no meu GitHub. Com o tempo, vou aprimorando este repositório para que sirva como registro da minha jornada de estudo e também como referência para todos que se interessarem por essa linguagem fantástica.

---

## 🚀 Funcionalidades

- **CRUD de usuários** (Create, Read, Update, Delete)
- **UUID como chave primária** para maior consistência e unicidade
- **Serialização e leitura de JSON com Go generics**
- **Canal de comunicação entre goroutines** para operações concorrentes
- **ConnectionMonitor** para monitorar o estado da conexão
- **Reconexão automática** em caso de falha no banco de dados
- **Encapsulamento e monitoramento** em `database.go`
- **Configuração via arquivo `.env`**

---

## 🛠️ Tecnologias utilizadas

- [Go](https://golang.org/) — linguagem principal.
- Banco de dados relacional PostgreSQL.
- Leitura de variáveis de ambiente com o pacote godotenv.
- Goroutines e channels para concorrência.
- Uso de generics para funções e tipos genéricos.
- UUID para identificação única de usuários.
- Serialização e desserialização JSON.
- Docker Compose para subir o serviço de banco de dados e scripts de inicialização para criar a tabela "user" e adicionar dados de teste.

---

## ⚙️ Configuração do ambiente

1. Clone o repositório:

   ```
   git clone https://github.com/seu-usuario/Database-manipulation-application.git
   ```

   ```
   cd Database-manipulation-application
   ```

2. Criar o arquivo .env na raiz do projeto.

   - Banco de dados em nuvem com NEON PostgreSQL:

     - Veja o video disponível no YouTube:
       - [Conectando uma aplicação Go ao Neon](https://youtu.be/AyKDQxnrrX4)

   - Utilizar o Docker `docker-compose.yml`:
     - Renomeie o arquivo `env.example` para apenas `.env`.

3. Criar ou iniciar o banco de dados:

   ```
   docker-compose up
   ```

   OBS: Apagar todos os containers

   ```
   docker rm -f $(docker ps -aq)
   ```

4. Baixar os pacotes externos do projeto:

   ```
   go mod tidy
   ```

5. Compilar o projeto:

   ```
   go build
   ```

6. Executar o projeto:
   - Linux
   ```
   ./app
   ```
   - Windows
   ```
   .\app.exe
   ```
7. Endpoints de teste

   - documentation / user.http

   OBS: (Opcional) No VSCode use a extensão - REST Client (Huachao Mao)
