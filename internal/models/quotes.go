// Filename: internal/models/questions.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Lets model the question table
type Quote struct {
	QuoteID   int64
	Quote     string
	Author    string
	CreatedAt time.Time
}

// setup dependency injection
type QuoteModel struct {
	DB *sql.DB //connection pool
}

// sql to insert
func (m *QuoteModel) Insert(quote string, author string) (int64, error) { //we use QuestionModel because it has acces to connection pool
	var id int64

	statement := `
			INSERT INTO quotes(quote, author)
			VALUES($1, $2)  
			RETURNING quote_id
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, quote, author).Scan(&id) //m is the instance, DB. connectionpool, and we want to send query row context
	if err != nil {
		return 0, err
	}
	return id, nil

}

// write SQL code to access the database
// TODO
func (m *QuoteModel) Get() (*Quote, error) { //we use QuestionModel because it has acces to connection pool
	var q Quote

	statement := `
			SELECT quote_id, quote, author
			FROM quotes
			ORDER BY RANDOM()
			LIMIT 1
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.QuoteID, &q.Quote, &q.Author) //m is the instNCE, DB. connectionpool, and we want to send query row context
	if err != nil {
		return nil, err
	}
	return &q, err

}
