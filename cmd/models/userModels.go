package models

import (
	"database/sql"
	"fmt"

	"github.com/RodrigoMS/app/database"
)

/*func GetUser() app.User {
	user := app.User{
		Email:    "localhost@localhost.com",
		Password: "123456789",
	}

	return user
}*/

func CreateUser(email, password string) (User, error) {
	var user User

	err := database.SQL.QueryRow(
		"INSERT INTO \"user\" (email, password) VALUES ($1, $2) RETURNING id, email", email, password,
	).Scan(&user.ID, &user.Email)

	// Inserir e retornas o ID do novo registro - Somente no PostgreSQL:
	// INSERT INTO User (email, password) VALUES ($1, $2) RETURNING id

	// Para o MySQL esta mesma função será
	// INSERT INTO User (email, password) VALUES (?, ?)
	// SELECT LAST_INSERT_ID();

	return user, err
}

func ReadAllUsers() ([]User, error) {
	rows, err := database.SQL.Query("SELECT id, email, password FROM \"user\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return users, nil
}

func ReadUser(id string) (*User, error) {
	var user User

	err := database.SQL.QueryRow(
		"SELECT id, email, password FROM \"user\" WHERE id = $1", id,
	).Scan(
		&user.ID, &user.Email, &user.Password,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(email, password, id string) (User, error) {
	var user User

	err := database.SQL.QueryRow(
		"UPDATE \"user\" SET email = $1, password = $2 WHERE id = $3 RETURNING id, email", email, password, id,
	).Scan(
		&user.ID, &user.Email,
	)

	return user, err
}

func DeleteUser(id string) error {
	// Verificar se o usuário existe
	var exists bool
	err := database.SQL.QueryRow("SELECT exists (SELECT 1 FROM \"user\" WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		//return fmt.Errorf("erro ao verificar a existência do usuário: %v", err)
		return err
	}
	if !exists {
		//return fmt.Errorf("Usuário com ID %s não existe", id)
		return fmt.Errorf(id)
	}

	// Deletar o usuário
	_, err = database.SQL.Exec("DELETE FROM \"user\" WHERE id = $1", id)
	if err != nil {
			//return fmt.Errorf("erro ao deletar o usuário: %v", err)
			return err
	}

	return nil
}
