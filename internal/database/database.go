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

//SQL, err = sql.Open("pgx", "postgres://docker:docker@localhost:5432/app")

package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib" // usando pgx driver
	"github.com/joho/godotenv"
)

// DB é uma struct que encapsula a conexão com o bancofmt.Println("\033[32mConnection established!\033[0m")
type DB struct {
	*sql.DB
}

var dbInstance *DB

//var ConnectedChan = make(chan bool)

func GetDB() *DB {
	ConnectionMonitor()

	return dbInstance
}

func GetSQLState(err error) error {
    if pgErr, ok := err.(*pgconn.PgError); ok {
        return fmt.Errorf(pgErr.Code)
    }
    return err

/*
	CREATE (INSERT):
	23505 - Unique violation
	23502 - Not null violation
	23503 - Foreign key violation
	22001 - Data too long

	READ (SELECT):
	P0002 - No data found (já capturado pelo PgError)
	42P01 - Table doesn't exist
	42703 - Column doesn't exist

	UPDATE:
	23505 - Data conflicts
	23502 - Required fields missing
	23503 - Invalid references

	DELETE:
	23503 - Foreign key violation (dependent records)
*/
}

func loadDBConnectionString(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf(".env file not found - Could not load environment variables")
	}

	raw := os.Getenv(key)
	if raw == "" {
		return "", fmt.Errorf("%s environment variable not set", key)
	}

	//fmt.Println(raw)

	// Expande variáveis internas ($PGUSER etc.) se houver
	dataSourceName := os.ExpandEnv(raw)

	//fmt.Println(dsn)

	return dataSourceName, nil
}

// NewConnection cria e retorna uma nova instância de DB com conexão estabelecida
func OpenConnection() error {
	dsn, err := loadDBConnectionString("DATABASE_URL")
	if err != nil {
		return fmt.Errorf("failed to load connection string: %w", err)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	// Configurar o pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("error connecting to database — check connection url or server availability: %w", err)
	}

	dbInstance = &DB{db}

	return nil
}

func ConnectionMonitor() {
	if dbInstance == nil || dbInstance.Ping() != nil {

		go func() {
			for {
				fmt.Println("\033[33mAttempting to connect to the database...\033[0m")

				err := OpenConnection()
				if err != nil {
					time.Sleep(10 * time.Second)
					fmt.Println("\033[31mFailed to connect:\033[0m", err)
					continue
				}
				break
			}

			defer fmt.Println("\033[32mConnection established!\033[0m")
			//ConnectedChan <- true // sinaliza sucesso
		}()
	}
}

/*func HandleRequest(w http.ResponseWriter, r *http.Request) {
    timeout := 3 * time.Second
    done := make(chan bool)

    go func() {
        for {
            if dbInstance == nil || dbInstance.Ping() != nil {
                err := OpenConnection()
                if err != nil {
                    time.Sleep(2 * time.Second)
            }
            done <- true
            return
        }
    }()

    select {
    case <-done:
        fmt.Fprintln(w, "Connection established.")
    case <-time.After(timeout):
        fmt.Fprintln(w, "Still trying to connect. Please wait or try again shortly.")
    }
}*/

func CloseConnection() error {
	if dbInstance == nil {
		return nil
	}

	err := dbInstance.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}

	return nil
}

// GetDBInfo retorna informações do banco (mantido como exemplo)
func (db *DB) GetDBInfo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		serverVersion     string
		maxConnections    int
		openedConnections int
	)

	err := db.QueryRowContext(
		ctx,
		`SELECT 
			current_setting('server_version') AS server_version, 
			current_setting('max_connections')::int AS max_connections, 
			(SELECT COUNT(*)::int FROM pg_stat_activity WHERE datname = $1)::int AS opened_connections;`,
		os.Getenv("PGDATABASE"),
	).Scan(&serverVersion, &maxConnections, &openedConnections)

	if err != nil {
		return fmt.Errorf("failed to get database info: %v", err)
	}

	fmt.Printf("Versão: %s \nMáx conexões: %d \nConexões abertas: %d\n",
		serverVersion, maxConnections, openedConnections)

	return nil
}
