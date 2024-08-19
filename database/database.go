/*

   Em um sistema em produção, geralmente é preferível manter uma conexão persistente
   com o banco de dados, em vez de abrir e fechar conexões para cada consulta. Abrir
   uma conexão com um banco de dados é uma operação relativamente cara em termos de
   tempo e recursos. Portanto, se você estiver fazendo muitas consultas, o custo de
   abrir e fechar a conexão repetidamente pode se somar.

   No entanto, manter uma conexão aberta indefinidamente também tem suas desvantagens.
   Por exemplo, se o seu aplicativo mantém muitas conexões abertas simultaneamente,
   isso pode sobrecarregar o banco de dados. Além disso, se a conexão for interrompida
   por algum motivo (por exemplo, se o banco de dados cair ou a rede falhar), seu
   aplicativo precisará ser capaz de lidar com isso.

   Uma abordagem comum é usar um pool de conexões. Um pool de conexões mantém um número
   de conexões abertas e reutiliza-as conforme necessário. Quando uma consulta é feita,
   uma conexão é retirada do pool, usada e depois retornada ao pool. Isso oferece um bom
   equilíbrio entre eficiência (porque você não precisa abrir uma nova conexão para cada
   consulta) e uso de recursos (porque você limita o número de conexões abertas).

   Neste código, SetMaxOpenConns(25) define o número máximo de conexões abertas para
   o banco de dados, SetMaxIdleConns(25) define o número máximo de conexões ociosas
   que podem existir simultaneamente e SetConnMaxLifetime(5 * time.Minute) define a
   duração máxima que uma conexão pode ser reutilizada.
*/

package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var SQL *sql.DB

func connectDatabase(driver, dataSourceName string) error {
	var err error
	SQL, err = sql.Open(driver, dataSourceName)
	if err != nil {
		return err
	}

	// Configurar o pool de conexões
	SQL.SetMaxOpenConns(25)
	SQL.SetMaxIdleConns(25)
	SQL.SetConnMaxLifetime(5 * time.Minute)

	// Verificar a conexão
	err = SQL.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Disconnect() {
	err := SQL.Close()
	if err != nil {
		log.Fatalf("Failed to close connection: %v", err)
	}
}

func createTable() {
	_, err := SQL.Exec("CREATE TABLE IF NOT EXISTS \"user\" (id SERIAL PRIMARY KEY, email VARCHAR(100) UNIQUE NOT NULL, password VARCHAR(100) NOT NULL)")
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func Connection() {

	err := connectDatabase("pgx", "postgres://docker:docker@localhost:5432/app")
	//err := connectDatabase("pgx", "postgres://postgres:postgres@localhost:5432/app")
	//err := connectDatabase("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	createTable()

	fmt.Println("Successfully connected!")
}
