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
		SELECT id, firstname, lastname, phone_number, date
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

		err := rows.Scan(
			&c.ID,
			&c.Firstname,
			&c.Lastname,
			&c.PhoneNumber,
			&c.Date,
		)
		if err != nil {
			return nil, err
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
		INSERT INTO completions (firstname, lastname, phone_number, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var newID int
	err := s.db.QueryRow(
		query,
		c.Firstname,
		c.Lastname,
		c.PhoneNumber,
		c.Date,
	).Scan(&newID)

	if err != nil {
		return nil, err
	}

	c.ID = newID
	return c, nil
}
