package service

import (
	"database/sql"

	"github.com/0x41gawor/pptx-quiz/internal/repo"
	"github.com/0x41gawor/pptx-quiz/internal/service/model"
)

type ServiceCompletions struct {
	db *sql.DB
}

func NewServiceCompletions() *ServiceCompletions {
	db := repo.GetDatabaseInstance().DB
	return &ServiceCompletions{db: db}
}

func (s *ServiceCompletions) List() ([]model.Completion, error) {
	rows, err := s.db.Query(`
		SELECT id, firstname, lastname, id_passport_number, company, date
		FROM completions
		ORDER BY date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var completions []model.Completion

	for rows.Next() {
		var c model.Completion
		var idPassportNumber, company sql.NullString

		err := rows.Scan(
			&c.ID,
			&c.Firstname,
			&c.Lastname,
			&idPassportNumber,
			&company,
			&c.Date,
		)
		if err != nil {
			return nil, err
		}

		// Konwersja sql.NullString → *string
		if idPassportNumber.Valid {
			c.IDPassportNumber = &idPassportNumber.String
		}
		if company.Valid {
			c.Company = &company.String
		}

		completions = append(completions, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return completions, nil
}

func (s *ServiceCompletions) Create(c *model.Completion) (*model.Completion, error) {
	query := `
		INSERT INTO completions (firstname, lastname, id_passport_number, company, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newID int
	err := s.db.QueryRow(
		query,
		c.Firstname,
		c.Lastname,
		c.IDPassportNumber,
		c.Company,
		c.Date,
	).Scan(&newID)

	if err != nil {
		return nil, err
	}

	c.ID = newID
	return c, nil
}
