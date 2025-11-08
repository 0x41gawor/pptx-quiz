package service

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type ServiceQuiz struct{}

func NewServiceQuiz() *ServiceQuiz {
	return &ServiceQuiz{}
}

type Answer struct {
	Content string
}

type Question struct {
	Content string
	Answers []Answer
	Correct int // 0–3 (A–D)
}

type Quiz struct {
	Questions []Question
}

func (s *ServiceQuiz) Get(lang string, howMany int) *Quiz {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("cannot open csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}

	if len(records) <= 1 {
		log.Fatal("no data in csv")
	}

	rows := records[1:]

	if howMany > len(rows) {
		howMany = len(rows)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(rows), func(i, j int) { rows[i], rows[j] = rows[j], rows[i] })

	selected := rows[:howMany]
	quiz := &Quiz{}

	for _, row := range selected {
		var (
			content     string
			a, b, c, d  string
			correctStr  string
			langOffset  int
			minRequired int
		)

		switch lang {
		case "en":
			langOffset = 6
			minRequired = 11
		case "pl":
			langOffset = 0
			minRequired = 6
		default:
			log.Fatalf("unsupported language: %s", lang)
		}

		if len(row) < minRequired {
			log.Println("invalid row length, skipping")
			continue
		}

		content = row[langOffset]
		a, b, c, d = row[langOffset+1], row[langOffset+2], row[langOffset+3], row[langOffset+4]
		correctStr = strings.TrimSpace(strings.ToUpper(row[5])) // kolumna "poprawna" wspólna

		answers := []Answer{}
		for _, ans := range []string{a, b, c, d} {
			if ans != "" {
				answers = append(answers, Answer{Content: ans})
			}
		}

		// konwersja "A"/"B"/"C"/"D" → 0–3
		correctIdx := -1
		switch correctStr {
		case "A":
			correctIdx = 0
		case "B":
			correctIdx = 1
		case "C":
			correctIdx = 2
		case "D":
			correctIdx = 3
		}

		quiz.Questions = append(quiz.Questions, Question{
			Content: content,
			Answers: answers,
			Correct: correctIdx,
		})
	}

	return quiz
}
