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

func CreateUser(name, email, password string) (User, error) {
    var user User

    err := database.SQL.QueryRow(
        `INSERT INTO "user" (name, email, password)
           VALUES ($1, $2, $3)
         RETURNING id, name, email, active`,
        name, email, password,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.Active)

    return user, err
}

func ReadAllUsers() ([]User, error) {
    rows, err := database.SQL.Query(`SELECT id, name, email, active FROM "user"`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Active)
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
        `SELECT id, name, email, active FROM "user" WHERE id = $1`, id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.Active)

    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(id string, name, email, password string) (User, error) {
    var user User

    err := database.SQL.QueryRow(
        `UPDATE "user"
           SET name = $1, email = $2, password = $3
           WHERE id = $4
         RETURNING id, name, email, active`,
        name, email, password, id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.Active)

    return user, err
}

func DeleteUser(id string) error {
    var exists bool
    err := database.SQL.QueryRow(
        `SELECT EXISTS (SELECT 1 FROM "user" WHERE id = $1)`, id,
    ).Scan(&exists)
    if err != nil {
        return fmt.Errorf("erro ao verificar existência do usuário: %v", err)
    }
    if !exists {
        return fmt.Errorf("usuário com ID %d não existe", id)
    }

    _, err = database.SQL.Exec(`DELETE FROM "user" WHERE id = $1`, id)
    if err != nil {
        return fmt.Errorf("erro ao deletar usuário: %v", err)
    }

    return nil
}