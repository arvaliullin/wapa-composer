package persistence

import (
	"database/sql"
	"encoding/json"

	"github.com/arvaliullin/wapa-composer/internal/domain"
)

type DesignRepositoryContract interface {
	Create(design domain.Design) error
	Delete(id string) error
	GetAll() ([]domain.Design, error)
}

type DesignRepository struct {
	DbConnection string
}

type Action func(conn *sql.DB) error

func (repo *DesignRepository) withConnection(action Action) error {
	db, err := sql.Open("postgres", repo.DbConnection)
	if err != nil {
		return err
	}
	defer db.Close()

	return action(db)
}

func (repo *DesignRepository) Create(design domain.Design) error {
	return repo.withConnection(func(conn *sql.DB) error {
		query := `
			INSERT INTO composer.design (name, description, js_filename, wasm_filename)
			VALUES ($1, $2, $3, $4)`
		_, err := conn.Exec(query, design.Name, design.Description, design.JSFilename, design.WASMFilename)
		return err
	})
}

func (repo *DesignRepository) Delete(id string) error {
	return repo.withConnection(func(conn *sql.DB) error {
		query := `DELETE FROM composer.design WHERE id = $1`
		_, err := conn.Exec(query, id)
		return err
	})
}

func (repo *DesignRepository) GetAll() ([]domain.Design, error) {
	var designs []domain.Design
	err := repo.withConnection(func(conn *sql.DB) error {
		query := `SELECT id, name, description, js_filename, wasm_filename FROM composer.design`
		rows, err := conn.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var design domain.Design
			var descriptionBytes []byte

			err := rows.Scan(&design.ID, &design.Name, &descriptionBytes, &design.JSFilename, &design.WASMFilename)
			if err != nil {
				return err
			}

			err = json.Unmarshal(descriptionBytes, &design.Description)
			if err != nil {
				return err
			}

			designs = append(designs, design)
		}
		return rows.Err()
	})
	return designs, err
}
