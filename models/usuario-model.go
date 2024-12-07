package models

import (
	"bookedup/config"
	"database/sql"
	"fmt"
)

type Usuario struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func BuscarUsuarios() ([]Usuario, error) {
	rows, err := config.DB.Query("SELECT * FROM Usuario")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var novoUsuario Usuario

		if err := rows.Scan(&novoUsuario.ID, &novoUsuario.Nome, &novoUsuario.Email); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, novoUsuario)
	}

	return usuarios, nil
}

func BuscarUsuario(id string) (Usuario, error) {
	var usuario Usuario

	err := config.DB.QueryRow("SELECT * FROM Usuario WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nome, &usuario.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return Usuario{}, fmt.Errorf("user with ID %s not found", id)
		}
		return Usuario{}, err
	}

	return usuario, nil
}

func CriarUsuario(usuario Usuario) (int64, error) {
	result, err := config.DB.Exec("INSERT INTO Usuario (Nome, Email) VALUES (?, ?)", usuario.Nome, usuario.Email)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return id, nil
}

func EditarUsuario(novoUsuario Usuario, id string) error {
	result, err := config.DB.Exec("UPDATE Usuario SET Nome=?, Email=? WHERE id = ?", novoUsuario.Nome, novoUsuario.Email, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("unable to update user %s, id %s not found", novoUsuario.Nome, id)
	}

	return nil
}

func RemoverUsuario(id string) (Usuario, error) {
	var usuario Usuario

	err := config.DB.QueryRow("SELECT * FROM Usuario WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nome, &usuario.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return Usuario{}, fmt.Errorf("user with ID %s not found", id)
		}
		return Usuario{}, err
	}

	_, err = config.DB.Exec("DELETE FROM Usuario WHERE id = ?", id)

	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}
