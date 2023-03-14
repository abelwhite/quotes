// Filename: internal/models/questions.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Lets model the question table
type Quotes struct {
	QuestionID int64
	Quote      string
	Author     string
	CreatedAt  time.Time
}

// setup dependency injection
type QuestionModel struct {
	DB *sql.DB //connection pool
}

// sql to insert
func (m *QuoteModel) Insert(body string) (int64, error) { //we use QuestionModel because it has acces to connection pool
	var id int64

	statement := `
			INSERT INTO questions(body)
			VALUES($1)  
			RETURNING question_id
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id) //m is the instance, DB. connectionpool, and we want to send query row context
	if err != nil {
		return 0, err
	}
	return id, nil

}

// write SQL code to access the database
// TODO
func (m *QuestionModel) Get() (*Question, error) { //we use QuestionModel because it has acces to connection pool
	var q Question

	statement := `
			SELECT question_id, body 
			FROM questions
			ORDER BY RANDOM()
			LIMIT 1
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.QuestionID, &q.Body) //m is the instNCE, DB. connectionpool, and we want to send query row context
	if err != nil {
		return nil, err
	}
	return &q, err

}
