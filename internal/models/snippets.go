package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, interval int) (int, error) {
	stmt := `
	insert into snippets (title, content, created, expires)
	values(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`

	result, err := m.DB.Exec(stmt, title, content, interval)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `
	select id, title, content, created, expires from snippets
	where id = ? and expires > UTC_TIMESTAMP();
	`

	var s Snippet
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `
	select id, title, content, created, expires from snippets
	where expires > UTC_TIMESTAMP() order by id desc limit 10;
	`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var s Snippet
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
