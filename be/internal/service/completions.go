package service

import (
	"strings"
	"unicode"

	"github.com/0x41gawor/pptx-quiz/internal/repo"
	"github.com/0x41gawor/pptx-quiz/internal/service/model"
	"golang.org/x/text/unicode/norm"
)

type ServiceCompletions struct {
	repo *repo.RepoCompletions
}

func NewServiceCompletions() *ServiceCompletions {
	return &ServiceCompletions{
		repo: repo.NewRepoCompletions(),
	}
}

func (s *ServiceCompletions) List(query string) ([]model.Completion, error) {
	if query != "" {
		models, err := s.repo.List()
		if err != nil {
			return nil, err
		}
		// concatenate models first name and last name and filter by query
		var filtered []model.Completion
		for _, m := range models {
			fullName := m.Firstname + " " + m.Lastname
			if containsIgnoreCase(fullName, query) {
				filtered = append(filtered, m)
			}
		}
		return filtered, nil
	}
	return s.repo.List()
}

func (s *ServiceCompletions) Create(c *model.Completion) (*model.Completion, error) {
	return s.repo.Create(c)
}

func normalizeASCII(s string) string {
	// Normalizacja do decomposed form (np. ą -> a +  ̨)
	t := norm.NFD.String(s)

	// Usuwanie wszystkich znaków diakrytycznych typu Mn (mark, nonspacing)
	b := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b = append(b, unicode.ToLower(r))
	}

	return string(b)
}

func containsIgnoreCase(s, substr string) bool {
	sNorm := normalizeASCII(s)
	substrNorm := normalizeASCII(substr)
	return strings.Contains(sNorm, substrNorm)
}
