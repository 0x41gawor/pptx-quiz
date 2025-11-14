package repo

import (
	"database/sql"

	"github.com/0x41gawor/pptx-quiz/internal/service/model"
)

type RepoCompletions struct {
	db *sql.DB
}

func NewRepoCompletions() *RepoCompletions {
	db := GetDatabaseInstance().DB
	return &RepoCompletions{db: db}
}

func (r *RepoCompletions) List() ([]model.Completion, error) {
	rows, err := r.db.Query(`
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

func (r *RepoCompletions) Create(c *model.Completion) (*model.Completion, error) {
	query := `
		INSERT INTO completions (firstname, lastname, phone_number, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var newID int
	err := r.db.QueryRow(
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
